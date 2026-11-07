// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"encoding/json"
	"testing"
)

func TestSecurityPreferenceMFAParameters(t *testing.T) {
	for _, tc := range []struct{ name, manifest string }{
		{"Adoption", `{"spec":{"forProvider":{}}}`},
		{"CanonicalField", `{"spec":{"forProvider":{"mfaOperationForLogin":"mandatory"}}}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pref := &SecurityPreference{}
			if err := json.Unmarshal([]byte(tc.manifest), pref); err != nil {
				t.Fatal(err)
			}
			// Include a legacy state value to ensure it cannot re-enter the spec.
			state := []byte(`{"enforce_mfa_for_login":true,"mfa_operation_for_login":"mandatory"}`)
			if _, err := pref.LateInitialize(state); err != nil {
				t.Fatal(err)
			}
			for _, merge := range []bool{false, true} {
				params, err := pref.GetMergedParameters(merge)
				if err != nil {
					t.Fatal(err)
				}
				if _, ok := params["enforce_mfa_for_login"]; ok {
					t.Fatal("unsupported MFA field remains in Terraform configuration")
				}
				if params["mfa_operation_for_login"] != "mandatory" {
					t.Fatalf("canonical MFA setting lost: %v", params)
				}
			}
			updated := "independent"
			pref.Spec.ForProvider.MfaOperationForLogin = &updated
			if _, err := pref.LateInitialize(state); err != nil {
				t.Fatal(err)
			}
			params, err := pref.GetParameters()
			if err != nil {
				t.Fatal(err)
			}
			if params["mfa_operation_for_login"] != updated {
				t.Fatalf("MFA update overwritten: %v", params)
			}
		})
	}
}
