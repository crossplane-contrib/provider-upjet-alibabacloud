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

// TestOOSServiceSettingDeliveryOSSFields covers the workaround for
// https://github.com/aliyun/terraform-provider-alicloud/issues/11154: upstream
// suppresses the diff on delivery_oss_bucket_name and delivery_oss_key_prefix
// unless delivery_sls_enabled is true, so Create never sent them.
func TestOOSServiceSettingDeliveryOSSFields(t *testing.T) {
	p, err := providerconfig.GetProvider(context.Background(), false)
	if err != nil {
		t.Fatalf("cannot build the runtime provider configuration: %v", err)
	}
	r := p.Resources["alicloud_oos_service_setting"]
	res := r.TerraformResource

	params := func(extra map[string]any) map[string]any {
		m := map[string]any{
			"region":                   "us-west-1",
			"delivery_oss_enabled":     true,
			"delivery_oss_bucket_name": "example-bucket",
			"delivery_oss_key_prefix":  "path1/",
		}
		for k, v := range extra {
			m[k] = v
		}
		return m
	}
	diff := func(state *terraform.InstanceState, params map[string]any, custom bool) *terraform.InstanceDiff {
		t.Helper()
		cfg := terraform.NewResourceConfigRaw(params)
		d, err := schema.InternalMap(res.Schema).Diff(context.Background(), state, cfg, nil, nil, false)
		if err != nil {
			t.Fatalf("cannot compute the diff: %v", err)
		}
		if custom {
			if d, err = r.TerraformCustomDiff(d, state, cfg); err != nil {
				t.Fatalf("TerraformCustomDiff: %v", err)
			}
		}
		return d
	}
	has := func(d *terraform.InstanceDiff, k string) bool { return d != nil && d.Attributes[k] != nil }
	existing := func(bucket string) *terraform.InstanceState {
		return &terraform.InstanceState{ID: "us-west-1", Attributes: map[string]string{
			"id": "us-west-1", "delivery_oss_enabled": "true",
			"delivery_oss_bucket_name": bucket, "delivery_oss_key_prefix": "path1/",
		}}
	}

	t.Run("CreateSendsBucket", func(t *testing.T) {
		if has(diff(nil, params(nil), false), "delivery_oss_bucket_name") {
			t.Fatal("precondition: upstream no longer suppresses delivery_oss_bucket_name; the workaround may be obsolete")
		}
		d := diff(nil, params(nil), true)
		for _, k := range []string{"delivery_oss_bucket_name", "delivery_oss_key_prefix"} {
			if !has(d, k) {
				t.Errorf("%s missing from the diff", k)
			}
		}
		data, err := schema.InternalMap(res.Schema).Data(nil, d)
		if err != nil {
			t.Fatalf("cannot build ResourceData: %v", err)
		}
		if v, ok := data.GetOk("delivery_oss_bucket_name"); !ok || v != "example-bucket" {
			t.Errorf("Create would read delivery_oss_bucket_name = %v, %v; want example-bucket", v, ok)
		}
	})
	t.Run("UnchangedIsNotAChange", func(t *testing.T) {
		if d := diff(existing("example-bucket"), params(nil), true); has(d, "delivery_oss_bucket_name") || has(d, "delivery_oss_key_prefix") {
			t.Errorf("unchanged fields must not diff, got %#v", d.Attributes)
		}
	})
	t.Run("ChangedBucketIsSent", func(t *testing.T) {
		d := diff(existing("old-bucket"), params(nil), true)
		if a := d.Attributes["delivery_oss_bucket_name"]; a == nil || a.Old != "old-bucket" || a.New != "example-bucket" {
			t.Errorf("want old-bucket -> example-bucket, got %#v", a)
		}
	})
	t.Run("SLSEnabledLeavesUpstreamDiff", func(t *testing.T) {
		extra := map[string]any{"delivery_sls_enabled": true, "delivery_sls_project_name": "p"}
		a := diff(nil, params(extra), false).Attributes["delivery_oss_bucket_name"]
		b := diff(nil, params(extra), true).Attributes["delivery_oss_bucket_name"]
		if a == nil || b == nil || a.Old != b.Old || a.New != b.New {
			t.Errorf("with SLS delivery on, the SDK's own diff should be kept as is: got %#v, want %#v", b, a)
		}
	})
}
