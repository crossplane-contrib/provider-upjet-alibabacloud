// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"regexp"
	"strings"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/types"
	"github.com/crossplane/upjet/v2/pkg/types/comments"
	"github.com/crossplane/upjet/v2/pkg/types/name"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
