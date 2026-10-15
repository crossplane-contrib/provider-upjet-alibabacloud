// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package ram_test

import (
	"context"
	"strings"
	"testing"

	providerconfig "github.com/crossplane-contrib/provider-alibabacloud/config"
)

// TestReferenceExtractorsPointAtLiveFields guards against a reference
// extracting a field that a resource configurator has deleted from the schema.
//
// Role, Group and SecurityPreference drop deprecated aliases such as "name",
// so an extractor still reading "name" resolves to an empty value: the
// reference never resolves, and the only symptom is
// "referenced field was empty" at reconcile time. That is what happened to
// RolePolicyAttachment.roleName and to the two Group references after the
// aliases were removed.
func TestReferenceExtractorsPointAtLiveFields(t *testing.T) {
	p, err := providerconfig.GetProvider(context.Background(), true)
	if err != nil {
		t.Fatalf("cannot build the generation provider configuration: %v", err)
	}

	for _, tc := range []struct {
		resource string
		field    string
		want     string
	}{
		{"alicloud_ram_role_policy_attachment", "role_name", "role_name"},
		{"alicloud_ram_group_membership", "group_name", "group_name"},
		{"alicloud_ram_group_policy_attachment", "group_name", "group_name"},
	} {
		t.Run(tc.resource+"/"+tc.field, func(t *testing.T) {
			r, ok := p.Resources[tc.resource]
			if !ok {
				t.Fatalf("%s is not configured", tc.resource)
			}
			ref, ok := r.References[tc.field]
			if !ok {
				t.Fatalf("%s has no reference on %s", tc.resource, tc.field)
			}
			if !strings.Contains(ref.Extractor, `"`+tc.want+`"`) {
				t.Errorf("extractor for %s.%s is %q, expected it to read %q",
					tc.resource, tc.field, ref.Extractor, tc.want)
			}
		})
	}

	// The aliases those extractors used to read are gone; assert that, so the
	// test fails loudly if a future change puts them back and makes the
	// extractors above look unnecessary.
	for _, tc := range []struct{ resource, gone string }{
		{"alicloud_ram_role", "name"},
		{"alicloud_ram_role", "document"},
		{"alicloud_ram_group", "name"},
	} {
		t.Run(tc.resource+"/no-"+tc.gone, func(t *testing.T) {
			r, ok := p.Resources[tc.resource]
			if !ok {
				t.Fatalf("%s is not configured", tc.resource)
			}
			if _, ok := r.TerraformResource.Schema[tc.gone]; ok {
				t.Errorf("%s still exposes the deprecated %q", tc.resource, tc.gone)
			}
		})
	}
}

// TestNoReferenceExtractsADeletedField sweeps every configured reference and
// fails if its extractor reads a field the target resource no longer has. This
// is the general form of the bug above: it catches a new deletion breaking an
// existing reference, which is otherwise invisible until a reconcile.
func TestNoReferenceExtractsADeletedField(t *testing.T) {
	p, err := providerconfig.GetProvider(context.Background(), true)
	if err != nil {
		t.Fatalf("cannot build the generation provider configuration: %v", err)
	}
	for name, r := range p.Resources {
		for field, ref := range r.References {
			param := extractedParam(ref.Extractor)
			if param == "" || ref.TerraformName == "" {
				continue
			}
			target, ok := p.Resources[ref.TerraformName]
			if !ok || target.TerraformResource == nil {
				continue
			}
			if _, ok := target.TerraformResource.Schema[param]; !ok {
				t.Errorf("%s.%s extracts %q from %s, which has no such field",
					name, field, param, ref.TerraformName)
			}
		}
	}
}

// extractedParam returns the field name an ExtractParamPath extractor reads,
// or "" for any other extractor.
func extractedParam(extractor string) string {
	_, rest, found := strings.Cut(extractor, `ExtractParamPath("`)
	if !found {
		return ""
	}
	param, _, found := strings.Cut(rest, `"`)
	if !found {
		return ""
	}
	return param
}
