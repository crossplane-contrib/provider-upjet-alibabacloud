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

// TestOOSServiceSettingDeliveryOSSFields guards against a regression of
// https://github.com/aliyun/terraform-provider-alicloud/issues/11154: upstream
// used to suppress the diff on delivery_oss_bucket_name and
// delivery_oss_key_prefix unless delivery_sls_enabled was true, so Create
// never sent them. Terraform provider 2.2.0 fixed it, and this provider
// dropped its own workaround.
func TestOOSServiceSettingDeliveryOSSFields(t *testing.T) {
	p, err := providerconfig.GetProvider(context.Background(), false)
	if err != nil {
		t.Fatalf("cannot build the runtime provider configuration: %v", err)
	}
	res := p.Resources["alicloud_oos_service_setting"].TerraformResource

	params := map[string]any{
		"region":                   "us-west-1",
		"delivery_oss_enabled":     true,
		"delivery_oss_bucket_name": "example-bucket",
		"delivery_oss_key_prefix":  "path1/",
	}
	diff := func(state *terraform.InstanceState) *terraform.InstanceDiff {
		t.Helper()
		d, err := schema.InternalMap(res.Schema).Diff(context.Background(), state, terraform.NewResourceConfigRaw(params), nil, nil, false)
		if err != nil {
			t.Fatalf("cannot compute the diff: %v", err)
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
		d := diff(nil)
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
		if d := diff(existing("example-bucket")); has(d, "delivery_oss_bucket_name") || has(d, "delivery_oss_key_prefix") {
			t.Errorf("unchanged fields must not diff, got %#v", d.Attributes)
		}
	})
	t.Run("ChangedBucketIsSent", func(t *testing.T) {
		d := diff(existing("old-bucket"))
		if a := d.Attributes["delivery_oss_bucket_name"]; a == nil || a.Old != "old-bucket" || a.New != "example-bucket" {
			t.Errorf("want old-bucket -> example-bucket, got %#v", a)
		}
	})
}
