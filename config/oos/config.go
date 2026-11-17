package oos

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_oos_application", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_parameter", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_template", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_application_group", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_default_patch_baseline", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_execution", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
		// Terraform provider 1.293.0 added "tags" here as tagsSchemaForceNew.
		// AddExternalTagsField would otherwise write the crossplane-* tags into
		// spec, and on an Execution created before that field existed the plan
		// wants to add them, which is a replacement. Upjet renders
		// prevent_destroy for resources that are not being deleted, so the
		// apply fails on every reconcile instead. Users can still set tags
		// themselves; we just do not inject them.
		r.InitializerFns = nil
	})
	p.AddResourceConfigurator("alicloud_oos_patch_baseline", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_secret_parameter", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_service_setting", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_state_configuration", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
}
