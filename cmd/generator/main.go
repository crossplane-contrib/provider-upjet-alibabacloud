/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/crossplane/upjet/v2/pkg/pipeline"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}
	p, err := config.GetProvider(context.Background(), true)
	if err != nil {
		panic(fmt.Sprintf("cannot initialize the provider configuration: %v", err))
	}
	dumpGeneratedResourceList(p, filepath.Join(absRootDir, "config", "generated.lst"))
	// Namespaced resource generation is deferred to a follow-up change; passing
	// nil keeps the existing cluster-scoped layout.
	pipeline.Run(p, nil, absRootDir)
}

// dumpGeneratedResourceList records the Terraform resource names that were
// generated. scripts/version_diff.py reads it to report native state schema
// version changes between two upstream provider versions; without it the
// schema-version-diff target has nothing to compare.
func dumpGeneratedResourceList(p *ujconfig.Provider, targetPath string) {
	generated := make([]string, 0, len(p.Resources))
	for name := range p.Resources {
		generated = append(generated, name)
	}
	sort.Strings(generated)
	// A flat JSON array does not need indenting, but the newlines keep
	// concurrent PRs adding non-adjacent resources from conflicting.
	buff, err := json.MarshalIndent(generated, "", "")
	if err != nil {
		panic(fmt.Sprintf("cannot marshal the generated resource list: %v", err))
	}
	if err := os.WriteFile(targetPath, buff, 0o600); err != nil {
		panic(fmt.Sprintf("cannot write the generated resource list to %s: %v", targetPath, err))
	}
}
