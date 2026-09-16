package ram

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_ram_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ram"
		r.ShortGroup = "ram"

		// Name has been deprecated in favor of groupName
	})

	p.AddResourceConfigurator("alicloud_ram_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ram"
		r.ShortGroup = "ram"

		// Document has been deprecated in favor of policyDocument
		// Name has been deprecated in favor of policyName
		// Statement has been deprecated
		// Version has been deprecated
	})

	p.AddResourceConfigurator("alicloud_ram_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ram"
		r.ShortGroup = "ram"

	})
}
