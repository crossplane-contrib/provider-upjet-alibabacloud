package tair

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_kvstore_account", func(r *config.Resource) {
		r.ShortGroup = string(common.Tair)
		r.Kind = "Account"
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_kvstore_instance",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_kvstore_audit_log_config", func(r *config.Resource) {
		r.ShortGroup = string(common.Tair)
		r.Kind = "AuditLogConfig"
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_kvstore_instance",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_kvstore_connection", func(r *config.Resource) {
		r.ShortGroup = string(common.Tair)
		r.Kind = "Connection"
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_kvstore_instance",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_kvstore_instance", func(r *config.Resource) {
		r.ShortGroup = string(common.Tair)
		r.Kind = "Instance"
		r.References["security_group_id"] = config.Reference{
			TerraformName: "alicloud_security_group",
			Extractor:     common.PathIdExtractor,
		}
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathIdExtractor,
		}
		delete(r.TerraformResource.Schema, "availability_zone")
		delete(r.TerraformResource.Schema, "connection_string_prefix")
		delete(r.TerraformResource.Schema, "enable_public")
		delete(r.TerraformResource.Schema, "instance_charge_type")
		delete(r.TerraformResource.Schema, "instance_name")
		delete(r.TerraformResource.Schema, "node_type")
		delete(r.TerraformResource.Schema, "parameters")
	})
	p.AddResourceConfigurator("alicloud_redis_tair_instance", func(r *config.Resource) {
		r.ShortGroup = string(common.Tair)
		r.Kind = "TairInstance"
		r.References["security_group_id"] = config.Reference{
			TerraformName: "alicloud_security_group",
			Extractor:     common.PathIdExtractor,
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
			Extractor:     common.PathIdExtractor,
		}
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathIdExtractor,
		}
	})
}
