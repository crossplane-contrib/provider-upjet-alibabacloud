// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Package legacyshape finds the ACK fields that provider 1.x stored as objects
// and that 2.0 types as lists.
//
// Terraform provider 2.0 changed certificate_authority, connections and
// runtime from maps to lists of one object, because its plugin SDK v2 rejects
// maps of nested blocks. Provider 1.x kept them as objects, so cluster-scoped
// resources it created still hold objects where 2.0 expects lists. Such a
// resource cannot be decoded into the 2.0 types, and one that cannot be
// decoded stops its kind's cache and then the whole provider. The 2.0 types
// therefore drop these values when decoding, and the provider reports the
// resource Synced=False until it is migrated, rather than reinterpreting what
// 1.x stored.
package legacyshape

import (
	"bytes"
	"encoding/json"
	"strings"
)

// MigrationScript is the script that converts resources 1.x created.
const MigrationScript = "hack/migrate-ack-2.0.sh"

var status = []string{"status.atProvider.certificateAuthority", "status.atProvider.connections"}

var runtime = []string{"spec.forProvider.runtime", "spec.initProvider.runtime", "status.atProvider.runtime"}

// Paths are the fields of each cluster-scoped ACK kind that 1.x stored as
// objects and 2.0 types as lists.
var Paths = map[string][]string{
	"EdgeKubernetes":    append(append([]string{}, runtime...), status...),
	"Kubernetes":        append(append([]string{}, runtime...), status...),
	"ManagedKubernetes": status,
}

// Find returns the paths whose value in obj, a decoded JSON object, is an
// object rather than a list.
func Find(obj map[string]any, paths []string) []string {
	var found []string
	for _, p := range paths {
		parent, key := walk(obj, p)
		if _, ok := parent[key].(map[string]any); ok {
			found = append(found, p)
		}
	}
	return found
}

// Drop returns data with the values at paths removed wherever they are
// objects, or data itself when there are none. Numbers keep their exact
// representation.
func Drop(data []byte, paths []string) ([]byte, error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	var obj map[string]any
	if err := dec.Decode(&obj); err != nil {
		// Leave malformed input to the caller's decoder, which reports it.
		return data, nil //nolint:nilerr // the error is the decoder's to report
	}
	found := Find(obj, paths)
	if len(found) == 0 {
		return data, nil
	}
	for _, p := range found {
		parent, key := walk(obj, p)
		delete(parent, key)
	}
	return json.Marshal(obj)
}

// walk returns the object holding the last segment of path, and that segment;
// the object is nil when a segment before it is missing or not an object.
func walk(obj map[string]any, path string) (map[string]any, string) {
	segments := strings.Split(path, ".")
	parent := obj
	for _, s := range segments[:len(segments)-1] {
		next, ok := parent[s].(map[string]any)
		if !ok {
			return nil, segments[len(segments)-1]
		}
		parent = next
	}
	return parent, segments[len(segments)-1]
}
