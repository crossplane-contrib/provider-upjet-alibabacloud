// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/json"
)

// A resource that provider 1.x created must decode, or the provider cannot
// list its kind at startup and stops.
func TestDecodeLegacyShape(t *testing.T) {
	t.Run("ManagedKubernetesFrom1.5", func(t *testing.T) {
		var mk ManagedKubernetesList
		err := json.Unmarshal([]byte(`{"items": [{"metadata": {"name": "mk"},
			"spec": {"forProvider": {"name": "mk"}},
			"status": {"atProvider": {"id": "c1",
				"certificateAuthority": {"clientCert": "cert"},
				"connections": {"apiServerInternet": "https://203.0.113.1:6443"}}}}]}`), &mk)
		if err != nil {
			t.Fatalf("cannot decode a 1.5 ManagedKubernetes: %v", err)
		}
		o := mk.Items[0].Status.AtProvider
		if o.CertificateAuthority != nil || o.Connections != nil {
			t.Errorf("1.x object-form status fields were kept: %+v %+v", o.CertificateAuthority, o.Connections)
		}
		if o.ID == nil || *o.ID != "c1" || *mk.Items[0].Spec.ForProvider.Name != "mk" {
			t.Error("the rest of the resource did not decode")
		}
	})

	t.Run("KubernetesFrom1.4", func(t *testing.T) {
		var k Kubernetes
		err := json.Unmarshal([]byte(`{"spec": {"forProvider": {"runtime": {"name": "containerd", "version": "1.6.28"}}},
			"status": {"atProvider": {"runtime": {"name": "containerd"}, "connections": {"api_server_internet": "x"}}}}`), &k)
		if err != nil {
			t.Fatalf("cannot decode a 1.4 Kubernetes: %v", err)
		}
		if k.Spec.ForProvider.Runtime != nil || k.Status.AtProvider.Runtime != nil {
			t.Error("1.x object-form runtime was kept")
		}
	})

	t.Run("EdgeKubernetesIn2.0Form", func(t *testing.T) {
		var ek EdgeKubernetes
		err := json.Unmarshal([]byte(`{"spec": {"forProvider": {"runtime": [{"name": "containerd", "version": "1.6.28"}]}}}`), &ek)
		if err != nil {
			t.Fatalf("cannot decode a 2.0 EdgeKubernetes: %v", err)
		}
		if rt := ek.Spec.ForProvider.Runtime; len(rt) != 1 || *rt[0].Name != "containerd" {
			t.Errorf("the 2.0 list was not kept: %+v", rt)
		}
	})
}
