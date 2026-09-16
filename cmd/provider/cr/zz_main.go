/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/alecthomas/kingpin/v2"
	apisCluster "github.com/crossplane-contrib/provider-alibabacloud/apis/cluster"
	apisNamespaced "github.com/crossplane-contrib/provider-alibabacloud/apis/namespaced"
	"github.com/crossplane-contrib/provider-alibabacloud/config"
	"github.com/crossplane-contrib/provider-alibabacloud/internal/clients"
	controllerCluster "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster"
	controllerNamespaced "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced"
	"github.com/crossplane-contrib/provider-alibabacloud/internal/features"
	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/gate"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/customresourcesgate"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/statemetrics"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/pkg/errors"
	authv1 "k8s.io/api/authorization/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

func main() {
	var (
		app                     = kingpin.New(filepath.Base(os.Args[0]), "Terraform based Crossplane provider for AlibabaCloud cr").DefaultEnvars()
		debug                   = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncPeriod              = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		pollInterval            = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		pollStateMetricInterval = app.Flag("poll-state-metric", "State metric recording interval").Default("5s").Duration()
		leaderElection          = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		maxReconcileRate        = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may be checked for drift from the desired state.").Default("10").Int()

		enableManagementPolicies = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()
	)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName("provider-upjet-alibabacloud-cr"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	log.Debug("Starting", "sync-period", syncPeriod.String(), "poll-interval", pollInterval.String(), "max-reconcile-rate", *maxReconcileRate)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-provider-upjet-alibabacloud-cr",
		Cache: cache.Options{
			SyncPeriod: syncPeriod,
		},
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
		RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")
	kingpin.FatalIfError(apisCluster.AddToScheme(mgr.GetScheme()), "Cannot add cluster-scoped AlibabaCloud APIs to scheme")
	kingpin.FatalIfError(apisNamespaced.AddToScheme(mgr.GetScheme()), "Cannot add namespaced AlibabaCloud APIs to scheme")
	kingpin.FatalIfError(authv1.AddToScheme(mgr.GetScheme()), "Cannot add Kubernetes authorization APIs to scheme")

	metricRecorder := managed.NewMRMetricRecorder()
	stateMetrics := statemetrics.NewMRStateMetrics()

	metrics.Registry.MustRegister(metricRecorder)
	metrics.Registry.MustRegister(stateMetrics)

	ctx := context.Background()
	providerCluster, err := config.GetProvider(ctx, false)
	kingpin.FatalIfError(err, "Cannot initialize the cluster-scoped provider configuration")
	providerNamespaced, err := config.GetNamespacedProvider(ctx, false)
	kingpin.FatalIfError(err, "Cannot initialize the namespaced provider configuration")

	xpOpts := xpcontroller.Options{
		Logger:                  log,
		GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
		PollInterval:            *pollInterval,
		MaxConcurrentReconciles: *maxReconcileRate,
		Features:                &feature.Flags{},
		MetricOptions: &xpcontroller.MetricOptions{
			PollStateMetricInterval: *pollStateMetricInterval,
			MRMetrics:               metricRecorder,
			MRStateMetrics:          stateMetrics,
		},
	}

	clusterOpts := tjcontroller.Options{
		Options:               xpOpts,
		Provider:              providerCluster,
		SetupFn:               clients.TerraformSetupBuilder(providerCluster.TerraformProvider),
		OperationTrackerStore: tjcontroller.NewOperationStore(log),
	}

	namespacedOpts := tjcontroller.Options{
		Options:               xpOpts,
		Provider:              providerNamespaced,
		SetupFn:               clients.TerraformSetupBuilder(providerNamespaced.TerraformProvider),
		OperationTrackerStore: tjcontroller.NewOperationStore(log),
	}
	// Options embeds xpcontroller.Options by value, so each copy needs its own
	// Features: enabling a flag on one must not enable it on the other.
	clusterOpts.Features = &feature.Flags{}
	namespacedOpts.Features = &feature.Flags{}

	if *enableManagementPolicies {
		clusterOpts.Features.Enable(features.EnableBetaManagementPolicies)
		namespacedOpts.Features.Enable(features.EnableBetaManagementPolicies)
		log.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	// SafeStart defers controller setup until the CRDs they watch exist, which
	// needs permission to watch CRDs. Without it the controllers are started
	// eagerly, as they were before, so a provider lacking the RBAC still runs.
	canSafeStart, err := canWatchCRD(ctx, mgr)
	kingpin.FatalIfError(err, "SafeStart precheck failed")
	if canSafeStart {
		crdGate := new(gate.Gate[schema.GroupVersionKind])
		clusterOpts.Gate = crdGate
		namespacedOpts.Gate = crdGate
		kingpin.FatalIfError(customresourcesgate.Setup(mgr, xpcontroller.Options{
			Logger:                  log,
			Gate:                    crdGate,
			MaxConcurrentReconciles: 1,
		}), "Cannot setup CRD gate")
		kingpin.FatalIfError(controllerCluster.SetupGated_cr(mgr, clusterOpts), "Cannot setup cluster-scoped AlibabaCloud cr controllers")
		kingpin.FatalIfError(controllerNamespaced.SetupGated_cr(mgr, namespacedOpts), "Cannot setup namespaced AlibabaCloud cr controllers")
	} else {
		log.Info("Provider has missing RBAC permissions for watching CRDs, controller SafeStart capability will be disabled")
		kingpin.FatalIfError(controllerCluster.Setup_cr(mgr, clusterOpts), "Cannot setup cluster-scoped AlibabaCloud cr controllers")
		kingpin.FatalIfError(controllerNamespaced.Setup_cr(mgr, namespacedOpts), "Cannot setup namespaced AlibabaCloud cr controllers")
	}

	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}

// canWatchCRD reports whether the provider's service account may get, list and
// watch CustomResourceDefinitions, which the CRD gate requires.
func canWatchCRD(ctx context.Context, mgr manager.Manager) (bool, error) {
	verbs := []string{"get", "list", "watch"}
	for _, verb := range verbs {
		sar := &authv1.SelfSubjectAccessReview{
			Spec: authv1.SelfSubjectAccessReviewSpec{
				ResourceAttributes: &authv1.ResourceAttributes{
					Group:    "apiextensions.k8s.io",
					Resource: "customresourcedefinitions",
					Verb:     verb,
				},
			},
		}
		if err := mgr.GetClient().Create(ctx, sar); err != nil {
			return false, errors.Wrapf(err, "unable to perform RBAC check for verb %s on CustomResourceDefinitions", verb)
		}
		if !sar.Status.Allowed {
			return false, nil
		}
	}
	return true, nil
}
