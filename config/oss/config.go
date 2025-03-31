package oss

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_oss", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "oss"
		r.ShortGroup = "oss"
	})
	p.AddResourceConfigurator("alicloud_oss_bucket", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "oss"
		r.ShortGroup = "oss"
		delete(r.TerraformResource.Schema, "acl")
		delete(r.TerraformResource.Schema, "logging_isenable")
		delete(r.TerraformResource.Schema, "referer_config")
		delete(r.TerraformResource.Schema, "policy")
	})
}
