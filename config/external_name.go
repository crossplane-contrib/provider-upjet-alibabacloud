/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	// ECS
	"alicloud_security_group": config.IdentifierFromProvider,
	// PolarDB
	"alicloud_polardb_account":                 config.IdentifierFromProvider,
	"alicloud_polardb_account_privilege":       config.IdentifierFromProvider,
	"alicloud_polardb_backup_policy":           config.IdentifierFromProvider,
	"alicloud_polardb_cluster":                 config.IdentifierFromProvider,
	"alicloud_polardb_cluster_endpoint":        config.IdentifierFromProvider,
	"alicloud_polardb_database":                config.IdentifierFromProvider,
	"alicloud_polardb_endpoint":                config.IdentifierFromProvider,
	"alicloud_polardb_endpoint_address":        config.IdentifierFromProvider,
	"alicloud_polardb_global_database_network": config.IdentifierFromProvider,
	"alicloud_polardb_parameter_group":         config.IdentifierFromProvider,
	"alicloud_polardb_primary_endpoint":        config.IdentifierFromProvider,
	// VPC
	"alicloud_vpc":     config.IdentifierFromProvider,
	"alicloud_vswitch": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
