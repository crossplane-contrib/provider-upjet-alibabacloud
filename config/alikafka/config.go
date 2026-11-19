// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package alikafka

import (
	"strings"

	"github.com/aliyun/terraform-provider-alicloud/alicloud"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// instanceChildren are the resources that live inside an alikafka instance and
// are released together with it.
var instanceChildren = []string{
	"alicloud_alikafka_topic",
	"alicloud_alikafka_consumer_group",
	"alicloud_alikafka_scheduled_scaling_rule",
	"alicloud_alikafka_sasl_user",
	"alicloud_alikafka_sasl_acl",
	"alicloud_alikafka_instance_allowed_ip_attachment",
}

// Configure configures individual resources by adding custom
// ResourceConfigurators. The Kind is intentionally left to upjet's default
// derivation from the Terraform resource name (e.g. "alicloud_alikafka_sasl_acl"
// -> "SaslACL"), matching how the other resource groups in this provider are
// configured.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_alikafka_instance", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIKAFKA)
		// "security_group" is the ECS security group id. KnownReferences only
		// wires "security_group_id"/"security_group_ids", so configure this one
		// explicitly to match the upstream example.
		r.References["security_group"] = config.Reference{
			TerraformName: "alicloud_security_group",
		}
		// "topic_quota" is deprecated in favor of "partition_num" and is marked
		// both Optional and Computed. The provider populates it on every read;
		// upjet then late-initializes it back into the spec and re-sends it on
		// the next apply alongside partition_num, which keeps the instance from
		// settling. Dropped during code generation; see generationOnlySchemaOmissions
		// in config/schema_omissions.go. "partition_num" fully replaces it.
	})

	p.AddResourceConfigurator("alicloud_alikafka_topic", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIKAFKA)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_alikafka_instance",
			Extractor:     common.PathIdExtractor,
		}
	})

	p.AddResourceConfigurator("alicloud_alikafka_consumer_group", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIKAFKA)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_alikafka_instance",
			Extractor:     common.PathIdExtractor,
		}
		// As of provider v1.268.0 "description" is deprecated in favor of the new
		// "remark" field and is now both Optional and Computed. Keeping it would
		// let upjet late-initialize the provider-populated value back into the
		// spec and re-send it on the next apply, fighting "remark". Dropped
		// during code generation; see generationOnlySchemaOmissions in
		// config/schema_omissions.go. "remark" replaces it.
	})

	p.AddResourceConfigurator("alicloud_alikafka_scheduled_scaling_rule", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIKAFKA)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_alikafka_instance",
			Extractor:     common.PathIdExtractor,
		}
	})

	p.AddResourceConfigurator("alicloud_alikafka_sasl_user", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIKAFKA)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_alikafka_instance",
			Extractor:     common.PathIdExtractor,
		}
	})

	p.AddResourceConfigurator("alicloud_alikafka_sasl_acl", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIKAFKA)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_alikafka_instance",
			Extractor:     common.PathIdExtractor,
		}
		// The SASL user's external-name id is "<instance_id>:<username>", so a
		// dedicated extractor is needed to resolve the bare username.
		r.References["username"] = config.Reference{
			TerraformName: "alicloud_alikafka_sasl_user",
			Extractor:     common.PathAliKafkaSaslUserUsernameExtractor,
		}
		// acl_resource_name is not configured here: upjet derives its
		// reference to Topic from the upstream documentation example
		// (alicloud_alikafka_topic.topic), which is what the ACL example uses.
		// The field is polymorphic though (topic name, group id, cluster name
		// or transaction id, and it may be an asterisk), so the reference is
		// only a convenience: set acl_resource_name directly for anything that
		// is not a topic managed by this provider.
	})

	p.AddResourceConfigurator("alicloud_alikafka_instance_allowed_ip_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIKAFKA)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_alikafka_instance",
			Extractor:     common.PathIdExtractor,
		}
	})
	for _, name := range instanceChildren {
		p.AddResourceConfigurator(name, func(r *config.Resource) {
			r.TerraformResource.Delete = goneWithReleasedInstance(r.TerraformResource.Delete) //nolint:staticcheck // upstream implements these resources with the legacy Delete field.
		})
	}
}

// goneWithReleasedInstance treats a delete that fails because the parent
// instance has already been released as a successful one: releasing the
// instance released the child with it.
//
// Upstream's describe calls map BIZ.INSTANCE.STATUS.ERROR to "not found" for
// some of these resources, but none of its delete calls do. Crossplane does not
// order deletes between managed resources, so when a Composition (or anything
// else) deletes an instance together with its children, the instance often
// goes first, every child delete then fails with that code, and the children
// keep their finalizers for good.
//
// Only the RELEASED state counts. The same code is returned while the instance
// is, for example, being upgraded, and a child must not be dropped then: upjet
// stops reading a resource once its delete has succeeded, so a wrongly
// swallowed error would orphan a child that still exists.
func goneWithReleasedInstance(del schema.DeleteFunc) schema.DeleteFunc { //nolint:staticcheck // upstream implements these resources with the legacy Delete field.
	if del == nil {
		// The code generation schema carries no CRUD functions.
		return nil
	}
	return func(d *schema.ResourceData, meta any) error {
		err := del(d, meta)
		if isInstanceReleased(err) {
			return nil
		}
		return err
	}
}

func isInstanceReleased(err error) bool {
	return alicloud.IsExpectedErrors(err, []string{"BIZ.INSTANCE.STATUS.ERROR", "BIZ_INSTANCE_STATUS_ERROR"}) &&
		strings.Contains(err.Error(), "instance is in RELEASED")
}
