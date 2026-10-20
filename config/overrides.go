// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types"
	"github.com/crossplane/upjet/pkg/types/comments"
	"github.com/crossplane/upjet/pkg/types/name"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
)

// RegionAddition adds region to the spec of all resources except ram group which
// does not have a region notion.
func RegionAddition() config.ResourceOption { //nolint:gocyclo
	return func(r *config.Resource) {
		if r.ShortGroup == "ram" {
			return
		}
		c := "Region is the region you'd like your resource to be created in.\n"
		comment, err := comments.New(c, comments.WithTFTag("-"))
		if err != nil {
			panic(errors.Wrap(err, "cannot build comment for region"))
		}

		// check if the underlying Terraform resource already has "region"
		// as a (state) attribute
		if s, ok := r.TerraformResource.Schema["region"]; ok && types.IsObservation(s) {
			r.SchemaElementOptions.SetAddToObservation("region")
		}

		r.TerraformResource.Schema["region"] = &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: comment.String(),
		}
		if r.MetaResource == nil {
			return
		}
		for _, ex := range r.MetaResource.Examples {
			defaultRegion := "cn-beijing"
			if err := ex.SetPathValue("region", defaultRegion); err != nil {
				panic(err)
			}
			for k := range ex.Dependencies {
				if strings.HasPrefix(k, "alicloud_ram") {
					continue
				}
				if err := ex.Dependencies.SetPathValue(k, "region", defaultRegion); err != nil {
					panic(err)
				}
			}
		}
	}
}

// IdentifierAssignedByAlibabaCloud will work for all AlibabaCloud resource types because even if the ID
// is assigned by user, we'll see it in the Terraform State ID.
// The resource-specific configurations should override this whenever possible.
func IdentifierAssignedByAlibabaCloud() config.ResourceOption {
	return func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	}
}

// NamePrefixRemoval makes sure we remove name_prefix from all since it is mostly
// for Terraform functionality.
func NamePrefixRemoval() config.ResourceOption {
	return func(r *config.Resource) {
		for _, f := range r.ExternalName.OmittedFields {
			if f == "name_prefix" {
				return
			}
		}
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
	}
}

// KnownReferences adds references for fields that are known and common among
// more than a few resources.
func KnownReferences() config.ResourceOption { //nolint:gocyclo
	return func(r *config.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add references for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}
			switch {
			case strings.HasSuffix(k, "role_arn"):
				r.References[k] = config.Reference{
					TerraformName: "alicloud_ram_role",
					Extractor:     common.PathRoleArnExtractor,
				}
			case strings.HasSuffix(k, "security_group_ids"):
				r.References[k] = config.Reference{
					TerraformName:     "alicloud_security_group",
					RefFieldName:      name.NewFromSnake(strings.TrimSuffix(k, "s")).Camel + "Refs",
					SelectorFieldName: name.NewFromSnake(strings.TrimSuffix(k, "s")).Camel + "Selector",
				}
			}
			switch k {
			case "vpc_id":
				r.References["vpc_id"] = config.Reference{
					TerraformName: "alicloud_vpc",
				}
			case "vswitch_ids":
				r.References["vswitch_ids"] = config.Reference{
					TerraformName: "alicloud_vswitch",
					// Added s because of ids versus id
					RefFieldName: "VswitchIDsRefs",
					// Added s because of ids versus id
					SelectorFieldName: "VswitchIDsSelector",
				}
			case "vswitch_id":
				r.References["vswitch_id"] = config.Reference{
					// Implicitly creates VswitchIDRefs and VswitchIDSelector
					// without s
					TerraformName: "alicloud_vswitch",
				}
			case "ram_roles":
				r.References["ram_roles"] = config.Reference{
					TerraformName:     "alicloud_ram_role",
					RefFieldName:      "RAMRoleRefs",
					SelectorFieldName: "RAMRoleSelector",
				}
			case "security_group_id":
				r.References["security_group_id"] = config.Reference{
					TerraformName: "alicloud_security_group",
				}
			case "kms_key_id":
				r.References["kms_key_id"] = config.Reference{
					TerraformName: "alicloud_kms_key",
				}
			case "oss_bucket":
				r.References["oss_bucket"] = config.Reference{
					TerraformName: "alicloud_oss_bucket",
				}
			case "bucket_name":
				r.References["bucket_name"] = config.Reference{
					TerraformName: "alicloud_oss_bucket",
				}
			}
		}
	}
}

// AddExternalTagsField adds ExternalTagsFieldName configuration for resources that have tags field.
func AddExternalTagsField() config.ResourceOption {
	return func(r *config.Resource) {
		if s, ok := r.TerraformResource.Schema["tags"]; ok && s.Type == schema.TypeMap {
			r.InitializerFns = append(r.InitializerFns, config.TagInitializer)
		}
	}
}

// backtickRun matches a run of backticks together with any spaces before it,
// so removing an empty span does not leave a space stranded before the
// following punctuation. The callback narrows this to runs of exactly two.
var backtickRun = regexp.MustCompile(" *`+")

// DocumentationForEmptyCodeSpans removes empty inline code spans from the
// scraped upstream documentation.
//
// Nine alicloud resources document their computed id as
//
//	The ID of the resource supplied above.The value is formulated as ``.
//
// where upstream left the inline code span empty (other resources carry
// something like <db_cluster_id>:<account_name> there). Those two backticks
// make the generated output depend on which Go toolchain built the goimports
// binary that upjet shells out to: typewriter emits field comments at column 1,
// so go/printer treats them as top-level doc comments and runs them through
// go/doc/comment, which rewrites an empty code span into a U+201C quote. Go
// 1.27's printer skips that for field comments, so the same source yields
// different bytes under different toolchains and check-diff fails depending on
// how goimports was compiled.
//
// Dropping the empty span removes the trigger, so generation is reproducible
// under every toolchain. Runs of three or more backticks are left alone so
// fenced blocks and real code spans are unaffected, as are single backticks,
// which is what a non-empty span is made of.
//
// Only alicloud_hbase_instance.core_disk_type has the span mid-sentence rather
// than before a full stop; there the surrounding punctuation is left as the
// upstream docs wrote it. That resource is not generated here today.
//
// A two-backtick span with content, as in “a`b“, would lose its delimiters
// rather than being left alone. No description in provider-metadata.yaml uses
// that form -- every occurrence of two backticks there is an empty span -- so
// this is noted rather than handled.
func DocumentationForEmptyCodeSpans() config.ResourceOption {
	return func(r *config.Resource) {
		if r.MetaResource == nil {
			return
		}
		for k, v := range r.MetaResource.ArgumentDocs {
			if !strings.Contains(v, "``") {
				continue
			}
			r.MetaResource.ArgumentDocs[k] = backtickRun.ReplaceAllStringFunc(v, func(match string) string {
				if strings.Count(match, "`") == 2 {
					return ""
				}
				return match
			})
		}
	}
}

// DocumentationForTags overrides the API documentation of the tags fields since
// it contains Terraform-specific feature call out.
func DocumentationForTags() config.ResourceOption {
	return func(r *config.Resource) {
		if r.MetaResource == nil {
			return
		}
		if _, ok := r.MetaResource.ArgumentDocs["tags"]; ok {
			r.MetaResource.ArgumentDocs["tags"] = "- (Optional) Key-value map of resource tags."
		}
	}
}

// UnknownComputedCountSuppression drops the "known after apply" count diffs
// the plugin SDK reports for an Optional+Computed collection whose count is
// absent from the prior state, on resources that already exist.
//
// Without a cached state -- after an import, or after a provider restart --
// upjet rebuilds the prior state from status.atProvider, or from
// spec.forProvider when that is empty. An empty collection survives neither:
// it is omitted from both, so the rebuilt state carries no "<field>.#" key at
// all, where a state produced by Create carries "<field>.#" = "0". Many
// upstream Read functions only d.Set a collection when it is non-empty
// (alicloud_db_instance never sets babelfish_config, and sets
// security_group_ids and parameters only conditionally), so the refresh does
// not restore the key either. schemaMap.diffSet then reports the count as
// NewComputed, which makes the resource permanently not up to date, and when
// the field is ForceNew Update refuses outright:
//
//	refuse to update the external resource because the following update
//	requires replacing it: cannot change the value of the argument
//	"babelfish_config.#" from "" to ""
//
// An Optional+Computed field left out of the configuration means "keep what
// the provider has", so an unknown count with no configured value and no
// element-level diff is not a change. Diffs for creates, configured fields and
// element-level changes are left alone.
func UnknownComputedCountSuppression() config.ResourceOption {
	return func(r *config.Resource) {
		next := r.TerraformCustomDiff
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, cfg *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if next != nil {
				var err error
				if diff, err = next(diff, state, cfg); err != nil {
					return nil, err
				}
			}
			suppressUnknownComputedCounts(r.TerraformResource.Schema, diff, state, cfg)
			return diff, nil
		}
	}
}

func suppressUnknownComputedCounts(sm map[string]*schema.Schema, diff *terraform.InstanceDiff, state *terraform.InstanceState, cfg *terraform.ResourceConfig) { //nolint:gocyclo
	if diff == nil || state == nil || state.ID == "" {
		return
	}
	for k, d := range diff.Attributes {
		if d == nil || !d.NewComputed || d.Old != "" {
			continue
		}
		field, ok := strings.CutSuffix(k, ".#")
		if !ok {
			if field, ok = strings.CutSuffix(k, ".%"); !ok {
				continue
			}
		}
		if _, ok := state.Attributes[k]; ok {
			continue
		}
		if s := schemaAt(sm, field); s == nil || !s.Optional || !s.Computed {
			continue
		}
		if cfg != nil {
			if _, ok := cfg.Get(field); ok {
				continue
			}
		}
		if hasElementDiff(diff, field) {
			continue
		}
		delete(diff.Attributes, k)
	}
}

// schemaAt returns the schema of the flatmap attribute path p, such as
// "babelfish_config" or "network.0.ports", or nil if there is none.
func schemaAt(sm map[string]*schema.Schema, p string) *schema.Schema {
	var s *schema.Schema
	for _, seg := range strings.Split(p, ".") {
		if s != nil {
			if _, err := strconv.Atoi(seg); err == nil {
				continue
			}
			r, ok := s.Elem.(*schema.Resource)
			if !ok {
				return nil
			}
			sm = r.Schema
		}
		if s = sm[seg]; s == nil {
			return nil
		}
	}
	return s
}

// hasElementDiff reports whether diff changes any element of the collection
// field, as opposed to only its count.
func hasElementDiff(diff *terraform.InstanceDiff, field string) bool {
	for k := range diff.Attributes {
		if rest, ok := strings.CutPrefix(k, field+"."); ok && rest != "#" && rest != "%" {
			return true
		}
	}
	return false
}
