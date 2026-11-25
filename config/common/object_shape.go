// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/name"
)

// KeepObjectShape generates the given top-level Terraform fields as objects
// instead of lists, the shape they had before Terraform provider 2.0.
//
// Provider 2.0 changed these fields from maps to lists of one object, because
// its plugin SDK v2 rejects maps of nested blocks. Objects that 1.4 stored keep
// the map in the cluster, and the generated list types cannot decode it, which
// stops the provider from starting. As objects they decode, and upjet converts
// them to and from Terraform's one-item list, as provider 2.0's own state
// upgrader does.
func KeepObjectShape(r *config.Resource, fields ...string) {
	for _, f := range fields {
		r.AddSingletonListConversion(f, name.NewFromSnake(f).LowerCamelComputed)
	}
	r.TerraformConversions = append(r.TerraformConversions, config.NewTFSingletonConversion())
}
