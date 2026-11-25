// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config_test

import (
	"context"
	"testing"

	providerconfig "github.com/crossplane-contrib/provider-alibabacloud/config"
)

// Only cluster-scoped ACK kinds can hold what provider 1.x stored, so only
// they get the legacy shape check, as their first initializer.
func TestLegacyShapeGuardScope(t *testing.T) {
	cluster, err := providerconfig.GetProvider(context.Background(), false)
	if err != nil {
		t.Fatal(err)
	}
	namespaced, err := providerconfig.GetNamespacedProvider(context.Background(), false)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"alicloud_cs_edge_kubernetes", "alicloud_cs_kubernetes", "alicloud_cs_managed_kubernetes"} {
		c, n := len(cluster.Resources[name].InitializerFns), len(namespaced.Resources[name].InitializerFns)
		if c != n+1 {
			t.Errorf("%s: %d initializers cluster-scoped, %d namespaced; want exactly one more, the legacy shape check", name, c, n)
		}
	}
	if c, n := len(cluster.Resources["alicloud_vpc"].InitializerFns), len(namespaced.Resources["alicloud_vpc"].InitializerFns); c != n {
		t.Errorf("alicloud_vpc: %d initializers cluster-scoped, %d namespaced; want the same", c, n)
	}
}
