// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	v1beta1 "github.com/crossplane-contrib/provider-alibabacloud/apis/namespaced/v1beta1"
)

// countingProvider returns a provider whose meta records how many times it was
// configured, so a test can tell a cached client from a fresh one.
func countingProvider(calls *int) schema.Provider {
	return schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {Type: schema.TypeString, Required: true},
			"region":     {Type: schema.TypeString, Optional: true},
		},
		ConfigureContextFunc: func(_ context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
			if d.Get("access_key").(string) == "" {
				return nil, diag.Diagnostics{{Severity: diag.Error, Summary: "missing access_key"}}
			}
			*calls++
			return fmt.Sprintf("client-%d", *calls), nil
		},
	}
}

func setup(cfg terraform.ProviderConfiguration) *terraform.Setup {
	return &terraform.Setup{Configuration: cfg}
}

func TestProviderMetaCacheReusesClient(t *testing.T) {
	calls := 0
	now := time.Unix(0, 0)
	c := newProviderMetaCache()
	c.now = func() time.Time { return now }
	p := countingProvider(&calls)
	cfg := func(region, ak string) terraform.ProviderConfiguration {
		return terraform.ProviderConfiguration{"access_key": ak, "region": region}
	}

	configure := func(cfg terraform.ProviderConfiguration) any {
		t.Helper()
		ps := setup(cfg)
		if err := c.configure(context.Background(), ps, p); err != nil {
			t.Fatalf("configure: %v", err)
		}
		return ps.Meta
	}

	first := configure(cfg("us-west-1", "ak"))
	if got := configure(cfg("us-west-1", "ak")); got != first || calls != 1 {
		t.Fatalf("same configuration: want the cached %v and 1 Configure call, got %v and %d", first, got, calls)
	}
	if got := configure(cfg("cn-beijing", "ak")); got == first || calls != 2 {
		t.Fatalf("different region: want a new client, got %v after %d calls", got, calls)
	}
	if got := configure(cfg("us-west-1", "rotated")); got == first || calls != 3 {
		t.Fatalf("rotated credentials: want a new client, got %v after %d calls", got, calls)
	}

	now = now.Add(defaultMetaTTL)
	if got := configure(cfg("us-west-1", "ak")); got == first || calls != 4 {
		t.Fatalf("after the TTL: want a new client, got %v after %d calls", got, calls)
	}
	if n := len(c.entries); n != 1 {
		t.Fatalf("want expired entries swept, got %d entries", n)
	}
}

func TestProviderMetaCacheDoesNotCacheErrors(t *testing.T) {
	calls := 0
	c := newProviderMetaCache()
	p := countingProvider(&calls)

	if err := c.configure(context.Background(), setup(terraform.ProviderConfiguration{"access_key": ""}), p); err == nil {
		t.Fatal("want an error configuring without access_key")
	}
	if n := len(c.entries); n != 0 {
		t.Fatalf("want a failed configuration left uncached, got %d entries", n)
	}
}

func TestMetaTTL(t *testing.T) {
	cases := map[string]struct {
		cfg  terraform.ProviderConfiguration
		want time.Duration
	}{
		"StaticCredentials": {
			cfg:  terraform.ProviderConfiguration{"access_key": "ak"},
			want: defaultMetaTTL,
		},
		"LongAssumeRoleSession": {
			cfg:  terraform.ProviderConfiguration{"assume_role": []any{map[string]any{"session_expiration": 3600}}},
			want: defaultMetaTTL,
		},
		"ShortAssumeRoleSession": {
			cfg:  terraform.ProviderConfiguration{"assume_role": []any{map[string]any{"session_expiration": 900}}},
			want: 450 * time.Second,
		},
		"ShortOIDCSession": {
			cfg:  terraform.ProviderConfiguration{"assume_role_with_oidc": []any{map[string]any{"session_expiration": 1000}}},
			want: 500 * time.Second,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := metaTTL(tc.cfg); got != tc.want {
				t.Errorf("metaTTL: want %v, got %v", tc.want, got)
			}
		})
	}
}

// TestMetaTTLMatchesBuildProviderConfiguration guards the shape metaTTL reads:
// if buildProviderConfiguration stops emitting session_expiration as an int in
// a single-element block, the short-session TTL silently stops applying.
func TestMetaTTLMatchesBuildProviderConfiguration(t *testing.T) {
	cfg := buildProviderConfiguration("us-west-1", map[string]any{"access_key": "ak"}, &providerConfigWithAssumeRoleSession(900).Spec)
	if got := metaTTL(cfg); got != 450*time.Second {
		t.Fatalf("metaTTL of a built configuration with a 900s session: want 7m30s, got %v", got)
	}
}

func providerConfigWithAssumeRoleSession(seconds int) *v1beta1.ProviderConfig {
	return &v1beta1.ProviderConfig{Spec: v1beta1.ProviderConfigSpec{
		AssumeRole: &v1beta1.AssumeRoleOptions{
			RoleARN:           "acs:ram::1234567890123456:role/crossplane",
			SessionExpiration: &seconds,
		},
	}}
}
