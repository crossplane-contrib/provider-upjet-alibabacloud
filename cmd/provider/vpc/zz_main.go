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
	"github.com/crossplane-contrib/provider-alibabacloud/apis"
	"github.com/crossplane-contrib/provider-alibabacloud/config"
	"github.com/crossplane-contrib/provider-alibabacloud/internal/clients"
	"github.com/crossplane-contrib/provider-alibabacloud/internal/controller"
	"github.com/crossplane-contrib/provider-alibabacloud/internal/features"
	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/statemetrics"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/crossplane/upjet/v2/pkg/diffserver"
	"k8s.io/apimachinery/pkg/runtime"
	corev1 "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

func main() {
	var (
		app                     = kingpin.New(filepath.Base(os.Args[0]), "Terraform based Crossplane provider for AlibabaCloud vpc").DefaultEnvars()
		debug                   = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncPeriod              = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		pollInterval            = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		pollStateMetricInterval = app.Flag("poll-state-metric", "State metric recording interval").Default("5s").Duration()
		leaderElection          = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		maxReconcileRate        = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may be checked for drift from the desired state.").Default("10").Int()

		enableManagementPolicies = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()

		// Running the provider's controllers is the default command, so that
		// the provider keeps behaving as it did before the internal commands
		// were introduced, i.e. when it's invoked without any arguments.
		startCmd = app.Command("start", "Start the provider's controllers.").Default()

		internalCmd   = app.Command("internal", "Commands used internally by the Crossplane ecosystem. No compatibility guarantees are made for them.")
		diffServerCmd = internalCmd.Command("diff-server", "Start a gRPC server serving the provider diff services.")
		diffNetwork   = diffServerCmd.Flag("network", "The network the diff gRPC server listens on.").Default("tcp").Envar("DIFF_SERVER_NETWORK").Enum("tcp", "unix")
		diffAddress   = diffServerCmd.Flag("address", "The address the diff gRPC server listens on. A socket path when the network is unix.").Default(":9099").Envar("DIFF_SERVER_ADDRESS").String()
	)

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName("provider-upjet-alibabacloud-vpc"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	ctx := context.Background()
	provider, err := config.GetProvider(ctx, false)
	kingpin.FatalIfError(err, "Cannot initialize the provider configuration")

	switch cmd {
	case diffServerCmd.FullCommand():
		// The diff server only needs to deserialize this provider's managed
		// resources, so its scheme holds just the provider's APIs plus the
		// Kubernetes core APIs the request's object store may carry
		// (ProviderConfigs reference Secrets, which are core objects).
		diffScheme := runtime.NewScheme()
		kingpin.FatalIfError(corev1.AddToScheme(diffScheme), "Cannot add Kubernetes core APIs to the diff server scheme")
		kingpin.FatalIfError(apis.AddToScheme(diffScheme), "Cannot add AlibabaCloud APIs to the diff server scheme")

		s := diffserver.NewServer(
			diffserver.WithProviderConfigurations(provider),
			diffserver.WithLogger(log),
			diffserver.WithTerraformSetupFn(clients.OfflineTerraformSetupBuilder(provider.TerraformProvider)),
			
			// This package's controllers reconcile a single API group, while
			// its scheme and provider configuration cover all of them. Declare
			// the package's group so that a plan request for any other one is
			// declined rather than answered from a configuration that does not
			// run here.
			diffserver.WithAPIGroups("vpc"),
			
		)
		kingpin.FatalIfError(s.Serve(ctrl.SetupSignalHandler(), *diffNetwork, *diffAddress, diffScheme), "Cannot run the diff gRPC server")
		return
	case startCmd.FullCommand():
		// the provider's controllers are started below.
	}

	log.Debug("Starting", "sync-period", syncPeriod.String(), "poll-interval", pollInterval.String(), "max-reconcile-rate", *maxReconcileRate)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-provider-upjet-alibabacloud-vpc",
		Cache: cache.Options{
			SyncPeriod: syncPeriod,
		},
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
		RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")
	kingpin.FatalIfError(apis.AddToScheme(mgr.GetScheme()), "Cannot add AlibabaCloud APIs to scheme")

	metricRecorder := managed.NewMRMetricRecorder()
	stateMetrics := statemetrics.NewMRStateMetrics()

	metrics.Registry.MustRegister(metricRecorder)
	metrics.Registry.MustRegister(stateMetrics)

	o := tjcontroller.Options{
		Options: xpcontroller.Options{
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
		},
		Provider:              provider,
		SetupFn:               clients.TerraformSetupBuilder(provider.TerraformProvider),
		OperationTrackerStore: tjcontroller.NewOperationStore(log),
	}

	if *enableManagementPolicies {
		o.Features.Enable(features.EnableBetaManagementPolicies)
		log.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	kingpin.FatalIfError(controller.Setup_vpc(mgr, o), "Cannot setup AlibabaCloud vpc controllers")
	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}
