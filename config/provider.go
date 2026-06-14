/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane-contrib/provider-alibabacloud/config/fcv3"
	"github.com/crossplane-contrib/provider-alibabacloud/config/slb"

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
	"github.com/crossplane/upjet/pkg/registry/reference"

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

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	defaultResourceOptions := []ujconfig.ResourceOption{
		ExternalNameConfigurations(),
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
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithMainTemplate(hack.MainTemplate),
		ujconfig.WithDefaultResourceOptions(defaultResourceOptions...))

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
	return pc
}
