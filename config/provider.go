/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"context"
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	alicloud "github.com/aliyun/terraform-provider-alicloud/alicloud"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/provider-alibabacloud/config/fcv3"
	"github.com/crossplane-contrib/provider-alibabacloud/config/slb"

	"github.com/crossplane/upjet/v2/pkg/registry/reference"

	"github.com/crossplane-contrib/provider-alibabacloud/config/ack"
	"github.com/crossplane-contrib/provider-alibabacloud/config/ackone"
	"github.com/crossplane-contrib/provider-alibabacloud/config/alb"
	"github.com/crossplane-contrib/provider-alibabacloud/config/alidns"
	"github.com/crossplane-contrib/provider-alibabacloud/config/alikafka"
	"github.com/crossplane-contrib/provider-alibabacloud/config/cdn"
	"github.com/crossplane-contrib/provider-alibabacloud/config/cloudmonitorservice"
	"github.com/crossplane-contrib/provider-alibabacloud/config/cr"
	"github.com/crossplane-contrib/provider-alibabacloud/config/ecs"
	"github.com/crossplane-contrib/provider-alibabacloud/config/kms"
	"github.com/crossplane-contrib/provider-alibabacloud/config/messageservice"
	"github.com/crossplane-contrib/provider-alibabacloud/config/oos"
	"github.com/crossplane-contrib/provider-alibabacloud/config/oss"
	"github.com/crossplane-contrib/provider-alibabacloud/config/polardb"
	"github.com/crossplane-contrib/provider-alibabacloud/config/privatelink"
	"github.com/crossplane-contrib/provider-alibabacloud/config/quotas"
	"github.com/crossplane-contrib/provider-alibabacloud/config/ram"
	"github.com/crossplane-contrib/provider-alibabacloud/config/tair"
	"github.com/crossplane-contrib/provider-alibabacloud/config/vpc"
	"github.com/crossplane-contrib/provider-alibabacloud/hack"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	resourcePrefix = "alicloud"
	modulePath     = "github.com/crossplane-contrib/provider-alibabacloud"

	// clusterRootGroup is the API group for the legacy cluster-scoped managed
	// resources, and namespacedRootGroup the one for Crossplane v2 namespaced
	// managed resources. The ".m." infix is the convention upjet recommends for
	// telling the two apart.
	clusterRootGroup    = "alibabacloud.crossplane.io"
	namespacedRootGroup = "alibabacloud.m.crossplane.io"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// getProviderSchema builds a schema.Provider out of the Terraform JSON schema
// document. It carries no CRUD implementations, so it is only good enough for
// code generation, where the schema is all that is read.
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal the Terraform JSON schema")
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration. When generationProvider is true,
// the Terraform provider is reconstructed from the embedded JSON schema, which
// keeps code generation independent of the upstream provider's Go code. At
// runtime it is the real upstream provider, whose CRUD functions the plugin SDK
// external client calls directly.
func GetProvider(ctx context.Context, generationProvider bool) (*ujconfig.Provider, error) {
	return newProvider(ctx, generationProvider, clusterRootGroup)
}

// GetNamespacedProvider returns the provider configuration for namespaced
// managed resources. It differs from GetProvider only in the root API group:
// namespaced MRs are served under "alibabacloud.m.crossplane.io" so they can be
// told apart from their cluster-scoped counterparts.
//
// The per-service resource configurations are shared between the two. They are
// scope-independent — none of them reference the apis packages, pin a CRD
// version or register conversions — so unlike some other upjet providers there
// is no need to maintain a duplicate copy of every config/<group> package.
func GetNamespacedProvider(ctx context.Context, generationProvider bool) (*ujconfig.Provider, error) {
	return newProvider(ctx, generationProvider, namespacedRootGroup)
}

func newProvider(_ context.Context, generationProvider bool, rootGroup string) (*ujconfig.Provider, error) {
	// The runtime schema is the upstream provider's own Go schema, which the
	// CRUD functions execute against. Code generation deliberately uses the
	// JSON schema instead, to keep the generated CRD APIs stable: the Go
	// schema would, among other things, widen some number fields differently.
	//
	// The two are not identical, though: the JSON schema does not faithfully
	// carry MaxItems, so a list the Go schema constrains to one element can
	// come back unconstrained. Left alone, the generated API would model such
	// a field as an array while the runtime conversion functions expect a
	// single object. Sync the constraints across before generating, as
	// provider-upjet-aws does for the same reason.
	//
	// Today this is a no-op for the resources we expose: the only divergence
	// in the upstream schema is alicloud_ros_stack_instances.deployment_options,
	// which is not in the include list. It is kept because the divergence is a
	// property of the two schema sources, not of the current include list.
	p := alicloud.Provider()
	if generationProvider {
		gp, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot get the Terraform provider schema from the embedded JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(p.ResourcesMap).Traverse(traverser.NewMaxItemsSync(gp.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		p = gp
	}

	defaultResourceOptions := []ujconfig.ResourceOption{
		ResourceConfigurator(),
		RegionAddition(),
		IdentifierAssignedByAlibabaCloud(),
		KnownReferences(),
		NamePrefixRemoval(),
		AddExternalTagsField(),
		DocumentationForTags(),
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithShortName("alibabacloud"),
		ujconfig.WithRootGroup(rootGroup),
		ujconfig.WithIncludeList(resourceList(CLIReconciledExternalNameConfigs)),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(terraformPluginSDKExternalNameConfigs)),
		ujconfig.WithTerraformProvider(p),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithMainTemplate(hack.MainTemplate),
		ujconfig.WithDefaultResourceOptions(defaultResourceOptions...))

	// Schema omissions shape the generated CRDs and must never touch the live
	// runtime schema, which the upstream CRUD functions execute against.
	if generationProvider {
		addGenerationOnlySchemaOmissions(pc)
	}

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		ack.Configure,
		ackone.Configure,
		alb.Configure,
		alikafka.Configure,
		alidns.Configure,
		cdn.Configure,
		cloudmonitorservice.Configure,
		cr.Configure,
		ecs.Configure,
		fcv3.Configure,
		kms.Configure,
		messageservice.Configure,
		oos.Configure,
		oss.Configure,
		polardb.Configure,
		privatelink.Configure,
		quotas.Configure,
		ram.Configure,
		slb.Configure,
		tair.Configure,
		vpc.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}
