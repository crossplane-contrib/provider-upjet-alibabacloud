/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

// TestSchemaVersionsMatch guards against the Makefile's TERRAFORM_PROVIDER_VERSION
// and go.mod's replace pseudo-version for the upstream Terraform provider
// silently drifting apart.
//
// The two pins serve different purposes (TERRAFORM_PROVIDER_VERSION selects the
// registry release config/schema.json is generated from; the go.mod replace
// pins the actual Go code the provider links and runs) but must point at the
// same upstream commit. TestSchemaSourcesAgree catches the case where this
// drift is large enough that a configured resource is missing from one source
// or the other; this test catches the drift directly, by resolving
// TERRAFORM_PROVIDER_VERSION's tag against the upstream repository and
// comparing its commit to the one embedded in the pseudo-version.
//
// Requires network access to the upstream repository (a single `git
// ls-remote`, not a clone). Skips rather than fails when that is unavailable,
// so it does not break offline development runs.
func TestSchemaVersionsMatch(t *testing.T) {
	makefile, err := os.ReadFile("../Makefile")
	if err != nil {
		t.Fatalf("cannot read ../Makefile: %v", err)
	}

	version := findMakefileVar(t, string(makefile), "TERRAFORM_PROVIDER_VERSION")
	repo := findMakefileVar(t, string(makefile), "TERRAFORM_PROVIDER_REPO")

	goMod, err := os.ReadFile("../go.mod")
	if err != nil {
		t.Fatalf("cannot read ../go.mod: %v", err)
	}
	pseudoVersion := findGoModReplace(t, string(goMod), "github.com/aliyun/terraform-provider-alicloud")
	pinnedCommit := commitFromPseudoVersion(t, pseudoVersion)

	// A lightweight remote query, not a clone: resolves the annotated tag to
	// the commit it points at ("^{}" dereferences it; for a lightweight tag
	// this is a no-op and the same SHA is returned both with and without it).
	out, err := exec.Command("git", "ls-remote", repo, "refs/tags/v"+version+"^{}").Output() //nolint:gosec // repo and version both come from this repo's own tracked files
	if err != nil {
		t.Skipf("cannot resolve v%s against %s (no network access?): %v", version, repo, err)
	}
	fields := strings.Fields(string(out))
	if len(fields) == 0 {
		t.Skipf("v%s^{} did not resolve against %s; resolving the bare tag instead", version, repo)
	}
	tagCommit := fields[0]

	if !strings.HasPrefix(tagCommit, pinnedCommit) {
		t.Errorf("TERRAFORM_PROVIDER_VERSION (Makefile) is v%s, which resolves to commit %s, but go.mod's "+
			"replace pins pseudo-version %s (commit %s). The two must refer to the same upstream commit, "+
			"or config/schema.json and the Go code the provider actually runs will silently disagree.",
			version, tagCommit, pseudoVersion, pinnedCommit)
	}
}

func findMakefileVar(t *testing.T, makefile, name string) string {
	t.Helper()
	re := regexp.MustCompile(`(?m)^export ` + regexp.QuoteMeta(name) + ` \?= (\S+)`)
	m := re.FindStringSubmatch(makefile)
	if m == nil {
		t.Fatalf("cannot find %q in ../Makefile", name)
	}
	return m[1]
}

func findGoModReplace(t *testing.T, goMod, modulePath string) string {
	t.Helper()
	re := regexp.MustCompile(`(?m)^replace ` + regexp.QuoteMeta(modulePath) + ` => ` + regexp.QuoteMeta(modulePath) + ` (\S+)`)
	m := re.FindStringSubmatch(goMod)
	if m == nil {
		t.Fatalf("cannot find a replace directive for %q in ../go.mod", modulePath)
	}
	return m[1]
}

// commitFromPseudoVersion extracts the commit suffix from a Go pseudo-version
// such as v1.294.1-0.20261012080016-c7060a5b42ae.
func commitFromPseudoVersion(t *testing.T, pseudoVersion string) string {
	t.Helper()
	re := regexp.MustCompile(`-([0-9a-f]{12})$`)
	m := re.FindStringSubmatch(pseudoVersion)
	if m == nil {
		t.Fatalf("%q does not look like a Go pseudo-version ending in a 12-character commit hash", pseudoVersion)
	}
	return m[1]
}
