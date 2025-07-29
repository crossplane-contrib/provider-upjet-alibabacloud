package fcv3

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_fcv3_alias", func(r *config.Resource) {
		r.ShortGroup = string(common.FCV3)
		r.Kind = "Alias"
		r.References["function_name"] = config.Reference{
			TerraformName: "alicloud_fcv3_function",
		}
	})
	p.AddResourceConfigurator("alicloud_fcv3_async_invoke_config", func(r *config.Resource) {
		r.ShortGroup = string(common.FCV3)
		r.Kind = "AsyncInvokeConfig"
		r.References["function_name"] = config.Reference{
			TerraformName: "alicloud_fcv3_function",
		}
		r.References["destination_config.on_success.destination"] = config.Reference{
			TerraformName: "alicloud_fcv3_function",
			Extractor:     common.PathFcv3FunctionArnExtractor,
		}
		r.References["destination_config.on_failure.destination"] = config.Reference{
			TerraformName: "alicloud_fcv3_function",
			Extractor:     common.PathFcv3FunctionArnExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_fcv3_concurrency_config", func(r *config.Resource) {
		r.ShortGroup = string(common.FCV3)
		r.Kind = "ConcurrencyConfig"
		r.References["function_name"] = config.Reference{
			TerraformName: "alicloud_fcv3_function",
		}
	})
	p.AddResourceConfigurator("alicloud_fcv3_custom_domain", func(r *config.Resource) {
		r.ShortGroup = string(common.FCV3)
		r.Kind = "CustomDomain"
		r.References["route_config.routes.function_name"] = config.Reference{
			TerraformName: "alicloud_fcv3_function",
		}
	})
	p.AddResourceConfigurator("alicloud_fcv3_concurrency_config", func(r *config.Resource) {
		r.ShortGroup = string(common.FCV3)
		r.Kind = "ConcurrencyConfig"
		r.References["function_name"] = config.Reference{
			TerraformName: "alicloud_fcv3_function",
		}
	})
	p.AddResourceConfigurator("alicloud_fcv3_function", func(r *config.Resource) {
		r.ShortGroup = string(common.FCV3)
		r.Kind = "Function"
		r.References["code.oss_bucket_name"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
		r.References["code.oss_object_name"] = config.Reference{
			TerraformName: "alicloud_oss_bucket_object",
		}
		r.References["role"] = config.Reference{
			TerraformName: "alicloud_ram_role",
		}
		r.References["vpc_config.vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["vpc_config.security_group_id"] = config.Reference{
			TerraformName: "alicloud_security_group",
		}
		r.References["vpc_config.vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "VSwitchIDRefs",
			SelectorFieldName: "VSwitchIDSelector",
		}
	})
	p.AddResourceConfigurator("alicloud_fcv3_function_version", func(r *config.Resource) {
		r.ShortGroup = string(common.FCV3)
		r.Kind = "FunctionVersion"
		r.References["function_name"] = config.Reference{
			TerraformName: "alicloud_fcv3_function",
		}
	})
	p.AddResourceConfigurator("alicloud_fcv3_layer_version", func(r *config.Resource) {
		r.ShortGroup = string(common.FCV3)
		r.Kind = "LayerVersion"
		r.References["code.oss_bucket_name"] = config.Reference{
			TerraformName: "alicloud_oss_bucket",
		}
		r.References["code.oss_object_name"] = config.Reference{
			TerraformName: "alicloud_oss_bucket_object",
		}
	})
	p.AddResourceConfigurator("alicloud_fcv3_provision_config", func(r *config.Resource) {
		r.ShortGroup = string(common.FCV3)
		r.Kind = "ProvisionConfig"
		r.References["function_name"] = config.Reference{
			TerraformName: "alicloud_fcv3_function",
		}
	})
	p.AddResourceConfigurator("alicloud_fcv3_trigger", func(r *config.Resource) {
		r.ShortGroup = string(common.FCV3)
		r.Kind = "Trigger"
		r.References["function_name"] = config.Reference{
			TerraformName: "alicloud_fcv3_function",
		}
		r.References["invocation_role"] = config.Reference{
			TerraformName: "alicloud_ram_role",
			Extractor:     common.PathRoleArnExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_fcv3_vpc_binding", func(r *config.Resource) {
		r.ShortGroup = string(common.FCV3)
		r.Kind = "VpcBinding"
		r.References["function_name"] = config.Reference{
			TerraformName: "alicloud_fcv3_function",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
	})
}
