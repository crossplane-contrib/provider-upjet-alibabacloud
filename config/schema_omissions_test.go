/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud"
)

// TestRuntimeSchemaIsPristine guards the no-fork invariant that the provider
// configuration built for runtime never mutates the upstream Terraform schema.
//
// upjet assigns Resource.TerraformResource straight from the live
// *schema.Provider without copying it, so any delete against
// r.TerraformResource.Schema outside code generation would strip fields from
// the map the upstream CRUD functions execute against. That breaks them at
// runtime only, where it is expensive to notice: alicloud_oss_bucket Create
// panics on d.Get("acl").(string), and alicloud_oss_bucket and
// alicloud_slb_acl Observe fail permanently on an error-checked d.Set.
func TestRuntimeSchemaIsPristine(t *testing.T) {
	pristine := alicloud.Provider()

	p, err := GetProvider(context.Background(), false)
	if err != nil {
		t.Fatalf("cannot build the runtime provider configuration: %v", err)
	}

	for name, r := range p.Resources {
		ref, ok := pristine.ResourcesMap[name]
		if !ok || ref.Schema == nil {
			continue
		}
		for field := range ref.Schema {
			if _, ok := r.TerraformResource.Schema[field]; !ok {
				t.Errorf("resource %q is missing field %q from its runtime Terraform schema; "+
					"schema fields may only be omitted from the code-generation provider "+
					"(see generationOnlySchemaOmissions)", name, field)
			}
		}
	}
}

// TestGenerationSchemaOmissionsApply is the other half: the omissions must
// still take effect for code generation, or fields would leak back into the
// generated CRDs.
func TestGenerationSchemaOmissionsApply(t *testing.T) {
	p, err := GetProvider(context.Background(), true)
	if err != nil {
		t.Fatalf("cannot build the generation provider configuration: %v", err)
	}

	applied := 0
	for name, fields := range generationOnlySchemaOmissions {
		r, ok := p.Resources[name]
		if !ok {
			t.Errorf("resource %q has schema omissions configured but is not a generated resource", name)
			continue
		}
		for _, f := range fields {
			if _, ok := r.TerraformResource.Schema[f]; ok {
				t.Errorf("resource %q still exposes omitted field %q to code generation", name, f)
				continue
			}
			applied++
		}
	}
	if applied == 0 {
		t.Fatal("no schema omissions were applied; the generation provider is not being configured")
	}
}
