// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package legacyshape

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFindAndDrop(t *testing.T) {
	cases := map[string]struct {
		in        string
		wantFound []string
		wantOut   string
	}{
		"1.5Object": {
			in:        `{"spec":{"forProvider":{"runtime":{"name":"containerd"},"size":3}},"status":{"atProvider":{"connections":{"apiServerInternet":"x"},"id":"c1"}}}`,
			wantFound: []string{"spec.forProvider.runtime", "status.atProvider.connections"},
			wantOut:   `{"spec":{"forProvider":{"size":3}},"status":{"atProvider":{"id":"c1"}}}`,
		},
		"1.4Map": {
			in:        `{"status":{"atProvider":{"certificateAuthority":{"client_cert":"c"}}}}`,
			wantFound: []string{"status.atProvider.certificateAuthority"},
			wantOut:   `{"status":{"atProvider":{}}}`,
		},
		"2.0List": {
			in: `{"spec":{"forProvider":{"runtime":[{"name":"containerd"}]}},"status":{"atProvider":{"connections":[{"apiServerInternet":"x"}]}}}`,
		},
		"Absent": {
			in: `{"spec":{"forProvider":{}}}`,
		},
		"LargeNumberKeepsItsDigits": {
			in:        `{"spec":{"forProvider":{"runtime":{"name":"x"},"bytes":9007199254740993}}}`,
			wantFound: []string{"spec.forProvider.runtime"},
			wantOut:   `{"spec":{"forProvider":{"bytes":9007199254740993}}}`,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var obj map[string]any
			if err := json.Unmarshal([]byte(tc.in), &obj); err != nil {
				t.Fatal(err)
			}
			if got := Find(obj, Paths["Kubernetes"]); !reflect.DeepEqual(got, tc.wantFound) {
				t.Errorf("Find(...) = %v, want %v", got, tc.wantFound)
			}
			out, err := Drop([]byte(tc.in), Paths["Kubernetes"])
			if err != nil {
				t.Fatalf("Drop(...): %v", err)
			}
			want := tc.wantOut
			if want == "" {
				want = tc.in
			}
			if string(out) != want {
				t.Errorf("Drop(...) =\n%s\nwant\n%s", out, want)
			}
		})
	}
}
