// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	v1beta1 "github.com/crossplane-contrib/provider-alibabacloud/apis/namespaced/v1beta1"
)

func TestBuildProviderConfiguration(t *testing.T) {
	t.Setenv("ALIBABA_CLOUD_ENDPOINT_SLS", "")
	t.Setenv("SLS_ENDPOINT", "")
	cfg := buildProviderConfiguration("cn-hangzhou", map[string]any{
		"access_key":     "ak",
		"secret_key":     "sk",
		"security_token": "token",
		"ignored":        "value",
	}, nil)

	want := map[string]any{
		"region":         "cn-hangzhou",
		"access_key":     "ak",
		"secret_key":     "sk",
		"security_token": "token",
		"endpoints": []any{map[string]any{
			"location": "location-readonly.aliyuncs.com",
			"log":      "cn-hangzhou.log.aliyuncs.com",
		}},
	}
	if !reflect.DeepEqual(map[string]any(cfg), want) {
		t.Fatalf("unexpected provider configuration:\n got: %#v\nwant: %#v", cfg, want)
	}
}

func TestEndpointSeeds(t *testing.T) {
	t.Setenv("ALIBABA_CLOUD_ENDPOINT_SLS", "")
	t.Setenv("SLS_ENDPOINT", "")
	want := map[string]any{"location": "location-readonly.aliyuncs.com", "log": "us-west-1.log.aliyuncs.com"}
	if got := endpointSeeds("us-west-1"); !reflect.DeepEqual(got, want) {
		t.Errorf("us-west-1: want %#v, got %#v", want, got)
	}
	if got := endpointSeeds(""); !reflect.DeepEqual(got, map[string]any{"location": "location-readonly.aliyuncs.com"}) {
		t.Errorf("no region: want only the location seed, got %#v", got)
	}
	for _, env := range []string{"ALIBABA_CLOUD_ENDPOINT_SLS", "SLS_ENDPOINT"} {
		t.Run(env, func(t *testing.T) {
			t.Setenv(env, "custom.example.com")
			if got := endpointSeeds("us-west-1"); got["log"] != nil {
				t.Errorf("with %s set: want no log seed so the override wins, got %#v", env, got)
			}
		})
	}
}

// TestEndpointSeedsFitProviderSchema checks the seeded endpoints block survives
// the real upstream provider schema with its keys intact: configure stores
// "log" in the client's endpoint cache, where loadApiEndpoint looks SLS up, and
// "location" in Config.LocationEndpoint, which the Location lookup asks.
func TestEndpointSeedsFitProviderSchema(t *testing.T) {
	t.Setenv("ALIBABA_CLOUD_ENDPOINT_SLS", "")
	t.Setenv("SLS_ENDPOINT", "")
	cfg := buildProviderConfiguration("us-west-1", map[string]any{"access_key": "ak", "secret_key": "sk"}, nil)
	p := alicloud.Provider()
	d := schema.TestResourceDataRaw(t, p.Schema, cfg)
	set, ok := d.Get("endpoints").(*schema.Set)
	if !ok || set.Len() != 1 {
		t.Fatalf("endpoints block did not survive the provider schema: %#v", d.Get("endpoints"))
	}
	got := set.List()[0].(map[string]any)
	if got["log"] != "us-west-1.log.aliyuncs.com" {
		t.Errorf("endpoints.log: want us-west-1.log.aliyuncs.com, got %#v", got["log"])
	}
	if got["location"] != "location-readonly.aliyuncs.com" {
		t.Errorf("endpoints.location: want location-readonly.aliyuncs.com, got %#v", got["location"])
	}
}

func TestBuildProviderConfigurationUsesTypedAssumeRole(t *testing.T) {
	roleARN := "acs:ram::1234567890123456:role/crossplane"
	sessionName := "crossplane"
	policy := "{}"
	sessionExpiration := 3600
	externalID := "external-id"

	cfg := buildProviderConfiguration("cn-hangzhou", map[string]any{
		"assume_role": map[string]any{
			"role_arn": "acs:ram::1234567890123456:role/from-secret",
		},
	}, &v1beta1.ProviderConfigSpec{
		AssumeRole: &v1beta1.AssumeRoleOptions{
			RoleARN:           roleARN,
			SessionName:       &sessionName,
			Policy:            &policy,
			SessionExpiration: &sessionExpiration,
			ExternalID:        &externalID,
		},
	})

	want := []any{
		map[string]any{
			"role_arn":           roleARN,
			"session_name":       sessionName,
			"policy":             policy,
			"session_expiration": sessionExpiration,
			"external_id":        externalID,
		},
	}
	if !reflect.DeepEqual(cfg["assume_role"], want) {
		t.Fatalf("unexpected assume_role value:\n got: %#v\nwant: %#v", cfg["assume_role"], want)
	}
}

func TestBuildProviderConfigurationUsesAssumeRoleWithOIDC(t *testing.T) {
	roleARN := "acs:ram::1234567890123456:role/crossplane-oidc"
	oidcProviderARN := "acs:ram::1234567890123456:oidc-provider/ack"
	oidcTokenFile := "/var/run/secrets/ack.alibabacloud.com/rrsa-tokens/token"
	roleSessionName := "crossplane"
	sessionExpiration := 3600

	cfg := buildProviderConfiguration("cn-hangzhou", map[string]any{}, &v1beta1.ProviderConfigSpec{
		AssumeRoleWithOIDC: &v1beta1.AssumeRoleWithOIDCOptions{
			RoleARN:           roleARN,
			OIDCProviderARN:   oidcProviderARN,
			OIDCTokenFile:     oidcTokenFile,
			RoleSessionName:   &roleSessionName,
			SessionExpiration: &sessionExpiration,
		},
	})

	wantOIDC := []any{
		map[string]any{
			"role_arn":           roleARN,
			"oidc_provider_arn":  oidcProviderARN,
			"oidc_token_file":    oidcTokenFile,
			"role_session_name":  roleSessionName,
			"session_expiration": sessionExpiration,
		},
	}
	if !reflect.DeepEqual(cfg["assume_role_with_oidc"], wantOIDC) {
		t.Fatalf("unexpected assume_role_with_oidc value:\n got: %#v\nwant: %#v", cfg["assume_role_with_oidc"], wantOIDC)
	}
}

func TestValidateProviderConfigRejectsDynamicAuthInSecret(t *testing.T) {
	err := validateProviderConfig(&v1beta1.ProviderConfigSpec{}, map[string]any{
		"assume_role": map[string]any{
			"role_arn": "acs:ram::1234567890123456:role/from-secret",
		},
	})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestValidateProviderConfigRequiresOIDCTokenFile(t *testing.T) {
	roleARN := "acs:ram::1234567890123456:role/crossplane-oidc"
	oidcProviderARN := "acs:ram::1234567890123456:oidc-provider/ack"

	err := validateProviderConfig(&v1beta1.ProviderConfigSpec{
		AssumeRoleWithOIDC: &v1beta1.AssumeRoleWithOIDCOptions{
			RoleARN:         roleARN,
			OIDCProviderARN: oidcProviderARN,
		},
	}, nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestValidateProviderConfigUsesDifferentSessionExpirationMaximums(t *testing.T) {
	tooLongForAssumeRole := 43200
	err := validateProviderConfig(&v1beta1.ProviderConfigSpec{
		AssumeRole: &v1beta1.AssumeRoleOptions{
			RoleARN:           "acs:ram::1234567890123456:role/crossplane",
			SessionExpiration: &tooLongForAssumeRole,
		},
	}, nil)
	if err == nil {
		t.Fatal("expected assumeRole sessionExpiration error")
	}

	oidcMax := 43200
	err = validateProviderConfig(&v1beta1.ProviderConfigSpec{
		AssumeRoleWithOIDC: &v1beta1.AssumeRoleWithOIDCOptions{
			RoleARN:           "acs:ram::1234567890123456:role/crossplane-oidc",
			OIDCProviderARN:   "acs:ram::1234567890123456:oidc-provider/ack",
			OIDCTokenFile:     "/var/run/secrets/ack.alibabacloud.com/rrsa-tokens/token",
			SessionExpiration: &oidcMax,
		},
	}, nil)
	if err != nil {
		t.Fatalf("unexpected assumeRoleWithOIDC sessionExpiration error: %v", err)
	}
}

// configureNoForkAlibabaCloudClient relies on p.Configure assigning p.meta
// from whatever p.ConfigureContextFunc returns, so a fake provider stands in
// for the real one: it declares one Required field and fails in the same way
// the real provider does when a credential is missing or of the wrong type
// (e.g. a nested assume_role block), without needing real credentials or a
// live Terraform provider.
func fakeCredentialedProvider() schema.Provider {
	return schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {Type: schema.TypeString, Required: true},
		},
		ConfigureContextFunc: func(_ context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
			if d.Get("access_key").(string) == "" {
				return nil, diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "missing access_key",
					Detail:   "access_key must be set",
				}}
			}
			return "configured", nil
		},
	}
}

func TestConfigureNoForkAlibabaCloudClientMissingCredentials(t *testing.T) {
	ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{}}

	err := configureNoForkAlibabaCloudClient(context.Background(), ps, fakeCredentialedProvider())
	if err == nil {
		t.Fatal("expected an error configuring the provider without access_key")
	}
	if !strings.Contains(err.Error(), "missing access_key") || !strings.Contains(err.Error(), "access_key must be set") {
		t.Fatalf("expected the error to surface both the diagnostic's Summary and Detail, got: %v", err)
	}
}

func TestConfigureNoForkAlibabaCloudClientSucceeds(t *testing.T) {
	ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{
		"access_key": "ak",
	}}

	if err := configureNoForkAlibabaCloudClient(context.Background(), ps, fakeCredentialedProvider()); err != nil {
		t.Fatalf("unexpected error configuring the provider: %v", err)
	}
	if ps.Meta != "configured" {
		t.Fatalf("expected ps.Meta to be set from the provider's ConfigureContextFunc, got: %#v", ps.Meta)
	}
}
