// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/name"
)

// GroupKindCalculator returns the correct group and kind name for given TF resource.
type GroupKindCalculator func(resource string) (string, string)

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		// "alicloud_instance": "ecs" -> (ecs, Instance)
		words := strings.Split(strings.TrimPrefix(resource, "alicloud_"), "_")
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

// GroupMap contains all overrides we'd like to make to the default group search.
// This maps Terraform resources to their appropriate Crossplane API groups.
// Based on the Alibaba Cloud service grouping for all 16 resource groups.
var GroupMap = map[string]GroupKindCalculator{
	// ACK - Alibaba Container Service for Kubernetes resources
	"alicloud_cs_autoscaling_config":     ReplaceGroupWords("ack", 1),
	"alicloud_cs_edge_kubernetes":        ReplaceGroupWords("ack", 1),
	"alicloud_cs_kubernetes":             ReplaceGroupWords("ack", 1),
	"alicloud_cs_kubernetes_addon":       ReplaceGroupWords("ack", 1),
	"alicloud_cs_kubernetes_node_pool":   ReplaceGroupWords("ack", 1),
	"alicloud_cs_kubernetes_permissions": ReplaceGroupWords("ack", 1),
	"alicloud_cs_managed_kubernetes":     ReplaceGroupWords("ack", 1),
	"alicloud_cs_serverless_kubernetes":  ReplaceGroupWords("ack", 1),

	// ACKONE - ACK One resources
	"alicloud_ack_one_cluster":               ReplaceGroupWords("ackone", 2),
	"alicloud_ack_one_membership_attachment": ReplaceGroupWords("ackone", 2),

	// ALB - Application Load Balancer resources
	"alicloud_alb_acl":                                     ReplaceGroupWords("alb", 1),
	"alicloud_alb_acl_entry_attachment":                    ReplaceGroupWords("alb", 1),
	"alicloud_alb_ascript":                                 ReplaceGroupWords("alb", 1),
	"alicloud_alb_health_check_template":                   ReplaceGroupWords("alb", 1),
	"alicloud_alb_listener":                                ReplaceGroupWords("alb", 1),
	"alicloud_alb_listener_acl_attachment":                 ReplaceGroupWords("alb", 1),
	"alicloud_alb_load_balancer":                           ReplaceGroupWords("alb", 1),
	"alicloud_alb_load_balancer_security_group_attachment": ReplaceGroupWords("alb", 1),
	"alicloud_alb_load_balancer_zone_shifted_attachment":   ReplaceGroupWords("alb", 1),
	"alicloud_alb_rule":                                    ReplaceGroupWords("alb", 1),
	"alicloud_alb_security_policy":                         ReplaceGroupWords("alb", 1),
	"alicloud_alb_server_group":                            ReplaceGroupWords("alb", 1),

	// ALIDNS - Alibaba DNS resources
	"alicloud_alidns_address_pool":      ReplaceGroupWords("alidns", 1),
	"alicloud_alidns_custom_line":       ReplaceGroupWords("alidns", 1),
	"alicloud_alidns_domain":            ReplaceGroupWords("alidns", 1),
	"alicloud_alidns_domain_attachment": ReplaceGroupWords("alidns", 1),
	"alicloud_alidns_domain_group":      ReplaceGroupWords("alidns", 1),
	"alicloud_alidns_gtm_instance":      ReplaceGroupWords("alidns", 1),
	"alicloud_alidns_instance":          ReplaceGroupWords("alidns", 1),
	"alicloud_alidns_monitor_config":    ReplaceGroupWords("alidns", 1),
	"alicloud_alidns_record":            ReplaceGroupWords("alidns", 1),

	// CDN - Content Delivery Network resources
	"alicloud_cdn_domain_config": ReplaceGroupWords("cdn", 1),
	"alicloud_cdn_domain_new":    ReplaceGroupWords("cdn", 1),
	"alicloud_cdn_fc_trigger":    ReplaceGroupWords("cdn", 1),

	// CLOUDMONITORSERVICE - Cloud Monitor Service resources
	"alicloud_cms_alarm_contact_group": ReplaceGroupWords("cloudmonitorservice", 1),

	// ECS - Elastic Compute Service resources
	"alicloud_instance":                            ReplaceGroupWords("ecs", 0),
	"alicloud_security_group":                      ReplaceGroupWords("ecs", 0),
	"alicloud_security_group_rule":                 ReplaceGroupWords("ecs", 0),
	"alicloud_image":                               ReplaceGroupWords("ecs", 0),
	"alicloud_image_copy":                          ReplaceGroupWords("ecs", 0),
	"alicloud_image_export":                        ReplaceGroupWords("ecs", 0),
	"alicloud_image_import":                        ReplaceGroupWords("ecs", 0),
	"alicloud_image_share_permission":              ReplaceGroupWords("ecs", 0),
	"alicloud_launch_template":                     ReplaceGroupWords("ecs", 0),
	"alicloud_auto_provisioning_group":             ReplaceGroupWords("ecs", 0),
	"alicloud_instance_set":                        ReplaceGroupWords("ecs", 0),
	"alicloud_network_interface":                   ReplaceGroupWords("ecs", 0),
	"alicloud_network_interface_attachment":        ReplaceGroupWords("ecs", 0),
	"alicloud_network_interface_permission":        ReplaceGroupWords("ecs", 0),
	"alicloud_deployment_set":                      ReplaceGroupWords("ecs", 0),
	"alicloud_command":                             ReplaceGroupWords("ecs", 0),
	"alicloud_invocation":                          ReplaceGroupWords("ecs", 0),
	"alicloud_activation":                          ReplaceGroupWords("ecs", 0),
	"alicloud_reserved_instance":                   ReplaceGroupWords("ecs", 0),
	"alicloud_ecs_auto_snapshot_policy":            ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_auto_snapshot_policy_attachment": ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_capacity_reservation":            ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_command":                         ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_dedicated_host":                  ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_dedicated_host_cluster":          ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_deployment_set":                  ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_disk":                            ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_disk_attachment":                 ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_elasticity_assurance":            ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_hpc_cluster":                     ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_image_component":                 ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_image_pipeline":                  ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_image_pipeline_execution":        ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_instance_set":                    ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_invocation":                      ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_key_pair":                        ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_key_pair_attachment":             ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_launch_template":                 ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_network_interface":               ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_network_interface_attachment":    ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_network_interface_permission":    ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_prefix_list":                     ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_session_manager_status":          ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_snapshot":                        ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_snapshot_group":                  ReplaceGroupWords("ecs", 1),
	"alicloud_ecs_storage_capacity_unit":           ReplaceGroupWords("ecs", 1),

	// KMS - Key Management Service resources
	"alicloud_kms_alias":    ReplaceGroupWords("kms", 1),
	"alicloud_kms_instance": ReplaceGroupWords("kms", 1),
	"alicloud_kms_key":      ReplaceGroupWords("kms", 1),
	"alicloud_kms_secret":   ReplaceGroupWords("kms", 1),

	// MESSAGESERVICE - Message Service resources
	"alicloud_message_service_endpoint":     ReplaceGroupWords("messageservice", 2),
	"alicloud_message_service_endpoint_acl": ReplaceGroupWords("messageservice", 2),
	"alicloud_message_service_queue":        ReplaceGroupWords("messageservice", 2),
	"alicloud_message_service_subscription": ReplaceGroupWords("messageservice", 2),
	"alicloud_message_service_topic":        ReplaceGroupWords("messageservice", 2),

	// OSS - Object Storage Service resources
	"alicloud_oss_bucket":                            ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_object":                     ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_acl":                        ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_cors":                       ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_logging":                    ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_website":                    ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_referer":                    ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_policy":                     ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_server_side_encryption":     ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_versioning":                 ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_request_payment":            ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_transfer_acceleration":      ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_worm":                       ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_cname":                      ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_cname_token":                ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_https_config":               ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_access_monitor":             ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_public_access_block":        ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_user_defined_log_fields":    ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_data_redundancy_transition": ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_replication":                ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_meta_query":                 ReplaceGroupWords("oss", 1),
	"alicloud_oss_bucket_style":                      ReplaceGroupWords("oss", 1),
	"alicloud_oss_access_point":                      ReplaceGroupWords("oss", 1),
	"alicloud_oss_account_public_access_block":       ReplaceGroupWords("oss", 1),

	// POLARDB - PolarDB Database resources
	"alicloud_polardb_account":                 ReplaceGroupWords("polardb", 1),
	"alicloud_polardb_account_privilege":       ReplaceGroupWords("polardb", 1),
	"alicloud_polardb_backup_policy":           ReplaceGroupWords("polardb", 1),
	"alicloud_polardb_cluster":                 ReplaceGroupWords("polardb", 1),
	"alicloud_polardb_cluster_endpoint":        ReplaceGroupWords("polardb", 1),
	"alicloud_polardb_database":                ReplaceGroupWords("polardb", 1),
	"alicloud_polardb_endpoint":                ReplaceGroupWords("polardb", 1),
	"alicloud_polardb_endpoint_address":        ReplaceGroupWords("polardb", 1),
	"alicloud_polardb_global_database_network": ReplaceGroupWords("polardb", 1),
	"alicloud_polardb_parameter_group":         ReplaceGroupWords("polardb", 1),
	"alicloud_polardb_primary_endpoint":        ReplaceGroupWords("polardb", 1),

	// PRIVATELINK - PrivateLink resources
	"alicloud_privatelink_vpc_endpoint":                  ReplaceGroupWords("privatelink", 1),
	"alicloud_privatelink_vpc_endpoint_connection":       ReplaceGroupWords("privatelink", 1),
	"alicloud_privatelink_vpc_endpoint_service":          ReplaceGroupWords("privatelink", 1),
	"alicloud_privatelink_vpc_endpoint_service_resource": ReplaceGroupWords("privatelink", 1),
	"alicloud_privatelink_vpc_endpoint_service_user":     ReplaceGroupWords("privatelink", 1),
	"alicloud_privatelink_vpc_endpoint_zone":             ReplaceGroupWords("privatelink", 1),

	// QUOTAS - Quota Management resources
	"alicloud_quotas_quota_alarm":           ReplaceGroupWords("quotas", 1),
	"alicloud_quotas_quota_application":     ReplaceGroupWords("quotas", 1),
	"alicloud_quotas_template_applications": ReplaceGroupWords("quotas", 1),
	"alicloud_quotas_template_quota":        ReplaceGroupWords("quotas", 1),
	"alicloud_quotas_template_service":      ReplaceGroupWords("quotas", 1),

	// RAM - Resource Access Management resources
	"alicloud_ram_user":                    ReplaceGroupWords("ram", 1),
	"alicloud_ram_group":                   ReplaceGroupWords("ram", 1),
	"alicloud_ram_role":                    ReplaceGroupWords("ram", 1),
	"alicloud_ram_policy":                  ReplaceGroupWords("ram", 1),
	"alicloud_ram_user_policy_attachment":  ReplaceGroupWords("ram", 1),
	"alicloud_ram_group_policy_attachment": ReplaceGroupWords("ram", 1),
	"alicloud_ram_role_policy_attachment":  ReplaceGroupWords("ram", 1),
	"alicloud_ram_group_membership":        ReplaceGroupWords("ram", 1),
	"alicloud_ram_user_group_attachment":   ReplaceGroupWords("ram", 1),
	"alicloud_ram_access_key":              ReplaceGroupWords("ram", 1),
	"alicloud_ram_login_profile":           ReplaceGroupWords("ram", 1),
	"alicloud_ram_account_alias":           ReplaceGroupWords("ram", 1),
	"alicloud_ram_account_password_policy": ReplaceGroupWords("ram", 1),
	"alicloud_ram_password_policy":         ReplaceGroupWords("ram", 1),
	"alicloud_ram_security_preference":     ReplaceGroupWords("ram", 1),
	"alicloud_ram_saml_provider":           ReplaceGroupWords("ram", 1),

	// SLB - Server Load Balancer resources
	"alicloud_slb_load_balancer": ReplaceGroupWords("slb", 1),

	// TAIR - Redis Tair resources
	"alicloud_kvstore_account":          ReplaceGroupWords("tair", 1),
	"alicloud_kvstore_audit_log_config": ReplaceGroupWords("tair", 1),
	"alicloud_kvstore_connection":       ReplaceGroupWords("tair", 1),
	"alicloud_kvstore_instance":         ReplaceGroupWords("tair", 1),
	"alicloud_redis_tair_instance":      ReplaceGroupWords("tair", 2),

	// VPC - Virtual Private Cloud resources
	"alicloud_route_table": ReplaceGroupWords("vpc", 0),
	"alicloud_vpc":         ReplaceGroupWords("vpc", 0),
	"alicloud_vswitch":     ReplaceGroupWords("vpc", 0),
}

// GroupKindOverrides overrides the group and kind of the resource if it matches
// any entry in the GroupMap.
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if f, ok := GroupMap[r.Name]; ok {
			r.ShortGroup, r.Kind = f(r.Name)
		}
	}
}
