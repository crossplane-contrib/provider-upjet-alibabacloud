// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package sls

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures the Simple Log Service (SLS / Log) resources by adding
// custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_log_project", func(r *config.Resource) {
		r.ShortGroup = string(common.SLS)
		r.Kind = "Project"
		// "name" is a deprecated alias of "project_name"; dropped during code
		// generation (see generationOnlySchemaOmissions in
		// config/schema_omissions.go) so only the canonical field is
		// generated. Upstream leaves project_name Optional+Computed only
		// because the deprecated "name" alias could supply it instead. With
		// that alias dropped there is no other way to name a project, so mark
		// it Required: otherwise a Project with just "region" set passes CRD
		// admission and only fails later at plan time. MarkAsRequired records
		// this without touching the live schema, unlike setting
		// Required/Optional/Computed directly, which would mutate the schema
		// the upstream CRUD functions execute against under no-fork.
		r.MarkAsRequired("project_name")
	})

	p.AddResourceConfigurator("alicloud_log_store", func(r *config.Resource) {
		r.ShortGroup = string(common.SLS)
		r.Kind = "Store"
		// "name"/"project" are deprecated aliases of "logstore_name"/"project_name".
		// The scraped registry example still wires the deprecated "project" to
		// alicloud_log_project's "name", which is deleted there as well. With the
		// field gone no resolver is generated for it, so drop the dead reference.
		delete(r.References, "project")
		// Same reasoning as Project.project_name above: upstream documents
		// logstore_name as "one of logstore_name, name", and the "name" alias is
		// gone, so it is effectively mandatory; MarkAsRequired records that
		// without touching the live schema.
		r.MarkAsRequired("logstore_name")
		// project_name deliberately stays optional: every example supplies it
		// through projectNameRef/projectNameSelector, and a Required field
		// cannot be omitted in favour of a reference.
		r.References["project_name"] = config.Reference{
			TerraformName: "alicloud_log_project",
			Extractor:     common.PathIdExtractor,
		}
	})

	p.AddResourceConfigurator("alicloud_log_store_index", func(r *config.Resource) {
		r.ShortGroup = string(common.SLS)
		r.Kind = "StoreIndex"
		r.References["project"] = config.Reference{
			TerraformName: "alicloud_log_project",
			Extractor:     common.PathIdExtractor,
		}
		r.References["logstore"] = config.Reference{
			TerraformName: "alicloud_log_store",
			Extractor:     common.PathLogStoreNameExtractor,
		}
	})

	p.AddResourceConfigurator("alicloud_log_machine_group", func(r *config.Resource) {
		r.ShortGroup = string(common.SLS)
		r.Kind = "MachineGroup"
		r.References["project"] = config.Reference{
			TerraformName: "alicloud_log_project",
			Extractor:     common.PathIdExtractor,
		}
	})

	p.AddResourceConfigurator("alicloud_logtail_config", func(r *config.Resource) {
		r.ShortGroup = string(common.SLS)
		r.Kind = "LogtailConfig"
		r.References["project"] = config.Reference{
			TerraformName: "alicloud_log_project",
			Extractor:     common.PathIdExtractor,
		}
		r.References["logstore"] = config.Reference{
			TerraformName: "alicloud_log_store",
			Extractor:     common.PathLogStoreNameExtractor,
		}
	})

	p.AddResourceConfigurator("alicloud_logtail_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.SLS)
		r.Kind = "LogtailAttachment"
		r.References["project"] = config.Reference{
			TerraformName: "alicloud_log_project",
			Extractor:     common.PathIdExtractor,
		}
		r.References["logtail_config_name"] = config.Reference{
			TerraformName: "alicloud_logtail_config",
			Extractor:     common.PathNameExtractor,
		}
		r.References["machine_group_name"] = config.Reference{
			TerraformName: "alicloud_log_machine_group",
			Extractor:     common.PathNameExtractor,
		}
	})
}
