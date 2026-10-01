// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package sls

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures the Simple Log Service (SLS / Log) resources by adding
// custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_log_project", func(r *config.Resource) {
		r.ShortGroup = string(common.SLS)
		r.Kind = "Project"
		// "name" is a deprecated alias of "project_name"; drop it so only the
		// canonical field is generated.
		delete(r.TerraformResource.Schema, "name")
	})

	p.AddResourceConfigurator("alicloud_log_store", func(r *config.Resource) {
		r.ShortGroup = string(common.SLS)
		r.Kind = "Store"
		// "name"/"project" are deprecated aliases of "logstore_name"/"project_name".
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "project")
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
