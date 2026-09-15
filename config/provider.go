/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"context"
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	alicloud "github.com/aliyun/terraform-provider-alicloud/alicloud"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/provider-alibabacloud/config/fcv3"
	"github.com/crossplane-contrib/provider-alibabacloud/config/slb"

	"github.com/crossplane/upjet/pkg/registry/reference"

	"github.com/crossplane-contrib/provider-alibabacloud/config/ack"
	"github.com/crossplane-contrib/provider-alibabacloud/config/ackone"
	"github.com/crossplane-contrib/provider-alibabacloud/config/alb"
	"github.com/crossplane-contrib/provider-alibabacloud/config/alidns"
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

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "alicloud"
	modulePath     = "github.com/crossplane-contrib/provider-alibabacloud"
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
func GetProvider(_ context.Context, generationProvider bool) (*ujconfig.Provider, error) {
	var p *schema.Provider
	var err error
	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p = alicloud.Provider()
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
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
		ujconfig.WithRootGroup("alibabacloud.crossplane.io"),
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
