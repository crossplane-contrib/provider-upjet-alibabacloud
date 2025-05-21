// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types"
	"github.com/crossplane/upjet/pkg/types/comments"
	"github.com/crossplane/upjet/pkg/types/name"
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
					Extractor:     common.PathRolaArnExtractor,
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
					TerraformName:     "alicloud_vswitch",
					RefFieldName:      "VswitchIDRefs",
					SelectorFieldName: "VswitchIDSelector",
				}
			case "vswitch_id":
				r.References["vswitch_id"] = config.Reference{
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
