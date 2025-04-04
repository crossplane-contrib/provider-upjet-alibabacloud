/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	// ALIDNS
	//"alicloud_alidns_access_strategy":   config.IdentifierFromProvider,
	"alicloud_alidns_address_pool":      config.IdentifierFromProvider,
	"alicloud_alidns_custom_line":       config.IdentifierFromProvider,
	"alicloud_alidns_domain":            config.IdentifierFromProvider,
	"alicloud_alidns_domain_attachment": config.IdentifierFromProvider,
	"alicloud_alidns_domain_group":      config.IdentifierFromProvider,
	"alicloud_alidns_gtm_instance":      config.IdentifierFromProvider,
	"alicloud_alidns_instance":          config.IdentifierFromProvider,
	"alicloud_alidns_monitor_config":    config.IdentifierFromProvider,
	"alicloud_alidns_record":            config.IdentifierFromProvider,
	// CDN
	"alicloud_cdn_domain_config": config.IdentifierFromProvider,
	"alicloud_cdn_domain_new":    config.IdentifierFromProvider,
	"alicloud_cdn_fc_trigger":    config.IdentifierFromProvider,
	// CloudMonitorService
	"alicloud_cms_alarm_contact_group": config.IdentifierFromProvider,
	// ECS
	"alicloud_security_group": config.IdentifierFromProvider,
	// MessageService
	"alicloud_message_service_endpoint":     config.IdentifierFromProvider,
	"alicloud_message_service_endpoint_acl": config.IdentifierFromProvider,
	"alicloud_message_service_queue":        config.IdentifierFromProvider,
	"alicloud_message_service_subscription": config.IdentifierFromProvider,
	"alicloud_message_service_topic":        config.IdentifierFromProvider,
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
	// Tair
	"alicloud_kvstore_account":          config.IdentifierFromProvider,
	"alicloud_kvstore_audit_log_config": config.IdentifierFromProvider,
	"alicloud_kvstore_connection":       config.IdentifierFromProvider,
	"alicloud_kvstore_instance":         config.IdentifierFromProvider,
	"alicloud_redis_tair_instance":      config.IdentifierFromProvider,
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
