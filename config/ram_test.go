// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config_test

import (
	"testing"

	"github.com/crossplane/upjet/pkg/registry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	providerconfig "github.com/crossplane-contrib/provider-alibabacloud/config"
)

func TestRAMDeprecatedFields(t *testing.T) {
	p := providerconfig.GetProvider()
	for _, tc := range []struct {
		resource          string
		removed, retained []string
	}{
		{"alicloud_ram_role", []string{"name", "document", "ram_users", "services", "version"}, []string{"role_name", "assume_role_policy_document"}},
		{"alicloud_ram_security_preference", []string{"enforce_mfa_for_login"}, []string{"mfa_operation_for_login"}},
	} {
		t.Run(tc.resource, func(t *testing.T) {
			r, ok := p.Resources[tc.resource]
			if !ok {
				t.Fatalf("resource %s is not configured", tc.resource)
			}
			for _, key := range tc.removed {
				if _, ok := r.TerraformResource.Schema[key]; ok {
					t.Errorf("deprecated field %q remains in schema", key)
				}
			}
			for _, key := range tc.retained {
				if _, ok := r.TerraformResource.Schema[key]; !ok {
					t.Errorf("canonical field %q is missing", key)
				}
			}
		})
	}
}

func TestRAMSecurityPreferenceDocumentation(t *testing.T) {
	for _, tc := range []struct {
		name     string
		metadata *registry.Resource
		wantDoc  bool
	}{
		{"MissingResource", nil, false},
		{"NilArgumentDocs", &registry.Resource{}, false},
		{"MissingKey", &registry.Resource{ArgumentDocs: map[string]string{"other": "unchanged"}}, false},
		{"ExistingKey", &registry.Resource{ArgumentDocs: map[string]string{"mfa_operation_for_login": "outdated", "other": "unchanged"}}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := providerconfig.GetProvider()
			r := p.Resources["alicloud_ram_security_preference"]
			r.MetaResource = tc.metadata
			// Removing the unsupported field must still run when documentation is absent.
			r.TerraformResource.Schema["enforce_mfa_for_login"] = &schema.Schema{Type: schema.TypeBool, Optional: true}
			p.ConfigureResources()
			if _, ok := r.TerraformResource.Schema["enforce_mfa_for_login"]; ok {
				t.Fatal("unsupported field remains")
			}
			if r.MetaResource == nil {
				return
			}
			got, exists := r.MetaResource.ArgumentDocs["mfa_operation_for_login"]
			if exists != tc.wantDoc {
				t.Fatalf("documentation present = %t, want %t", exists, tc.wantDoc)
			}
			if tc.wantDoc && got != "The login MFA policy for RAM users." {
				t.Errorf("unexpected documentation: %q", got)
			}
			if other, ok := r.MetaResource.ArgumentDocs["other"]; ok && other != "unchanged" {
				t.Error("unrelated documentation changed")
			}
		})
	}
}
