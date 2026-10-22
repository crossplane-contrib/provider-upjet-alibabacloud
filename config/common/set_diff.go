// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// SuppressSetChurn drops the element churn the plugin SDK reports for res's
// TypeSet fields when the configured elements are already present in the
// prior state.
//
// A set element's hash covers its Optional+Computed fields. Once the provider
// fills those in, an element read back into state hashes differently from the
// same element in the configuration, which leaves them empty, and the SDK
// reports every element as removed and re-added on every reconcile. Terraform
// itself never sees this: core merges the prior state's computed values into
// the planned configuration before the provider diffs, which the no-fork path
// does not do, and late-initialization does not fill fields inside elements
// the user already set.
//
// A configured element matches a state element when every field set in the
// configuration has the same value in state, recursing into nested blocks;
// fields left empty in the configuration are the provider's to fill. When each
// configured element matches a distinct state element and the counts agree,
// nothing the user asked for differs and the field's diff is dropped. Any
// other difference -- a changed or added value, an added or removed element --
// is kept.
func SuppressSetChurn(res *schema.Resource, diff *terraform.InstanceDiff, state *terraform.InstanceState, cfg *terraform.ResourceConfig) {
	if res == nil || diff == nil || state == nil || state.ID == "" || cfg == nil {
		return
	}
	var d *schema.ResourceData
	for field, s := range res.Schema {
		if s.Type != schema.TypeSet || !hasOptionalComputed(s) {
			continue
		}
		prefix := field + "."
		churn := false
		for k := range diff.Attributes {
			if strings.HasPrefix(k, prefix) {
				churn = true
				break
			}
		}
		if !churn {
			continue
		}
		want, ok := cfg.Get(field)
		if !ok {
			continue
		}
		if d == nil {
			d = res.Data(state)
		}
		if !matchValue(want, d.Get(field)) {
			continue
		}
		for k := range diff.Attributes {
			if strings.HasPrefix(k, prefix) {
				delete(diff.Attributes, k)
			}
		}
	}
}

// hasOptionalComputed reports whether a set's elements, at any depth, have an
// Optional+Computed field, the only kind that can cause the churn.
func hasOptionalComputed(s *schema.Schema) bool {
	r, ok := s.Elem.(*schema.Resource)
	if !ok {
		return false
	}
	for _, es := range r.Schema {
		if es.Optional && es.Computed {
			return true
		}
		if hasOptionalComputed(es) {
			return true
		}
	}
	return false
}

// matchValue reports whether the configured value want is satisfied by the
// state value got: maps match on the keys set in want, and lists and sets
// match element for element, in any order.
func matchValue(want, got interface{}) bool {
	if s, ok := got.(*schema.Set); ok {
		got = s.List()
	}
	switch w := want.(type) {
	case map[string]interface{}:
		g, ok := got.(map[string]interface{})
		if !ok {
			return false
		}
		for k, wv := range w {
			if isEmpty(wv) {
				continue
			}
			if !matchValue(wv, g[k]) {
				return false
			}
		}
		return true
	case []interface{}:
		g, ok := got.([]interface{})
		if !ok || len(g) != len(w) {
			return false
		}
		used := make([]bool, len(g))
		for _, wv := range w {
			found := false
			for i, gv := range g {
				if !used[i] && matchValue(wv, gv) {
					used[i] = true
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return true
	default:
		return fmt.Sprint(want) == fmt.Sprint(got)
	}
}

func isEmpty(v interface{}) bool {
	switch t := v.(type) {
	case nil:
		return true
	case string:
		return t == ""
	case []interface{}:
		return len(t) == 0
	case map[string]interface{}:
		return len(t) == 0
	}
	return false
}
