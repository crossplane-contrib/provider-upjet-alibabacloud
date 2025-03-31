package ram

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_ram_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ram"
		r.ShortGroup = "ram"

		// Document has been deprecated in favor of policyDocument
		delete(r.TerraformResource.Schema, "document")
		// Name has been deprecated in favor of policyName
		delete(r.TerraformResource.Schema, "name")
		// Statement has been deprecated in favor of document
		delete(r.TerraformResource.Schema, "statement")
		// Version has been derecated in favor of document
		delete(r.TerraformResource.Schema, "version")
	})

	p.AddResourceConfigurator("alicloud_ram_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ram"
		r.ShortGroup = "ram"

		delete(r.TerraformResource.Schema, "ram_users")
		delete(r.TerraformResource.Schema, "services")
		delete(r.TerraformResource.Schema, "version")
	})
}
