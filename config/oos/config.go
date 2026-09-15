package oos

import (
	"github.com/crossplane/upjet/v2/pkg/config"

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
