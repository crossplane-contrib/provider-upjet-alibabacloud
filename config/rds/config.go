// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package rds

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_db_instance", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathIdExtractor,
		}
		// The deprecated singular form is dropped during code generation (see
		// generationOnlySchemaOmissions in config/schema_omissions.go) for two
		// reasons: its generated selector fields collide with the ones
		// KnownReferences() produces for security_group_ids, and upstream
		// declares ConflictsWith: ["security_group_ids"] while both are
		// Optional+Computed, so late-init would populate the pair and every
		// plan would then error.
		// Neither secret is flagged Sensitive upstream, so upjet would emit
		// both as plain strings in spec and status.atProvider. Marked
		// sensitive during code generation instead; see
		// generationOnlySensitiveFields in config/schema_omissions.go, so
		// they become secret references.
	})
	p.AddResourceConfigurator("alicloud_db_readonly_instance", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["master_db_instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathIdExtractor,
		}
		// The scraped registry example contributes an
		// instance_storage -> alicloud_db_instance reference, but the generated
		// resolver reads it with ExtractParamPath("instance_storage", false),
		// which uses paved.GetString. instance_storage is a number, so it
		// always resolves to "" and reference resolution fails with
		// "referenced field was empty". Supply the value literally instead.
		delete(r.References, "instance_storage")
		// The same example also contributes instance_type -> alicloud_db_instance,
		// which resolves to the primary's class. A read-only instance takes a
		// read-only class (mysqlro.*), and the primary's is rejected with
		// InvalidSaleComponentFault, so the reference can only ever produce an
		// invalid value. Supply the class literally instead.
		delete(r.References, "instance_type")
		// Also not Sensitive upstream; same reasoning as alicloud_db_instance,
		// handled the same way in generationOnlySensitiveFields.
	})
	// NOTE: alicloud_db_account and alicloud_db_account_privilege (the
	// legacy account resources, superseded by alicloud_rds_account) are
	// intentionally NOT configured here and are excluded from the provider
	// via config/external_name.go (they carry deprecated ConflictsWith
	// field aliases and collide with alicloud_rds_account on kind Account).
	// Use alicloud_rds_account instead.
	p.AddResourceConfigurator("alicloud_db_backup_policy", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
		// These five aliases have no ConflictsWith upstream, but Read sets both
		// the alias and its canonical field, so late-init would copy the
		// aliases into spec alongside the canonical values. Dropped during
		// code generation; see generationOnlySchemaOmissions in
		// config/schema_omissions.go, so only the canonical field is managed.
	})
	p.AddResourceConfigurator("alicloud_db_connection", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_db_database", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
		// "name" is a deprecated, computed alias of "data_base_name"; both
		// being computed causes TF refresh to fail with a ConflictsWith error
		// ("data_base_name" conflicts with "name"). Dropped during code
		// generation; see generationOnlySchemaOmissions in
		// config/schema_omissions.go. With the alias gone there is no other
		// way to name a database, and TF create errors if neither is set, so
		// mark it required at the CRD level rather than letting it fail at
		// apply time. MarkAsRequired records this without touching the live
		// schema, unlike setting Required/Optional/Computed directly, which
		// would mutate the schema the upstream CRUD functions execute
		// against under no-fork.
		r.MarkAsRequired("data_base_name")
	})
	p.AddResourceConfigurator("alicloud_db_read_write_splitting_connection", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
	})

	p.AddResourceConfigurator("alicloud_rds_account", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["db_instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
		// Deprecated, computed aliases that collide with their non-deprecated
		// replacements on TF refresh (ConflictsWith). Dropped during code
		// generation; see generationOnlySchemaOmissions in
		// config/schema_omissions.go. Same reasoning as data_base_name
		// above: the alias is gone, so account_name is effectively
		// mandatory; MarkAsRequired records that without touching the live
		// schema.
		r.MarkAsRequired("account_name")
	})
	p.AddResourceConfigurator("alicloud_rds_backup", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["db_instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_rds_db_instance_endpoint", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["db_instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_rds_db_instance_endpoint_address", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["db_instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
		// The endpoint's own id is composite ("<DBInstanceId>:<EndpointId>"),
		// but CreateDBInstanceEndpointAddress wants the bare endpoint id, which
		// the endpoint exposes as the computed db_instance_endpoint_id. Read
		// that field rather than status.atProvider.id.
		r.References["db_instance_endpoint_id"] = config.Reference{
			TerraformName: "alicloud_rds_db_instance_endpoint",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("db_instance_endpoint_id",true)`,
		}
	})
	p.AddResourceConfigurator("alicloud_rds_db_node", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["db_instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_rds_db_proxy", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_rds_db_proxy_public", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["db_instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_rds_instance_cross_backup_policy", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_db_instance",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_rds_parameter_group", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
	})
	p.AddResourceConfigurator("alicloud_rds_service_linked_role", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
	})
	p.AddResourceConfigurator("alicloud_rds_whitelist_template", func(r *config.Resource) {
		r.ShortGroup = string(common.RDS)
	})
}
