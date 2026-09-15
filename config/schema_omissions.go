/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// generationOnlySchemaOmissions lists, per Terraform resource, the top-level
// schema fields that must not appear in the generated CRD.
//
// These MUST only ever be applied to the code-generation provider. Under the
// no-fork architecture upjet assigns Resource.TerraformResource directly from
// the live *schema.Provider returned by alicloud.Provider(), without copying
// it (see upjet pkg/config/provider.go, the TerraformPluginSDK branch of
// NewProvider). Deleting from that schema at runtime therefore removes fields
// from the very map the upstream CRUD functions execute against, which breaks
// them in ways the Terraform provider has no reason to defend against:
//
//   - alicloud_oss_bucket Create does oss.ACLType(d.Get("acl").(string)); with
//     "acl" absent from the schema d.Get returns nil and the assertion panics.
//   - alicloud_oss_bucket Read error-checks d.Set("policy", ...), and d.Set on
//     a field absent from the schema returns an error, so Observe fails
//     permanently. alicloud_slb_acl does the same with "entry_list".
//   - alicloud_slb_listener Create/Update read d.Get("ssl_certificate_id") when
//     an HTTPS listener has no server_certificate_id, panicking instead of
//     returning the intended validation error.
//
// Applying them during generation only reproduces the pre-no-fork behaviour
// exactly: the field stays in the runtime schema, the CRD never sets it, and
// d.Get returns the zero value rather than nil.
var generationOnlySchemaOmissions = map[string][]string{
	"alicloud_cs_kubernetes_node_pool":             {"security_group_id"},
	"alicloud_alb_acl":                             {"acl_entries"},
	"alicloud_alb_listener":                        {"acl_config", "xforwarded_for_config"},
	"alicloud_alidns_domain_group":                 {"group_name"},
	"alicloud_cr_ee_sync_rule":                     {"name"},
	"alicloud_ecs_auto_snapshot_policy":            {"name"},
	"alicloud_ecs_auto_snapshot_policy_attachment": {"name"},
	"alicloud_ecs_deployment_set":                  {"domain", "granularity"},
	"alicloud_ecs_disk":                            {"availability_zone", "name"},
	"alicloud_ecs_key_pair":                        {"key_name"},
	"alicloud_ecs_key_pair_attachment":             {"key_name"},
	"alicloud_ecs_launch_template":                 {"name", "system_disk_category", "system_disk_description", "system_disk_name", "system_disk_size", "userdata"},
	"alicloud_ecs_network_interface":               {"name", "security_groups", "private_ips", "private_ip", "private_ips_count"},
	"alicloud_ecs_snapshot":                        {"instant_access", "instant_access_retention_days", "name"},
	"alicloud_image":                               {"name"},
	"alicloud_image_copy":                          {"name"},
	"alicloud_instance":                            {"allocate_public_ip", "internet_max_bandwidth_in", "io_optimized", "subnet_id"},
	"alicloud_reserved_instance":                   {"name"},
	"alicloud_security_group":                      {"name", "inner_access"},
	"alicloud_kms_key":                             {"deletion_window_in_days", "is_enabled", "key_state"},
	"alicloud_message_service_topic":               {"logging_enabled"},
	"alicloud_oss_bucket":                          {"acl", "logging_isenable", "referer_config", "policy"},
	"alicloud_ram_group":                           {"name"},
	"alicloud_ram_policy":                          {"document", "name", "statement", "version"},
	"alicloud_ram_role":                            {"ram_users", "services", "version"},
	"alicloud_slb_acl":                             {"entry_list"},
	"alicloud_slb_load_balancer":                   {"name", "specification"},
	"alicloud_slb_listener":                        {"acl_id", "ssl_certificate_id"},
	"alicloud_kvstore_instance":                    {"availability_zone", "connection_string_prefix", "enable_public", "instance_charge_type", "instance_name", "node_type", "parameters"},
	"alicloud_route_table":                         {"name"},
	"alicloud_vswitch":                             {"name", "availability_zone"},
}

// addGenerationOnlySchemaOmissions registers the schema omissions above. It is
// called only when building the provider configuration for code generation;
// see GetProvider.
func addGenerationOnlySchemaOmissions(pc *config.Provider) {
	for name, fields := range generationOnlySchemaOmissions {
		pc.AddResourceConfigurator(name, func(r *config.Resource) {
			for _, f := range fields {
				delete(r.TerraformResource.Schema, f)
			}
		})
	}
}
