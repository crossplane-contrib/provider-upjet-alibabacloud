// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/json"
)

// Objects that provider 1.4 stored hold these fields as maps, with
// Terraform's snake_case keys. They must still decode: the provider lists
// every object of a kind at startup, and one that fails to decode stops it.
func TestDecodeLegacyShape(t *testing.T) {
	t.Run("ManagedKubernetes status", func(t *testing.T) {
		var mk ManagedKubernetes
		err := json.Unmarshal([]byte(`{"status": {"atProvider": {
			"certificateAuthority": {"client_cert": "cert", "client_key": "key", "cluster_cert": "ca"},
			"connections": {"api_server_internet": "https://203.0.113.1:6443"}}}}`), &mk)
		if err != nil {
			t.Fatalf("cannot decode a 1.4 ManagedKubernetes: %v", err)
		}
		if mk.Status.AtProvider.CertificateAuthority == nil || mk.Status.AtProvider.Connections == nil {
			t.Error("the legacy status fields did not decode as objects")
		}
	})

	t.Run("Kubernetes spec runtime", func(t *testing.T) {
		var k Kubernetes
		err := json.Unmarshal([]byte(`{"spec": {
			"forProvider": {"runtime": {"name": "containerd", "version": "1.6.28"}},
			"initProvider": {"runtime": {"name": "containerd"}}}}`), &k)
		if err != nil {
			t.Fatalf("cannot decode a 1.4 Kubernetes: %v", err)
		}
		rt := k.Spec.ForProvider.Runtime
		if rt == nil || rt.Name == nil || *rt.Name != "containerd" || rt.Version == nil || *rt.Version != "1.6.28" {
			t.Errorf("spec.forProvider.runtime = %+v; want containerd 1.6.28, as 1.4 stored it", rt)
		}
		if k.Spec.InitProvider.Runtime == nil || *k.Spec.InitProvider.Runtime.Name != "containerd" {
			t.Errorf("spec.initProvider.runtime did not keep its name")
		}
	})

	t.Run("EdgeKubernetes status", func(t *testing.T) {
		var ek EdgeKubernetes
		err := json.Unmarshal([]byte(`{"status": {"atProvider": {
			"certificateAuthority": {"client_cert": "cert"},
			"connections": {"service_domain": "example.invalid"},
			"runtime": {"name": "containerd"}}}}`), &ek)
		if err != nil {
			t.Fatalf("cannot decode a 1.4 EdgeKubernetes: %v", err)
		}
	})
}
