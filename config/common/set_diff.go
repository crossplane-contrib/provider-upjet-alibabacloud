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
// configuration has the same value in state, recursing into nested blocks. A
// field left empty in the configuration matches when it is Computed -- the
// provider's to fill -- or when state holds its zero value: late
// initialization cannot represent a false or 0 it observed, so it leaves an
// empty element (alicloud_instance image_options churned this way). An empty
// field facing a non-zero, non-Computed value is a real change and does not
// match. When each configured element matches a distinct state element and
// the counts agree, nothing the user asked for differs and the field's diff is
// dropped. Any other difference -- a changed or added value, an added or
// removed element -- is kept.
func SuppressSetChurn(res *schema.Resource, diff *terraform.InstanceDiff, state *terraform.InstanceState, cfg *terraform.ResourceConfig) { //nolint:gocyclo
	if res == nil || diff == nil || state == nil || state.ID == "" || cfg == nil {
		return
	}
	var d *schema.ResourceData
	for field, s := range res.Schema {
		elem, ok := s.Elem.(*schema.Resource)
		if s.Type != schema.TypeSet || !ok {
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
		if !matchValue(want, d.Get(field), elem) {
			continue
		}
		for k := range diff.Attributes {
			if strings.HasPrefix(k, prefix) {
				delete(diff.Attributes, k)
			}
		}
	}
}

// matchValue reports whether the configured value want is satisfied by the
// state value got. elem is the schema of the elements when want is a list or
// set of nested blocks, or of the block itself when want is a map; it is nil
// for scalars and lists of scalars. Lists and sets match element for element,
// in any order.
func matchValue(want, got interface{}, elem *schema.Resource) bool { //nolint:gocyclo
	if s, ok := got.(*schema.Set); ok {
		got = s.List()
	}
	switch w := want.(type) {
	case map[string]interface{}:
		g, ok := got.(map[string]interface{})
		if !ok {
			return false
		}
		if elem == nil {
			for k, wv := range w {
				if !isEmpty(wv) && !matchValue(wv, g[k], nil) {
					return false
				}
			}
			return true
		}
		for k, es := range elem.Schema {
			wv, gv := w[k], g[k]
			if isEmpty(wv) {
				if !es.Computed && !isZero(gv) {
					return false
				}
				continue
			}
			sub, _ := es.Elem.(*schema.Resource)
			if !matchValue(wv, gv, sub) {
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
				if !used[i] && matchValue(wv, gv, elem) {
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

// isZero reports whether a state value is absent or its type's zero value.
func isZero(v interface{}) bool {
	if s, ok := v.(*schema.Set); ok {
		return s.Len() == 0
	}
	switch t := v.(type) {
	case nil:
		return true
	case bool:
		return !t
	case int:
		return t == 0
	case float64:
		return t == 0
	}
	return isEmpty(v)
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
