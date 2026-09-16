// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"reflect"
	"testing"

	v1beta1 "github.com/crossplane-contrib/provider-alibabacloud/apis/namespaced/v1beta1"
)

func TestBuildProviderConfiguration(t *testing.T) {
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
	}
	if !reflect.DeepEqual(map[string]any(cfg), want) {
		t.Fatalf("unexpected provider configuration:\n got: %#v\nwant: %#v", cfg, want)
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
