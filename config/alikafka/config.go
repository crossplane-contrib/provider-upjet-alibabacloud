// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package alikafka

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

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
		// settling. Drop the dead field; "partition_num" fully replaces it.
		delete(r.TerraformResource.Schema, "topic_quota")
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
		// spec and re-send it on the next apply, fighting "remark". Drop the dead
		// field; "remark" replaces it.
		delete(r.TerraformResource.Schema, "description")
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
		// acl_resource_name is intentionally left unreferenced: it is
		// polymorphic (topic name, group id, cluster name, or transaction id)
		// and may be an asterisk, so it cannot be tied to a single resource
		// kind.
	})

	p.AddResourceConfigurator("alicloud_alikafka_instance_allowed_ip_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIKAFKA)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_alikafka_instance",
			Extractor:     common.PathIdExtractor,
		}
	})
}
