// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"encoding/json"
	"testing"
)

// Terraform Read returns the deprecated aliases as well as the canonical fields.
// They must not become configuration arguments during adoption or later refreshes.
func TestRoleRefreshAfterAdoption(t *testing.T) {
	const roleName = "caip1267-poc-role"
	const policy = `{"Version":"1","Statement":[{"Effect":"Allow","Action":"sts:AssumeRole","Principal":{"Service":["ecs.aliyuncs.com"]}}]}`
	state := map[string]any{
		"id": roleName, "name": roleName, "role_name": roleName,
		"document": policy, "assume_role_policy_document": policy,
		"description": "original", "max_session_duration": 3600,
	}
	attrs, err := json.Marshal(state)
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name string
		spec map[string]any
	}{
		{name: "adoption", spec: map[string]any{}},
		{name: "canonical manifest", spec: map[string]any{
			"forProvider": map[string]any{"roleName": roleName, "assumeRolePolicyDocument": policy},
		}},
		{name: "previously late initialized aliases", spec: map[string]any{
			"forProvider":  map[string]any{"roleName": roleName, "assumeRolePolicyDocument": policy, "name": roleName, "document": policy},
			"initProvider": map[string]any{"name": roleName, "document": policy},
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			manifest, err := json.Marshal(map[string]any{"spec": tc.spec})
			if err != nil {
				t.Fatal(err)
			}
			role := &Role{}
			if err := json.Unmarshal(manifest, role); err != nil {
				t.Fatal(err)
			}
			if err := role.SetObservation(state); err != nil {
				t.Fatal(err)
			}
			if _, err := role.LateInitialize(attrs); err != nil {
				t.Fatal(err)
			}
			updated := "updated description"
			role.Spec.ForProvider.Description = &updated
			// A second observation must not undo the requested update or add aliases.
			if _, err := role.LateInitialize(attrs); err != nil {
				t.Fatal(err)
			}
			for _, mergeInit := range []bool{false, true} {
				params, err := role.GetMergedParameters(mergeInit)
				if err != nil {
					t.Fatal(err)
				}
				for _, alias := range []string{"name", "document"} {
					if _, ok := params[alias]; ok {
						t.Errorf("mergeInit=%t: conflicting alias %q in Terraform configuration", mergeInit, alias)
					}
				}
				for key, want := range map[string]any{"role_name": roleName, "assume_role_policy_document": policy, "description": updated} {
					if got := params[key]; got != want {
						t.Errorf("mergeInit=%t: %s = %v, want %v", mergeInit, key, got, want)
					}
				}
			}
			if role.GetID() != roleName {
				t.Errorf("adopted ID = %q, want %q", role.GetID(), roleName)
			}
		})
	}
}
