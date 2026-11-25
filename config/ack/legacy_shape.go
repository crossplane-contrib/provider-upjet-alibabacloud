// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package ack

import (
	"context"
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/config"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-alibabacloud/internal/legacyshape"
)

// clusterRootGroup is the API group of the cluster-scoped resources, the only
// ones provider 1.x could have created.
const clusterRootGroup = "alibabacloud.crossplane.io"

// guardLegacyShape makes the resource's first initializer report it
// Synced=False while it still holds fields in the provider 1.x object form;
// see package legacyshape. Decoding drops those fields, so the check reads the
// resource as stored. It runs before anything that could write the decoded
// resource back, and so lose a 1.x spec.forProvider.runtime.
func guardLegacyShape(p *config.Provider, r *config.Resource) {
	if p.RootGroup != clusterRootGroup {
		return
	}
	gvk := schema.GroupVersionKind{Group: r.ShortGroup + "." + p.RootGroup, Version: r.Version, Kind: r.Kind}
	paths := legacyshape.Paths[r.Kind]
	r.InitializerFns = append([]config.NewInitializerFn{func(c client.Client) managed.Initializer {
		return managed.InitializerFn(func(ctx context.Context, mg resource.Managed) error {
			return checkLegacyShape(ctx, c, gvk, mg, paths)
		})
	}}, r.InitializerFns...)
}

func checkLegacyShape(ctx context.Context, c client.Client, gvk schema.GroupVersionKind, mg resource.Managed, paths []string) error {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	if err := c.Get(ctx, client.ObjectKeyFromObject(mg), u); err != nil {
		return errors.Wrap(err, "cannot read the resource to check it for provider 1.x field shapes")
	}
	found := legacyshape.Find(u.Object, paths)
	if len(found) == 0 {
		return nil
	}
	return errors.Errorf("%s: provider 1.x stored these as objects, and since 2.0 they are lists, as in Terraform provider 2.0. "+
		"Run %s --apply to convert this resource (a runtime needs the values its save step recorded before the upgrade), "+
		"and change manifests that set runtime to a list; see the 2.0.0 release notes",
		strings.Join(found, ", "), legacyshape.MigrationScript)
}
