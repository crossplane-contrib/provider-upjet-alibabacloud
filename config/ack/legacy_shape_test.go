// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package ack

import (
	"context"
	"strings"
	"testing"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	clientfake "sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/crossplane-contrib/provider-alibabacloud/internal/legacyshape"
)

func TestCheckLegacyShape(t *testing.T) {
	gvk := schema.GroupVersionKind{Group: "ack." + clusterRootGroup, Version: "v1alpha1", Kind: "Kubernetes"}
	legacy := map[string]any{
		"spec":   map[string]any{"forProvider": map[string]any{"runtime": map[string]any{"name": "containerd"}}},
		"status": map[string]any{"atProvider": map[string]any{"connections": map[string]any{"apiServerInternet": "x"}}},
	}
	cases := map[string]struct {
		stored   map[string]any
		deleting bool
		wantErr  []string
	}{
		"1.xRuntime": {
			stored:  legacy,
			wantErr: []string{"spec.forProvider.runtime, status.atProvider.connections", legacyshape.MigrationScript},
		},
		// The reconciler initializes before it deletes, so a resource being
		// deleted must get through, or its cloud resource is never deleted.
		"1.xRuntimeBeingDeleted": {
			stored:   legacy,
			deleting: true,
		},
		"2.0Lists": {
			stored: map[string]any{
				"spec": map[string]any{"forProvider": map[string]any{"runtime": []any{map[string]any{"name": "containerd"}}}},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			u := &unstructured.Unstructured{Object: tc.stored}
			u.SetGroupVersionKind(gvk)
			u.SetName("k")
			c := clientfake.NewClientBuilder().WithScheme(runtime.NewScheme()).WithObjects(u).Build()
			mg := &fake.Managed{}
			mg.SetName("k")
			if tc.deleting {
				now := metav1.Now()
				mg.SetDeletionTimestamp(&now)
			}

			err := checkLegacyShape(context.Background(), c, gvk, mg, legacyshape.Paths["Kubernetes"])
			if len(tc.wantErr) == 0 {
				if err != nil {
					t.Errorf("checkLegacyShape(...): unexpected error: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatal("checkLegacyShape(...): want an error for 1.x object-form fields")
			}
			for _, w := range tc.wantErr {
				if !strings.Contains(err.Error(), w) {
					t.Errorf("checkLegacyShape(...) = %q, want it to mention %q", err, w)
				}
			}
		})
	}
}
