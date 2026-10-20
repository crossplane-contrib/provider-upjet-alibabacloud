// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	providerconfig "github.com/crossplane-contrib/provider-alibabacloud/config"
)

// rdsInstanceParams is spec.forProvider of examples/rds/v1alpha1/account.yaml's
// Instance, in Terraform attribute names.
func rdsInstanceParams() map[string]any {
	return map[string]any{
		"region":                   "cn-beijing",
		"engine":                   "MySQL",
		"engine_version":           "8.0",
		"category":                 "Basic",
		"db_instance_storage_type": "cloud_essd",
		"instance_charge_type":     "Postpaid",
		"instance_name":            "crossplane-example",
		"instance_storage":         20,
		"instance_type":            "mysql.n2.medium.1",
		"zone_id":                  "cn-beijing-l",
		"security_ips":             []any{"10.0.0.0/8"},
		"vswitch_id":               "vsw-example",
	}
}

// TestUnknownComputedCountSuppression reproduces the RDS Instance import
// failure: with no cached state, upjet rebuilds the prior state from
// spec.forProvider, which leaves empty Optional+Computed sets out entirely,
// and the plugin SDK then reports their counts as NewComputed. For
// babelfish_config that diff is ForceNew, so Update refuses forever.
func TestUnknownComputedCountSuppression(t *testing.T) {
	p, err := providerconfig.GetProvider(context.Background(), false)
	if err != nil {
		t.Fatalf("cannot build the runtime provider configuration: %v", err)
	}
	r, ok := p.Resources["alicloud_db_instance"]
	if !ok {
		t.Fatal("alicloud_db_instance is not configured")
	}
	if r.TerraformCustomDiff == nil {
		t.Fatal("alicloud_db_instance has no TerraformCustomDiff; UnknownComputedCountSuppression is not wired in")
	}

	// The keys the uptest import run reported.
	counts := []string{"babelfish_config.#", "parameters.#", "security_group_ids.#"}

	cases := map[string]struct {
		id string
		// params is the prior spec.forProvider the state is rebuilt from;
		// config, when set, is the current one diffed against it.
		params      func() map[string]any
		config      func() map[string]any
		wantKept    []string
		wantDropped []string
	}{
		"ImportedInstance": {
			id:          "rm-example",
			params:      rdsInstanceParams,
			wantDropped: counts,
		},
		"Create": {
			params:   rdsInstanceParams,
			wantKept: counts,
		},
		"AddedSetIsKept": {
			id:     "rm-example",
			params: rdsInstanceParams,
			config: func() map[string]any {
				m := rdsInstanceParams()
				m["security_group_ids"] = []any{"sg-example"}
				return m
			},
			wantKept:    []string{"security_group_ids.#"},
			wantDropped: []string{"babelfish_config.#", "parameters.#"},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			params := tc.params()
			state := importedState(t, r.TerraformResource, params, tc.id)
			if tc.config != nil {
				params = tc.config()
			}
			cfg := terraform.NewResourceConfigRaw(params)
			diff, err := schema.InternalMap(r.TerraformResource.Schema).Diff(context.Background(), state, cfg, nil, nil, false)
			if err != nil {
				t.Fatalf("cannot compute the diff: %v", err)
			}
			for _, k := range append(append([]string{}, tc.wantKept...), tc.wantDropped...) {
				if diff == nil || diff.Attributes[k] == nil {
					t.Fatalf("precondition: the plugin SDK no longer reports %q; this test no longer reproduces the bug", k)
				}
			}

			diff, err = r.TerraformCustomDiff(diff, state, cfg)
			if err != nil {
				t.Fatalf("TerraformCustomDiff: %v", err)
			}
			for _, k := range tc.wantDropped {
				if d := diff.Attributes[k]; d != nil {
					t.Errorf("%q: want dropped, got %#v", k, d)
				}
			}
			for _, k := range tc.wantKept {
				if diff.Attributes[k] == nil {
					t.Errorf("%q: want kept, got dropped", k)
				}
			}
		})
	}
}

// importedState rebuilds a prior state the way upjet's
// TerraformPluginSDKConnector.Connect does when status.atProvider is empty:
// spec.forProvider plus the id, shimmed through cty.
func importedState(t *testing.T, res *schema.Resource, params map[string]any, id string) *terraform.InstanceState {
	t.Helper()
	m := make(map[string]any, len(params)+1)
	for k, v := range params {
		m[k] = v
	}
	m["id"] = id
	v, err := schema.JSONMapToStateValue(m, res.CoreConfigSchema())
	if err != nil {
		t.Fatalf("cannot convert the state map to cty: %v", err)
	}
	s, err := res.ShimInstanceStateFromValue(v)
	if err != nil {
		t.Fatalf("cannot shim the instance state: %v", err)
	}
	s.ID = id
	return s
}
