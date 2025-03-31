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
	"alicloud_auto_provisioning_group":             config.IdentifierFromProvider,
	"alicloud_disk":                                config.IdentifierFromProvider,
	"alicloud_disk_attachment":                     config.IdentifierFromProvider,
	"alicloud_ecs_activation":                      config.IdentifierFromProvider,
	"alicloud_ecs_auto_snapshot_policy":            config.IdentifierFromProvider,
	"alicloud_ecs_auto_snapshot_policy_attachment": config.IdentifierFromProvider,
	"alicloud_ecs_capacity_reservation":            config.IdentifierFromProvider,
	"alicloud_ecs_command":                         config.IdentifierFromProvider,
	"alicloud_ecs_dedicated_host":                  config.IdentifierFromProvider,
	"alicloud_ecs_dedicated_host_cluster":          config.IdentifierFromProvider,
	"alicloud_ecs_deployment_set":                  config.IdentifierFromProvider,
	"alicloud_ecs_disk":                            config.IdentifierFromProvider,
	"alicloud_ecs_disk_attachment":                 config.IdentifierFromProvider,
	"alicloud_ecs_elasticity_assurance":            config.IdentifierFromProvider,
	"alicloud_ecs_hpc_cluster":                     config.IdentifierFromProvider,
	"alicloud_ecs_image_component":                 config.IdentifierFromProvider,
	"alicloud_ecs_image_pipeline":                  config.IdentifierFromProvider,
	"alicloud_ecs_image_pipeline_execution":        config.IdentifierFromProvider,
	"alicloud_ecs_instance_set":                    config.IdentifierFromProvider,
	"alicloud_ecs_invocation":                      config.IdentifierFromProvider,
	"alicloud_ecs_key_pair":                        config.IdentifierFromProvider,
	"alicloud_ecs_key_pair_attachment":             config.IdentifierFromProvider,
	"alicloud_ecs_launch_template":                 config.IdentifierFromProvider,
	"alicloud_ecs_network_interface":               config.IdentifierFromProvider,
	"alicloud_ecs_network_interface_attachment":    config.IdentifierFromProvider,
	"alicloud_ecs_network_interface_permission":    config.IdentifierFromProvider,
	"alicloud_ecs_prefix_list":                     config.IdentifierFromProvider,
	"alicloud_ecs_session_manager_status":          config.IdentifierFromProvider,
	"alicloud_ecs_snapshot":                        config.IdentifierFromProvider,
	"alicloud_ecs_snapshot_group":                  config.IdentifierFromProvider,
	"alicloud_ecs_storage_capacity_unit":           config.IdentifierFromProvider,
	"alicloud_image":                               config.IdentifierFromProvider,
	"alicloud_image_copy":                          config.IdentifierFromProvider,
	"alicloud_image_export":                        config.IdentifierFromProvider,
	"alicloud_image_import":                        config.IdentifierFromProvider,
	"alicloud_image_share_permission":              config.IdentifierFromProvider,
	"alicloud_instance":                            config.IdentifierFromProvider,
	"alicloud_key_pair":                            config.IdentifierFromProvider,
	"alicloud_key_pair_attachment":                 config.IdentifierFromProvider,
	"alicloud_launch_template":                     config.IdentifierFromProvider,
	"alicloud_network_interface":                   config.IdentifierFromProvider,
	"alicloud_network_interface_attachment":        config.IdentifierFromProvider,
	"alicloud_ram_role_attachment":                 config.IdentifierFromProvider,
	"alicloud_reserved_instance":                   config.IdentifierFromProvider,
	"alicloud_security_group":                      config.IdentifierFromProvider,
	"alicloud_security_group_rule":                 config.IdentifierFromProvider,
	"alicloud_snapshot":                            config.IdentifierFromProvider,
	"alicloud_snapshot_policy":                     config.IdentifierFromProvider,

	// KMS
	"alicloud_kms_alias":                    config.IdentifierFromProvider,
	"alicloud_kms_application_access_point": config.IdentifierFromProvider,
	"alicloud_kms_ciphertext":               config.IdentifierFromProvider,
	"alicloud_kms_client_key":               config.IdentifierFromProvider,
	"alicloud_kms_instance":                 config.IdentifierFromProvider,
	"alicloud_kms_key":                      config.IdentifierFromProvider,
	"alicloud_kms_key_version":              config.IdentifierFromProvider,
	"alicloud_kms_network_rule":             config.IdentifierFromProvider,
	"alicloud_kms_policy":                   config.IdentifierFromProvider,
	"alicloud_kms_secret":                   config.IdentifierFromProvider,

	// OSS
	"alicloud_oss_access_point":                      config.IdentifierFromProvider,
	"alicloud_oss_account_public_access_block":       config.IdentifierFromProvider,
	"alicloud_oss_bucket":                            config.IdentifierFromProvider,
	"alicloud_oss_bucket_access_monitor":             config.IdentifierFromProvider,
	"alicloud_oss_bucket_acl":                        config.IdentifierFromProvider,
	"alicloud_oss_bucket_cname":                      config.IdentifierFromProvider,
	"alicloud_oss_bucket_cname_token":                config.IdentifierFromProvider,
	"alicloud_oss_bucket_cors":                       config.IdentifierFromProvider,
	"alicloud_oss_bucket_data_redundancy_transition": config.IdentifierFromProvider,
	"alicloud_oss_bucket_https_config":               config.IdentifierFromProvider,
	"alicloud_oss_bucket_logging":                    config.IdentifierFromProvider,
	"alicloud_oss_bucket_meta_query":                 config.IdentifierFromProvider,
	"alicloud_oss_bucket_object":                     config.IdentifierFromProvider,
	"alicloud_oss_bucket_policy":                     config.IdentifierFromProvider,
	"alicloud_oss_bucket_public_access_block":        config.IdentifierFromProvider,
	"alicloud_oss_bucket_referer":                    config.IdentifierFromProvider,
	"alicloud_oss_bucket_replication":                config.IdentifierFromProvider,
	"alicloud_oss_bucket_request_payment":            config.IdentifierFromProvider,
	"alicloud_oss_bucket_server_side_encryption":     config.IdentifierFromProvider,
	"alicloud_oss_bucket_transfer_acceleration":      config.IdentifierFromProvider,
	"alicloud_oss_bucket_user_defined_log_fields":    config.IdentifierFromProvider,
	"alicloud_oss_bucket_versioning":                 config.IdentifierFromProvider,
	"alicloud_oss_bucket_website":                    config.IdentifierFromProvider,
	"alicloud_oss_bucket_worm":                       config.IdentifierFromProvider,

	// Quotas
	"alicloud_quotas_quota_alarm":           config.IdentifierFromProvider,
	"alicloud_quotas_quota_application":     config.IdentifierFromProvider,
	"alicloud_quotas_template_applications": config.IdentifierFromProvider,
	"alicloud_quotas_template_quota":        config.IdentifierFromProvider,
	"alicloud_quotas_template_service":      config.IdentifierFromProvider,

	// RAM
	"alicloud_ram_access_key":              config.IdentifierFromProvider,
	"alicloud_ram_account_alias":           config.IdentifierFromProvider,
	"alicloud_ram_account_password_policy": config.IdentifierFromProvider,
	"alicloud_ram_group":                   config.IdentifierFromProvider,
	"alicloud_ram_group_membership":        config.IdentifierFromProvider,
	"alicloud_ram_group_policy_attachment": config.IdentifierFromProvider,
	"alicloud_ram_login_profile":           config.IdentifierFromProvider,
	"alicloud_ram_policy":                  config.IdentifierFromProvider,
	"alicloud_ram_role":                    config.IdentifierFromProvider,
	"alicloud_ram_role_policy_attachment":  config.IdentifierFromProvider,
	"alicloud_ram_saml_provider":           config.IdentifierFromProvider,
	"alicloud_ram_security_preference":     config.IdentifierFromProvider,
	"alicloud_ram_user":                    config.IdentifierFromProvider,
	"alicloud_ram_user_policy_attachment":  config.IdentifierFromProvider,

	// VPC
	"alicloud_vpc":                                         config.IdentifierFromProvider,
	"alicloud_vpc_dhcp_options_set":                        config.IdentifierFromProvider,
	"alicloud_vpc_dhcp_options_set_attachment":             config.IdentifierFromProvider,
	"alicloud_vpc_flow_log":                                config.IdentifierFromProvider,
	"alicloud_vpc_gateway_endpoint":                        config.IdentifierFromProvider,
	"alicloud_vpc_gateway_endpoint_route_table_attachment": config.IdentifierFromProvider,
	"alicloud_vpc_gateway_route_table_attachment":          config.IdentifierFromProvider,
	"alicloud_vpc_ha_vip":                                  config.IdentifierFromProvider,
	"alicloud_vpc_ipv4_cidr_block":                         config.IdentifierFromProvider,
	"alicloud_vpc_ipv4_gateway":                            config.IdentifierFromProvider,
	"alicloud_vpc_ipv6_address":                            config.IdentifierFromProvider,
	"alicloud_vpc_ipv6_egress_rule":                        config.IdentifierFromProvider,
	"alicloud_vpc_ipv6_gateway":                            config.IdentifierFromProvider,
	"alicloud_vpc_ipv6_internet_bandwidth":                 config.IdentifierFromProvider,
	"alicloud_vpc_network_acl_attachment":                  config.IdentifierFromProvider,
	"alicloud_vpc_peer_connection":                         config.IdentifierFromProvider,
	"alicloud_vpc_peer_connection_accepter":                config.IdentifierFromProvider,
	"alicloud_vpc_prefix_list":                             config.IdentifierFromProvider,
	"alicloud_vpc_public_ip_address_pool":                  config.IdentifierFromProvider,
	"alicloud_vpc_public_ip_address_pool_cidr_block":       config.IdentifierFromProvider,
	"alicloud_vpc_traffic_mirror_filter":                   config.IdentifierFromProvider,
	"alicloud_vpc_traffic_mirror_filter_egress_rule":       config.IdentifierFromProvider,
	"alicloud_vpc_traffic_mirror_filter_ingress_rule":      config.IdentifierFromProvider,
	"alicloud_vpc_traffic_mirror_session":                  config.IdentifierFromProvider,
	"alicloud_vpc_vswitch_cidr_reservation":                config.IdentifierFromProvider,
	"alicloud_vswitch":                                     config.IdentifierFromProvider,
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
