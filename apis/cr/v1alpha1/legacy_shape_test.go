// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/json"
)

// Provider 1.4 stored domainList as a map. It must still decode, or the
// provider cannot list Repos at startup and stops.
func TestDecodeLegacyDomainList(t *testing.T) {
	var r Repo
	if err := json.Unmarshal([]byte(`{"status": {"atProvider": {"domainList": {"public": "registry.example.invalid"}}}}`), &r); err != nil {
		t.Fatalf("cannot decode a 1.4 Repo: %v", err)
	}
	if d := r.Status.AtProvider.DomainList; d == nil || d.Public == nil || *d.Public != "registry.example.invalid" {
		t.Errorf("domainList = %+v; want the public domain 1.4 stored", d)
	}
}
