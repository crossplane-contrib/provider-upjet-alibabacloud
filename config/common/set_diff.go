// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"fmt"
	"sort"
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// SetElementsKeyedBy returns a resource option that suppresses the element
// churn the plugin SDK reports for the TypeSet field when its elements in the
// configuration and in the prior state carry the same values for keys.
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
// When the key values match, the elements are the same and the churn is
// dropped. A change to any key field still produces a diff. Any custom diff
// already configured on the resource runs first.
func SetElementsKeyedBy(field string, keys ...string) config.ResourceOption {
	return func(r *config.Resource) {
		next := r.TerraformCustomDiff
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, cfg *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if next != nil {
				var err error
				if diff, err = next(diff, state, cfg); err != nil {
					return nil, err
				}
			}
			suppressSetElementChurn(field, keys, diff, state, cfg)
			return diff, nil
		}
	}
}

func suppressSetElementChurn(field string, keys []string, diff *terraform.InstanceDiff, state *terraform.InstanceState, cfg *terraform.ResourceConfig) {
	if diff == nil || state == nil || state.ID == "" || cfg == nil {
		return
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
		return
	}
	old, ok := stateSetKeys(state.Attributes, field, keys)
	if !ok {
		return
	}
	cur, ok := configSetKeys(cfg, field, keys)
	if !ok || !equalKeys(old, cur) {
		return
	}
	for k := range diff.Attributes {
		if strings.HasPrefix(k, prefix) {
			delete(diff.Attributes, k)
		}
	}
}

// stateSetKeys returns the sorted key tuples of the set elements in a flatmap
// state, whose element attributes are "<field>.<hash>.<key>".
func stateSetKeys(attrs map[string]string, field string, keys []string) ([]string, bool) {
	prefix := field + "."
	elems := map[string]map[string]string{}
	for k, v := range attrs {
		rest, ok := strings.CutPrefix(k, prefix)
		if !ok {
			continue
		}
		hash, name, ok := strings.Cut(rest, ".")
		if !ok {
			continue
		}
		if elems[hash] == nil {
			elems[hash] = map[string]string{}
		}
		elems[hash][name] = v
	}
	out := make([]string, 0, len(elems))
	for _, e := range elems {
		out = append(out, keyTuple(keys, func(k string) (string, bool) { v, ok := e[k]; return v, ok }))
	}
	sort.Strings(out)
	return out, true
}

// configSetKeys returns the sorted key tuples of the set elements in the
// configuration, or false when the field is not a list of objects there.
func configSetKeys(cfg *terraform.ResourceConfig, field string, keys []string) ([]string, bool) {
	raw, ok := cfg.Get(field)
	if !ok {
		return nil, false
	}
	list, ok := raw.([]interface{})
	if !ok {
		return nil, false
	}
	out := make([]string, 0, len(list))
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			return nil, false
		}
		out = append(out, keyTuple(keys, func(k string) (string, bool) {
			v, ok := m[k]
			if !ok || v == nil {
				return "", false
			}
			return fmt.Sprint(v), true
		}))
	}
	sort.Strings(out)
	return out, true
}

func keyTuple(keys []string, get func(string) (string, bool)) string {
	parts := make([]string, len(keys))
	for i, k := range keys {
		v, _ := get(k)
		parts[i] = k + "=" + v
	}
	return strings.Join(parts, "\x00")
}

func equalKeys(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
