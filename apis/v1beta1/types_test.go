/*
Copyright 2022 Upbound Inc.
*/

package v1beta1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

const credentialsSourceWebIdentity xpv1.CredentialsSource = "WebIdentity"

func TestProviderConfig(t *testing.T) {
	tests := map[string]struct {
		reason string
		pc     ProviderConfig
		want   ProviderConfig
	}{
		"NewProviderConfig": {
			reason: "A new ProviderConfig should have the correct structure",
			pc: ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config",
				},
				Spec: ProviderConfigSpec{
					Credentials: ProviderCredentials{
						Source: xpv1.CredentialsSourceSecret,
						CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
							SecretRef: &xpv1.SecretKeySelector{
								SecretReference: xpv1.SecretReference{
									Name:      "test-secret",
									Namespace: "test-namespace",
								},
								Key: "credentials",
							},
						},
					},
				},
			},
			want: ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config",
				},
				Spec: ProviderConfigSpec{
					Credentials: ProviderCredentials{
						Source: xpv1.CredentialsSourceSecret,
						CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
							SecretRef: &xpv1.SecretKeySelector{
								SecretReference: xpv1.SecretReference{
									Name:      "test-secret",
									Namespace: "test-namespace",
								},
								Key: "credentials",
							},
						},
					},
				},
			},
		},
		"ProviderConfigWithOIDC": {
			reason: "A ProviderConfig with OIDC options should have the correct nested structure",
			pc: ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config-oidc",
				},
				Spec: ProviderConfigSpec{
					Credentials: ProviderCredentials{
						Source: credentialsSourceWebIdentity,
						OIDC: &OIDCOptions{
							ProviderARN: "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
							RoleARN:     "acs:ram::1234567890123456:role/test-role",
							Region:      "cn-hangzhou",
						},
					},
				},
			},
			want: ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config-oidc",
				},
				Spec: ProviderConfigSpec{
					Credentials: ProviderCredentials{
						Source: credentialsSourceWebIdentity,
						OIDC: &OIDCOptions{
							ProviderARN: "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
							RoleARN:     "acs:ram::1234567890123456:role/test-role",
							Region:      "cn-hangzhou",
						},
					},
				},
			},
		},
		"ProviderConfigWithOIDCDefaultRegion": {
			reason: "A ProviderConfig with OIDC options and no explicit region should be valid",
			pc: ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config-oidc-default-region",
				},
				Spec: ProviderConfigSpec{
					Credentials: ProviderCredentials{
						Source: credentialsSourceWebIdentity,
						OIDC: &OIDCOptions{
							ProviderARN: "acs:ram::1234567890123456:oidc-provider/example-oidc-provider",
							RoleARN:     "acs:ram::1234567890123456:role/example-oidc-role",
						},
					},
				},
			},
			want: ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config-oidc-default-region",
				},
				Spec: ProviderConfigSpec{
					Credentials: ProviderCredentials{
						Source: credentialsSourceWebIdentity,
						OIDC: &OIDCOptions{
							ProviderARN: "acs:ram::1234567890123456:oidc-provider/example-oidc-provider",
							RoleARN:     "acs:ram::1234567890123456:role/example-oidc-role",
						},
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(tc.want, tc.pc); diff != "" {
				t.Errorf("\n%s\nProviderConfig: -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestProviderCredentials(t *testing.T) {
	tests := map[string]struct {
		reason string
		pc     ProviderCredentials
		want   ProviderCredentials
	}{
		"StandardCredentials": {
			reason: "Standard credentials should be correctly structured",
			pc: ProviderCredentials{
				Source: xpv1.CredentialsSourceSecret,
				CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
					SecretRef: &xpv1.SecretKeySelector{
						SecretReference: xpv1.SecretReference{
							Name:      "test-secret",
							Namespace: "test-namespace",
						},
						Key: "credentials",
					},
				},
			},
			want: ProviderCredentials{
				Source: xpv1.CredentialsSourceSecret,
				CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
					SecretRef: &xpv1.SecretKeySelector{
						SecretReference: xpv1.SecretReference{
							Name:      "test-secret",
							Namespace: "test-namespace",
						},
						Key: "credentials",
					},
				},
			},
		},
		"OIDCCredentials": {
			reason: "OIDC credentials should use the nested OIDCOptions struct",
			pc: ProviderCredentials{
				Source: credentialsSourceWebIdentity,
				OIDC: &OIDCOptions{
					ProviderARN: "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
					RoleARN:     "acs:ram::1234567890123456:role/test-role",
					Region:      "cn-hangzhou",
				},
			},
			want: ProviderCredentials{
				Source: credentialsSourceWebIdentity,
				OIDC: &OIDCOptions{
					ProviderARN: "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
					RoleARN:     "acs:ram::1234567890123456:role/test-role",
					Region:      "cn-hangzhou",
				},
			},
		},
		"OIDCCredentialsWithSessionName": {
			reason: "OIDC credentials with a custom session name should be correctly structured",
			pc: ProviderCredentials{
				Source: credentialsSourceWebIdentity,
				OIDC: &OIDCOptions{
					ProviderARN:     "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
					RoleARN:         "acs:ram::1234567890123456:role/test-role",
					Region:          "cn-hangzhou",
					RoleSessionName: "my-custom-session",
				},
			},
			want: ProviderCredentials{
				Source: credentialsSourceWebIdentity,
				OIDC: &OIDCOptions{
					ProviderARN:     "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
					RoleARN:         "acs:ram::1234567890123456:role/test-role",
					Region:          "cn-hangzhou",
					RoleSessionName: "my-custom-session",
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(tc.want, tc.pc); diff != "" {
				t.Errorf("\n%s\nProviderCredentials: -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestOIDCOptions(t *testing.T) {
	tests := map[string]struct {
		reason string
		opts   OIDCOptions
		want   OIDCOptions
	}{
		"RequiredFieldsOnly": {
			reason: "OIDCOptions with only required fields should be valid",
			opts: OIDCOptions{
				RoleARN:     "acs:ram::1234567890123456:role/test-role",
				ProviderARN: "acs:ram::1234567890123456:oidc-provider/test-provider",
			},
			want: OIDCOptions{
				RoleARN:     "acs:ram::1234567890123456:role/test-role",
				ProviderARN: "acs:ram::1234567890123456:oidc-provider/test-provider",
			},
		},
		"AllFields": {
			reason: "OIDCOptions with all fields set should be correctly structured",
			opts: OIDCOptions{
				RoleARN:         "acs:ram::1234567890123456:role/test-role",
				ProviderARN:     "acs:ram::1234567890123456:oidc-provider/test-provider",
				Region:          "cn-beijing",
				RoleSessionName: "custom-session",
			},
			want: OIDCOptions{
				RoleARN:         "acs:ram::1234567890123456:role/test-role",
				ProviderARN:     "acs:ram::1234567890123456:oidc-provider/test-provider",
				Region:          "cn-beijing",
				RoleSessionName: "custom-session",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(tc.want, tc.opts); diff != "" {
				t.Errorf("\n%s\nOIDCOptions: -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
