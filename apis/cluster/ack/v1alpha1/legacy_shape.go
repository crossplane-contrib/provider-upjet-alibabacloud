// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/util/json"

	"github.com/crossplane-contrib/provider-alibabacloud/internal/legacyshape"
)

// These kinds drop the fields that provider 1.x stored as objects when they are
// decoded, so that a resource 1.x created does not stop the provider; see
// package legacyshape. The provider's legacy shape check then reports it
// Synced=False until it is migrated.
//
// The hooks are on the kinds, not on their parameter and observation types,
// because upjet decodes those with Terraform's field names, which a custom
// decoder would bypass.

// UnmarshalJSON decodes an EdgeKubernetes, dropping 1.x object-form fields.
func (mg *EdgeKubernetes) UnmarshalJSON(data []byte) error {
	type plain EdgeKubernetes
	return unmarshalDroppingLegacyShape(data, (*plain)(mg), "EdgeKubernetes")
}

// UnmarshalJSON decodes a Kubernetes, dropping 1.x object-form fields.
func (mg *Kubernetes) UnmarshalJSON(data []byte) error {
	type plain Kubernetes
	return unmarshalDroppingLegacyShape(data, (*plain)(mg), "Kubernetes")
}

// UnmarshalJSON decodes a ManagedKubernetes, dropping 1.x object-form fields.
func (mg *ManagedKubernetes) UnmarshalJSON(data []byte) error {
	type plain ManagedKubernetes
	return unmarshalDroppingLegacyShape(data, (*plain)(mg), "ManagedKubernetes")
}

func unmarshalDroppingLegacyShape(data []byte, into any, kind string) error {
	data, err := legacyshape.Drop(data, legacyshape.Paths[kind])
	if err != nil {
		return err
	}
	// The same decoder Kubernetes uses: case-sensitive keys, and integers kept
	// as integers.
	return json.Unmarshal(data, into)
}
