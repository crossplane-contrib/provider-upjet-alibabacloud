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
			reason: "A ProviderConfig with OIDC fields should have the correct structure",
			pc: ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config-oidc",
				},
				Spec: ProviderConfigSpec{
					Credentials: ProviderCredentials{
						Source:      xpv1.CredentialsSourceNone,
						Region:      "cn-hangzhou",
						ProviderARN: "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
						RoleARN:     "acs:ram::1234567890123456:role/test-role",
					},
				},
			},
			want: ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config-oidc",
				},
				Spec: ProviderConfigSpec{
					Credentials: ProviderCredentials{
						Source:      xpv1.CredentialsSourceNone,
						Region:      "cn-hangzhou",
						ProviderARN: "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
						RoleARN:     "acs:ram::1234567890123456:role/test-role",
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
			reason: "OIDC credentials should be correctly structured",
			pc: ProviderCredentials{
				Source:      xpv1.CredentialsSourceNone,
				Region:      "cn-hangzhou",
				ProviderARN: "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
				RoleARN:     "acs:ram::1234567890123456:role/test-role",
			},
			want: ProviderCredentials{
				Source:      xpv1.CredentialsSourceNone,
				Region:      "cn-hangzhou",
				ProviderARN: "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
				RoleARN:     "acs:ram::1234567890123456:role/test-role",
			},
		},
		"MixedCredentials": {
			reason: "Credentials with both standard and OIDC fields should be correctly structured",
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
				Region:      "cn-hangzhou",
				ProviderARN: "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
				RoleARN:     "acs:ram::1234567890123456:role/test-role",
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
				Region:      "cn-hangzhou",
				ProviderARN: "acs:ram::1234567890123456:oidc-provider/ack-rrsa-abc123",
				RoleARN:     "acs:ram::1234567890123456:role/test-role",
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
