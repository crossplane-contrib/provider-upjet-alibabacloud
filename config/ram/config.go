package ram

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_ram_security_preference", func(r *config.Resource) {
		// The provider no longer reads or writes this deprecated field.
		// Use mfa_operation_for_login to configure the login MFA policy.
		delete(r.TerraformResource.Schema, "enforce_mfa_for_login")
		// Upstream documentation still claims the removed field is effective.
		r.MetaResource.ArgumentDocs["mfa_operation_for_login"] = "The login MFA policy for RAM users."
	})

	p.AddResourceConfigurator("alicloud_ram_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ram"
		r.ShortGroup = "ram"

		// Name has been deprecated in favor of groupName
		delete(r.TerraformResource.Schema, "name")
	})

	p.AddResourceConfigurator("alicloud_ram_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ram"
		r.ShortGroup = "ram"

		// Document has been deprecated in favor of policyDocument
		delete(r.TerraformResource.Schema, "document")
		// Name has been deprecated in favor of policyName
		delete(r.TerraformResource.Schema, "name")
		// Statement has been deprecated
		delete(r.TerraformResource.Schema, "statement")
		// Version has been deprecated
		delete(r.TerraformResource.Schema, "version")
	})

	p.AddResourceConfigurator("alicloud_ram_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ram"
		r.ShortGroup = "ram"

		// Read returns both canonical fields and their deprecated aliases.
		// Exclude the aliases so late initialization cannot add conflicting
		// arguments to the configuration on the next refresh.
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "document")
		delete(r.TerraformResource.Schema, "ram_users")
		delete(r.TerraformResource.Schema, "services")
		delete(r.TerraformResource.Schema, "version")
	})
}
