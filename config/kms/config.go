package kms

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_kms_key", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "kms"
		r.ShortGroup = "kms"

		delete(r.TerraformResource.Schema, "deletion_window_in_days")
		delete(r.TerraformResource.Schema, "is_enabled")
		delete(r.TerraformResource.Schema, "key_state")
	})
}
