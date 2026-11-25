// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	providerconfig "github.com/crossplane-contrib/provider-alibabacloud/config"
)

// The same backend attached to an endpoint service in two zones is two
// attachments, and both must reach CreateVpcEndpointService (#212).
func TestEndpointServiceResourceZones(t *testing.T) {
	p, err := providerconfig.GetProvider(context.Background(), false)
	if err != nil {
		t.Fatalf("cannot build the provider configuration: %v", err)
	}
	r := p.Resources["alicloud_privatelink_vpc_endpoint_service"].TerraformResource
	nlb := func(zone any) map[string]any {
		m := map[string]any{"resource_id": "nlb-gfvi1lgpi28b7v72qa", "resource_type": "nlb"}
		if zone != nil {
			m["zone_id"] = zone
		}
		return m
	}
	cases := map[string]struct {
		resources []any
		want      int
	}{
		"SameNLBInTwoZones":   {resources: []any{nlb("cn-heyuan-a"), nlb("cn-heyuan-b")}, want: 2},
		"SameNLBAndZoneTwice": {resources: []any{nlb("cn-heyuan-a"), nlb("cn-heyuan-a")}, want: 1},
		"SLBWithoutAZone":     {resources: []any{nlb(nil), nlb("")}, want: 1},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// Create reads the block exactly like this, through d.GetOk.
			d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{"resource": tc.resources})
			v, ok := d.GetOk("resource")
			if !ok {
				t.Fatal("resource is not set")
			}
			if got := v.(*schema.Set).Len(); got != tc.want {
				t.Errorf("resource holds %d attachments, want %d: %v", got, tc.want, v.(*schema.Set).List())
			}
		})
	}
}
