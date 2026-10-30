// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud"

	providerconfig "github.com/crossplane-contrib/provider-alibabacloud/config"
)

// TestAliKafkaChildDeleteWrapped checks that every resource living inside an
// alikafka instance runs through the provider's delete wrapper at runtime, so
// deleting it after its instance was released does not leave it stuck.
func TestAliKafkaChildDeleteWrapped(t *testing.T) {
	children := []string{
		"alicloud_alikafka_topic",
		"alicloud_alikafka_consumer_group",
		"alicloud_alikafka_scheduled_scaling_rule",
		"alicloud_alikafka_sasl_user",
		"alicloud_alikafka_sasl_acl",
		"alicloud_alikafka_instance_allowed_ip_attachment",
	}
	p, err := providerconfig.GetProvider(context.Background(), false)
	if err != nil {
		t.Fatalf("cannot build the runtime provider configuration: %v", err)
	}
	upstream := alicloud.Provider().ResourcesMap
	for _, name := range children {
		r, ok := p.Resources[name]
		if !ok {
			t.Errorf("%s is not configured", name)
			continue
		}
		got := r.TerraformResource.Delete //nolint:staticcheck // upstream implements these resources with the legacy Delete field.
		want := upstream[name].Delete     //nolint:staticcheck // as above.
		if got == nil || want == nil {
			t.Errorf("%s: upstream no longer sets Delete; the wrapper must follow it to the field it uses", name)
			continue
		}
		if reflect.ValueOf(got).Pointer() == reflect.ValueOf(want).Pointer() {
			t.Errorf("%s: Delete is upstream's own function, not the provider's wrapper", name)
		}
	}
}

// TestPolarDBChildReadWrapped checks that the PolarDB resources whose upstream
// Read retries for minutes while their cluster is missing run through the
// provider's Read wrapper at runtime.
func TestPolarDBChildReadWrapped(t *testing.T) {
	children := []string{
		"alicloud_polardb_endpoint",
		"alicloud_polardb_cluster_endpoint",
		"alicloud_polardb_primary_endpoint",
		"alicloud_polardb_database",
	}
	p, err := providerconfig.GetProvider(context.Background(), false)
	if err != nil {
		t.Fatalf("cannot build the runtime provider configuration: %v", err)
	}
	upstream := alicloud.Provider().ResourcesMap
	for _, name := range children {
		r, ok := p.Resources[name]
		if !ok {
			t.Errorf("%s is not configured", name)
			continue
		}
		got := r.TerraformResource.Read //nolint:staticcheck // upstream implements these resources with the legacy Read field.
		want := upstream[name].Read     //nolint:staticcheck // as above.
		if got == nil || want == nil {
			t.Errorf("%s: upstream no longer sets Read; the wrapper must follow it to the field it uses", name)
			continue
		}
		if reflect.ValueOf(got).Pointer() == reflect.ValueOf(want).Pointer() {
			t.Errorf("%s: Read is upstream's own function, not the provider's wrapper", name)
		}
	}
}
