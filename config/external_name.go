/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda

	// ECS
	"alicloud_security_group": config.IdentifierFromProvider,

	// KMS
	"alicloud_kms_alias":                    config.IdentifierFromProvider,
	"alicloud_kms_application_access_point": config.IdentifierFromProvider,
	"alicloud_kms_ciphertext":               config.IdentifierFromProvider,
	"alicloud_kms_client_key":               config.IdentifierFromProvider,
	"alicloud_kms_instance":                 config.IdentifierFromProvider,
	"alicloud_kms_key":                      config.IdentifierFromProvider,
	"alicloud_kms_key_version":              config.IdentifierFromProvider,
	"alicloud_kms_network_rule":             config.IdentifierFromProvider,
	"alicloud_kms_policy":                   config.IdentifierFromProvider,
	"alicloud_kms_secret":                   config.IdentifierFromProvider,

	// Quotas
	"alicloud_quotas_quota_alarm":           config.IdentifierFromProvider,
	"alicloud_quotas_quota_application":     config.IdentifierFromProvider,
	"alicloud_quotas_template_applications": config.IdentifierFromProvider,
	"alicloud_quotas_template_quota":        config.IdentifierFromProvider,
	"alicloud_quotas_template_service":      config.IdentifierFromProvider,

	// RAM
	"alicloud_ram_access_key":              config.IdentifierFromProvider,
	"alicloud_ram_account_alias":           config.IdentifierFromProvider,
	"alicloud_ram_account_password_policy": config.IdentifierFromProvider,
	"alicloud_ram_group":                   config.IdentifierFromProvider,
	"alicloud_ram_group_membership":        config.IdentifierFromProvider,
	"alicloud_ram_group_policy_attachment": config.IdentifierFromProvider,
	"alicloud_ram_login_profile":           config.IdentifierFromProvider,
	"alicloud_ram_policy":                  config.IdentifierFromProvider,
	"alicloud_ram_role":                    config.IdentifierFromProvider,
	"alicloud_ram_role_policy_attachment":  config.IdentifierFromProvider,
	"alicloud_ram_saml_provider":           config.IdentifierFromProvider,
	"alicloud_ram_security_preference":     config.IdentifierFromProvider,
	"alicloud_ram_user":                    config.IdentifierFromProvider,
	"alicloud_ram_user_policy_attachment":  config.IdentifierFromProvider,

	// VPC
	"alicloud_vpc": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
