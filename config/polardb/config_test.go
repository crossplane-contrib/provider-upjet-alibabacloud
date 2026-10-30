// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package polardb

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestGoneWithCluster(t *testing.T) {
	cases := map[string]struct {
		id          string
		clusterGone bool
		wantChecked string
		wantRead    bool
		wantID      string
	}{
		"ClusterGone": {
			id:          "pc-rj95jayt9up7dox06:pe-rj9kf0rpzgewonyx5",
			clusterGone: true,
			wantChecked: "pc-rj95jayt9up7dox06",
			wantID:      "",
		},
		"ClusterExists": {
			id:          "pc-rj95jayt9up7dox06:pe-rj9kf0rpzgewonyx5",
			wantChecked: "pc-rj95jayt9up7dox06",
			wantRead:    true,
			wantID:      "pc-rj95jayt9up7dox06:pe-rj9kf0rpzgewonyx5",
		},
		// Without a cluster id there is nothing to check; upstream decides.
		"UnexpectedID": {
			id:          "pe-rj9kf0rpzgewonyx5",
			clusterGone: true,
			wantRead:    true,
			wantID:      "pe-rj9kf0rpzgewonyx5",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			d := (&schema.Resource{Schema: map[string]*schema.Schema{}}).Data(&terraform.InstanceState{ID: tc.id})
			checked, read := "", false
			r := goneWithCluster(func(*schema.ResourceData, any) error {
				read = true
				return nil
			}, func(_ any, clusterID string) bool {
				checked = clusterID
				return tc.clusterGone
			})
			if err := r(d, nil); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if checked != tc.wantChecked {
				t.Errorf("checked cluster %q, want %q", checked, tc.wantChecked)
			}
			if read != tc.wantRead {
				t.Errorf("upstream Read called: %t, want %t", read, tc.wantRead)
			}
			if d.Id() != tc.wantID {
				t.Errorf("id after Read %q, want %q", d.Id(), tc.wantID)
			}
		})
	}
}

// A resource that is being created is read at the end of upstream's Create;
// it must never be reported gone then.
func TestGoneWithClusterNewResource(t *testing.T) {
	d := (&schema.Resource{Schema: map[string]*schema.Schema{}}).Data(nil)
	d.MarkNewResource()
	d.SetId("pc-rj95jayt9up7dox06:pe-rj9kf0rpzgewonyx5")
	read := false
	r := goneWithCluster(func(*schema.ResourceData, any) error {
		read = true
		return nil
	}, func(any, string) bool { return true })
	if err := r(d, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !read || d.Id() == "" {
		t.Fatal("a new resource must go through upstream's Read")
	}
}

func TestGoneWithClusterNoRead(t *testing.T) {
	if goneWithCluster(nil, clusterNotFound) != nil {
		t.Fatal("wrapping a nil Read function must leave it nil")
	}
}

func TestClusterNotFoundWithoutClient(t *testing.T) {
	if clusterNotFound(nil, "pc-rj95jayt9up7dox06") {
		t.Fatal("without a client the cluster must not be reported gone")
	}
}
