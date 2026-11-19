// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config_test

import (
	"context"
	"strings"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry"

	providerconfig "github.com/crossplane-contrib/provider-alibabacloud/config"
)

// TestDocumentationForEmptyCodeSpans guards the scrub that keeps generated doc
// comments identical regardless of which Go toolchain compiled the goimports
// binary upjet shells out to. An empty inline code span left in a description
// is rewritten to U+201C by go/doc/comment under pre-1.27 toolchains and left
// alone under 1.27+, which made check-diff depend on the formatter rather than
// on the repository contents.
func TestDocumentationForEmptyCodeSpans(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		want string
	}{
		{
			name: "EmptySpanAndItsLeadingSpaceAreRemoved",
			in:   "- The ID of the resource supplied above.The value is formulated as ``.",
			want: "- The ID of the resource supplied above.The value is formulated as.",
		},
		{
			name: "EmptySpanWithoutLeadingSpace",
			in:   "- Valid values are local_ssd_pro，``, local_disk size is fixed.",
			want: "- Valid values are local_ssd_pro，, local_disk size is fixed.",
		},
		{
			name: "RealSingleBacktickSpanIsPreserved",
			in:   "- Set this to `true` to enable the feature.",
			want: "- Set this to `true` to enable the feature.",
		},
		{
			name: "FencedBlockIsPreserved",
			in:   "- Example:\n```\nfoo = bar\n```",
			want: "- Example:\n```\nfoo = bar\n```",
		},
		{
			name: "DescriptionWithoutBackticksIsUntouched",
			in:   "- (Optional) The name of the instance.",
			want: "- (Optional) The name of the instance.",
		},
		{
			name: "LegitimateSpaceBeforePeriodSurvives",
			in:   "- See the docs for `foo` . Also ``.",
			want: "- See the docs for `foo` . Also.",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := &config.Resource{
				MetaResource: &registry.Resource{
					ArgumentDocs: map[string]string{"field": tc.in},
				},
			}
			providerconfig.DocumentationForEmptyCodeSpans()(r)
			if got := r.MetaResource.ArgumentDocs["field"]; got != tc.want {
				t.Errorf("got  %q\nwant %q", got, tc.want)
			}
		})
	}

	t.Run("NilMetaResourceIsTolerated", func(t *testing.T) {
		r := &config.Resource{}
		providerconfig.DocumentationForEmptyCodeSpans()(r)
	})

	// The whole point is that no empty span reaches the generated comments, so
	// assert it across the real configured provider rather than only in table
	// cases.
	t.Run("NoEmptySpanSurvivesInAnyConfiguredResource", func(t *testing.T) {
		p, err := providerconfig.GetProvider(context.Background(), true)
		if err != nil {
			t.Fatalf("cannot build the generation provider configuration: %v", err)
		}
		for name, r := range p.Resources {
			if r.MetaResource == nil {
				continue
			}
			for field, doc := range r.MetaResource.ArgumentDocs {
				if strings.Contains(doc, "``") {
					t.Errorf("%s.%s still contains an empty code span: %q", name, field, doc)
				}
				// go/doc/comment rewrites '' to a right double quote in the
				// same toolchain-dependent way. Nothing carries one today;
				// fail here if a provider bump introduces one.
				if strings.Contains(doc, "''") {
					t.Errorf("%s.%s contains '', which go/doc/comment rewrites the same way: %q", name, field, doc)
				}
			}
		}
	})
}
