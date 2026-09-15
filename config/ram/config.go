package ram

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_ram_security_preference", func(r *config.Resource) {
		// The provider no longer reads or writes this deprecated field.
		// Use mfa_operation_for_login to configure the login MFA policy.
		// Upstream documentation still claims the removed field is effective.
		if r.MetaResource == nil {
			return
		}
		if _, ok := r.MetaResource.ArgumentDocs["mfa_operation_for_login"]; ok {
			r.MetaResource.ArgumentDocs["mfa_operation_for_login"] = "The login MFA policy for RAM users."
		}
	})

	p.AddResourceConfigurator("alicloud_ram_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ram"
		r.ShortGroup = "ram"

		// "name" has been deprecated in favor of "group_name" and is dropped
		// during code generation; see generationOnlySchemaOmissions in
		// config/schema_omissions.go.
	})

	p.AddResourceConfigurator("alicloud_ram_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ram"
		r.ShortGroup = "ram"

		// "document", "name", "statement" and "version" are all deprecated
		// aliases of "policy_document"/"policy_name" and dropped during code
		// generation; see generationOnlySchemaOmissions in
		// config/schema_omissions.go.
	})

	// alicloud_ram_group drops the deprecated "name" above, so references that
	// extract it from a Group resolve to an empty value. Point them at the
	// canonical field, as for Role below.
	p.AddResourceConfigurator("alicloud_ram_group_membership", func(r *config.Resource) {
		r.References["group_name"] = config.Reference{
			TerraformName: "alicloud_ram_group",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("group_name",false)`,
		}
	})

	p.AddResourceConfigurator("alicloud_ram_group_policy_attachment", func(r *config.Resource) {
		r.References["group_name"] = config.Reference{
			TerraformName: "alicloud_ram_group",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("group_name",false)`,
		}
	})

	p.AddResourceConfigurator("alicloud_ram_role_policy_attachment", func(r *config.Resource) {
		// Upstream documentation at the pinned provider version still points
		// this reference at the Role's deprecated "name", which the Role
		// configurator below removes from the schema. Left alone, roleNameRef
		// and roleNameSelector resolve to an empty value and the attachment
		// never reconciles. Extract the canonical field instead.
		r.References["role_name"] = config.Reference{
			TerraformName: "alicloud_ram_role",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("role_name",false)`,
		}
	})

	p.AddResourceConfigurator("alicloud_ram_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ram"
		r.ShortGroup = "ram"

	})
}
