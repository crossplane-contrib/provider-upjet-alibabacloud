// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config_test

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	providerconfig "github.com/crossplane-contrib/provider-alibabacloud/config"
)

// TestALBZoneMappingsChurn reproduces the ALB LoadBalancer that never synced:
// once the cloud fills a zone mapping's Optional+Computed fields
// (intranet_address, eip_type), the element in state hashes differently from
// the same element in spec.forProvider, and the plugin SDK reports every zone
// mapping as removed and re-added on every reconcile.
func TestALBZoneMappingsChurn(t *testing.T) {
	p, err := providerconfig.GetProvider(context.Background(), false)
	if err != nil {
		t.Fatalf("cannot build the runtime provider configuration: %v", err)
	}
	r := p.Resources["alicloud_alb_load_balancer"]
	res := r.TerraformResource

	zones := func(vsw1 string) []any {
		return []any{
			map[string]any{"zone_id": "us-west-1a", "vswitch_id": vsw1},
			map[string]any{"zone_id": "us-west-1b", "vswitch_id": "vsw-b"},
		}
	}
	base := map[string]any{
		"region":                 "us-west-1",
		"vpc_id":                 "vpc-x",
		"address_type":           "Intranet",
		"address_allocated_mode": "Fixed",
		"load_balancer_edition":  "Standard",
	}
	with := func(z []any) map[string]any {
		m := map[string]any{}
		for k, v := range base {
			m[k] = v
		}
		m["zone_mappings"] = z
		return m
	}

	// The observed state: the same zone mappings, with the computed fields
	// the cloud filled in.
	observed := with([]any{
		map[string]any{"zone_id": "us-west-1a", "vswitch_id": "vsw-a", "intranet_address": "172.16.0.5", "eip_type": "Common"},
		map[string]any{"zone_id": "us-west-1b", "vswitch_id": "vsw-b", "intranet_address": "172.16.1.5", "eip_type": "Common"},
	})
	observed["id"] = "alb-x"
	v, err := schema.JSONMapToStateValue(observed, res.CoreConfigSchema())
	if err != nil {
		t.Fatalf("cannot convert the state: %v", err)
	}
	state, err := res.ShimInstanceStateFromValue(v)
	if err != nil {
		t.Fatalf("cannot shim the state: %v", err)
	}
	state.ID = "alb-x"

	churn := func(d *terraform.InstanceDiff) []string {
		var out []string
		if d == nil {
			return nil
		}
		for k := range d.Attributes {
			if strings.HasPrefix(k, "zone_mappings.") {
				out = append(out, k)
			}
		}
		return out
	}
	diff := func(params map[string]any) (*terraform.InstanceDiff, *terraform.ResourceConfig) {
		t.Helper()
		cfg := terraform.NewResourceConfigRaw(params)
		d, err := schema.InternalMap(res.Schema).Diff(context.Background(), state, cfg, nil, nil, false)
		if err != nil {
			t.Fatalf("cannot compute the diff: %v", err)
		}
		return d, cfg
	}

	t.Run("SameZonesAreNotAChange", func(t *testing.T) {
		d, cfg := diff(with(zones("vsw-a")))
		if len(churn(d)) == 0 {
			t.Fatal("precondition: the plugin SDK no longer reports zone_mappings churn; this test no longer reproduces the bug")
		}
		d, err := r.TerraformCustomDiff(d, state, cfg)
		if err != nil {
			t.Fatalf("TerraformCustomDiff: %v", err)
		}
		if c := churn(d); len(c) != 0 {
			t.Errorf("want the churn dropped, still have %v", c)
		}
	})
	t.Run("ChangedVSwitchIsKept", func(t *testing.T) {
		d, cfg := diff(with(zones("vsw-new")))
		d, err := r.TerraformCustomDiff(d, state, cfg)
		if err != nil {
			t.Fatalf("TerraformCustomDiff: %v", err)
		}
		if len(churn(d)) == 0 {
			t.Error("a changed vswitch_id must still produce a zone_mappings diff")
		}
	})
}

// TestALBServerGroupServersChurn checks the set-churn suppression compares only
// what the user configured: a weight the cloud filled in is not a change, but a
// user's own weight change, or an added server, still is. Keying the elements
// by server_id alone would hide the weight change.
func TestALBServerGroupServersChurn(t *testing.T) {
	p, err := providerconfig.GetProvider(context.Background(), false)
	if err != nil {
		t.Fatalf("cannot build the runtime provider configuration: %v", err)
	}
	r := p.Resources["alicloud_alb_server_group"]
	res := r.TerraformResource

	observed := map[string]any{
		"id":                "sgp-x",
		"region":            "us-west-1",
		"server_group_name": "example",
		"servers": []any{map[string]any{
			"server_type": "Ecs", "server_id": "i-1", "port": 80,
			"weight": 100, "server_ip": "10.0.0.1", "remote_ip_enabled": false,
		}},
	}
	v, err := schema.JSONMapToStateValue(observed, res.CoreConfigSchema())
	if err != nil {
		t.Fatalf("cannot convert the state: %v", err)
	}
	state, err := res.ShimInstanceStateFromValue(v)
	if err != nil {
		t.Fatalf("cannot shim the state: %v", err)
	}
	state.ID = "sgp-x"

	server := func(extra map[string]any) map[string]any {
		m := map[string]any{"server_type": "Ecs", "server_id": "i-1", "port": 80}
		for k, v := range extra {
			m[k] = v
		}
		return m
	}
	run := func(servers []any) []string {
		t.Helper()
		params := map[string]any{"region": "us-west-1", "server_group_name": "example", "servers": servers}
		cfg := terraform.NewResourceConfigRaw(params)
		d, err := schema.InternalMap(res.Schema).Diff(context.Background(), state, cfg, nil, nil, false)
		if err != nil {
			t.Fatalf("cannot compute the diff: %v", err)
		}
		if d == nil {
			return nil
		}
		d, err = r.TerraformCustomDiff(d, state, cfg)
		if err != nil {
			t.Fatalf("TerraformCustomDiff: %v", err)
		}
		var out []string
		for k := range d.Attributes {
			if strings.HasPrefix(k, "servers.") {
				out = append(out, k)
			}
		}
		return out
	}

	if c := run([]any{server(nil)}); len(c) != 0 {
		t.Errorf("cloud-filled weight/server_ip: want no servers diff, got %v", c)
	}
	if c := run([]any{server(map[string]any{"weight": 50})}); len(c) == 0 {
		t.Error("a user's weight change must still produce a servers diff")
	}
	if c := run([]any{server(nil), server(map[string]any{"server_id": "i-2"})}); len(c) == 0 {
		t.Error("an added server must still produce a servers diff")
	}
}
