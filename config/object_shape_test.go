// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/crossplane/upjet/pkg/config"

	providerconfig "github.com/crossplane-contrib/provider-alibabacloud/config"
)

// Terraform provider 2.0 changed these fields from maps to lists of one object.
// They must stay objects, the shape 1.4 stored, or the provider cannot decode
// existing resources and stops.
var objectShapeFields = map[string][]string{
	"alicloud_cs_edge_kubernetes":    {"certificate_authority", "connections", "runtime"},
	"alicloud_cs_kubernetes":         {"certificate_authority", "connections", "runtime"},
	"alicloud_cs_managed_kubernetes": {"certificate_authority", "connections"},
	"alicloud_cr_repo":               {"domain_list"},
}

func TestKeepObjectShape(t *testing.T) {
	for _, generation := range []bool{true, false} {
		p, err := providerconfig.GetProvider(context.Background(), generation)
		if err != nil {
			t.Fatalf("cannot build the provider configuration (generation %t): %v", generation, err)
		}
		for resource, fields := range objectShapeFields {
			r, ok := p.Resources[resource]
			if !ok {
				t.Fatalf("resource %s is not configured", resource)
			}
			for _, f := range fields {
				if !r.SchemaElementOptions.EmbeddedObject(f) {
					t.Errorf("%s: %s is not generated as an object (generation %t)", resource, f, generation)
				}
			}
		}
	}
}

func TestKeepObjectShapeConversion(t *testing.T) {
	p, err := providerconfig.GetProvider(context.Background(), false)
	if err != nil {
		t.Fatalf("cannot build the provider configuration: %v", err)
	}
	r := p.Resources["alicloud_cs_kubernetes"]
	runtime := map[string]any{"name": "containerd", "version": "1.6.28"}
	ca := map[string]any{"cluster_cert": "ca", "client_cert": "cert", "client_key": "key"}

	fromTF, err := r.ApplyTFConversions(map[string]any{
		"runtime":               []any{runtime},
		"certificate_authority": []any{ca},
		"name":                  "c1",
	}, config.FromTerraform)
	if err != nil {
		t.Fatalf("cannot convert from Terraform: %v", err)
	}
	want := map[string]any{"runtime": runtime, "certificate_authority": ca, "name": "c1"}
	if !reflect.DeepEqual(fromTF, want) {
		t.Errorf("from Terraform:\n got %v\nwant %v", fromTF, want)
	}

	toTF, err := r.ApplyTFConversions(map[string]any{"runtime": runtime}, config.ToTerraform)
	if err != nil {
		t.Fatalf("cannot convert to Terraform: %v", err)
	}
	if want := map[string]any{"runtime": []any{runtime}}; !reflect.DeepEqual(toTF, want) {
		t.Errorf("to Terraform:\n got %v\nwant %v", toTF, want)
	}

	if _, err := r.ApplyTFConversions(map[string]any{"runtime": []any{runtime, runtime}}, config.FromTerraform); err == nil {
		t.Error("a list of two runtimes converted to an object; want an error")
	}
}
