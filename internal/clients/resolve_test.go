// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"testing"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	v1beta1 "github.com/crossplane-contrib/provider-alibabacloud/apis/namespaced/v1beta1"
	vpcv1alpha1 "github.com/crossplane-contrib/provider-alibabacloud/apis/namespaced/vpc/v1alpha1"
)

const (
	mrNamespace     = "team-a"
	sharedNamespace = "crossplane-system"
)

func credentialsSecret(namespace, accessKey string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: "creds", Namespace: namespace},
		Data: map[string][]byte{
			"credentials": []byte(`{"access_key":"` + accessKey + `","secret_key":"sk"}`),
		},
	}
}

func secretCredentials(namespace string) v1beta1.ProviderCredentials {
	return v1beta1.ProviderCredentials{
		Source: xpv1.CredentialsSourceSecret,
		CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
			SecretRef: &xpv1.SecretKeySelector{
				SecretReference: xpv1.SecretReference{Name: "creds", Namespace: namespace},
				Key:             "credentials",
			},
		},
	}
}

// TestResolveModernSecretNamespace checks where each modern config kind reads
// its credentials secret from. Both namespaces hold a secret with the same
// name and different contents, so the access key shows which one was read.
func TestResolveModernSecretNamespace(t *testing.T) {
	cases := map[string]struct {
		reason    string
		ref       *xpv1.ProviderConfigReference
		config    runtime.Object
		wantKeyID string
	}{
		"ClusterProviderConfig": {
			reason: "A ClusterProviderConfig is shared across namespaces, so the secret should be read from the namespace its reference names.",
			ref:    &xpv1.ProviderConfigReference{Kind: v1beta1.ClusterProviderConfigKind, Name: "default"},
			config: &v1beta1.ClusterProviderConfig{
				ObjectMeta: metav1.ObjectMeta{Name: "default"},
				Spec:       v1beta1.ProviderConfigSpec{Credentials: secretCredentials(sharedNamespace)},
			},
			wantKeyID: "from-" + sharedNamespace,
		},
		"NamespacedProviderConfig": {
			reason: "A namespaced ProviderConfig's secret reference is local, so the secret should be read from the MR's namespace whatever the reference names.",
			ref:    &xpv1.ProviderConfigReference{Kind: v1beta1.ProviderConfigKind, Name: "default"},
			config: &v1beta1.ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{Name: "default", Namespace: mrNamespace},
				Spec:       v1beta1.ProviderConfigSpec{Credentials: secretCredentials(sharedNamespace)},
			},
			wantKeyID: "from-" + mrNamespace,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := runtime.NewScheme()
			for _, add := range []func(*runtime.Scheme) error{corev1.AddToScheme, v1beta1.SchemeBuilder.AddToScheme, vpcv1alpha1.SchemeBuilder.AddToScheme} {
				if err := add(s); err != nil {
					t.Fatal(err)
				}
			}
			mg := &vpcv1alpha1.VPC{ObjectMeta: metav1.ObjectMeta{Name: "vpc", Namespace: mrNamespace, UID: "uid"}}
			mg.SetGroupVersionKind(vpcv1alpha1.VPC_GroupVersionKind)
			mg.Spec.ProviderConfigReference = tc.ref
			c := fake.NewClientBuilder().WithScheme(s).WithRuntimeObjects(
				tc.config,
				credentialsSecret(sharedNamespace, "from-"+sharedNamespace),
				credentialsSecret(mrNamespace, "from-"+mrNamespace),
			).Build()

			_, creds, err := resolveProviderConfig(context.Background(), c, mg)
			if err != nil {
				t.Fatalf("\n%s\nresolveProviderConfig(...): unexpected error: %v", tc.reason, err)
			}
			if got := creds["access_key"]; got != tc.wantKeyID {
				t.Errorf("\n%s\nresolveProviderConfig(...): access_key = %v, want %q", tc.reason, got, tc.wantKeyID)
			}
		})
	}
}
