/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud"
)

// TestSchemaSourcesAgree guards the two independent sources of truth for the
// provider schema against drifting apart.
//
// Code generation reads config/schema.json, produced by `terraform providers
// schema` against the registry release pinned in the Makefile
// (TERRAFORM_PROVIDER_VERSION). The runtime schema comes from the upstream Go
// module pinned in go.mod, which for a prerelease is an opaque pseudo-version.
// Nothing couples the two, and the identifiers look nothing alike, so a bump to
// one and not the other is easy to miss.
//
// The loud failure mode is upjet's NewProvider, which panics at controller
// startup if a resource in the plugin SDK include list has no entry in the Go
// ResourcesMap. The quiet one is a resource present in both whose fields differ,
// generating a CRD that does not match what executes. This test catches the
// first directly and gives the second somewhere obvious to grow.
func TestSchemaSourcesAgree(t *testing.T) {
	goSchema := alicloud.Provider().ResourcesMap

	jsonProvider, err := getProviderSchema(providerSchema)
	if err != nil {
		t.Fatalf("cannot read the embedded JSON schema: %v", err)
	}

	for name := range terraformPluginSDKExternalNameConfigs {
		if _, ok := goSchema[name]; !ok {
			t.Errorf("resource %q is configured for plugin SDK reconciliation but is absent from the "+
				"upstream Go schema (go.mod); upjet will panic at controller startup. Are go.mod and "+
				"the Makefile's TERRAFORM_PROVIDER_VERSION pointing at the same upstream commit?", name)
		}
		if _, ok := jsonProvider.ResourcesMap[name]; !ok {
			t.Errorf("resource %q is configured but is absent from config/schema.json; regenerate it "+
				"with `make generate` against the pinned TERRAFORM_PROVIDER_VERSION", name)
		}
	}
}
