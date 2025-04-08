package oss

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_oss_access_point", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
		r.References["vpc_configuration.vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_account_public_access_block", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		delete(r.TerraformResource.Schema, "acl")
		delete(r.TerraformResource.Schema, "logging_isenable")
		delete(r.TerraformResource.Schema, "referer_config")
		delete(r.TerraformResource.Schema, "policy")
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_access_monitor", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_acl", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_cname", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_cname_token", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_cors", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_data_redundancy_transition", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_https_config", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_logging", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
		r.References["target_bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_meta_query", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_object", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_policy", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_public_access_block", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_referer", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_replication", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
		r.References["destination.bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
		r.References["destination.location"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
			Extractor:     common.PathOssBucketLocationExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_request_payment", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_server_side_encryption", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_style", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_transfer_acceleration", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_user_defined_log_fields", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_versioning", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_website", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
	p.AddResourceConfigurator("alicloud_oss_bucket_worm", func(r *config.Resource) {
		r.ShortGroup = string(common.OSS)
		r.References["bucket"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
	})
}
