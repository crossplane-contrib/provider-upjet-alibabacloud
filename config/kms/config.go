package kms

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_kms_alias", func(r *config.Resource) {
		r.ShortGroup = string(common.KMS)
		r.References["key_id"] = config.Reference{
			TerraformName: "alicloud_kms_key",
		}
	})
	p.AddResourceConfigurator("alicloud_kms_instance", func(r *config.Resource) {
		r.ShortGroup = string(common.KMS)
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "VswitchRefs",
			SelectorFieldName: "VswitchSelector",
		}
		r.References["bind_vpcs.vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["bind_vpcs.vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
	})
	p.AddResourceConfigurator("alicloud_kms_key", func(r *config.Resource) {
		r.ShortGroup = string(common.KMS)

		delete(r.TerraformResource.Schema, "deletion_window_in_days")
		delete(r.TerraformResource.Schema, "is_enabled")
		delete(r.TerraformResource.Schema, "key_state")
	})
	p.AddResourceConfigurator("alicloud_kms_secret", func(r *config.Resource) {
		r.ShortGroup = string(common.KMS)
		r.References["encryption_key_id"] = config.Reference{
			TerraformName: "alicloud_kms_key",
		}
		r.References["dkms_instance_id"] = config.Reference{
			TerraformName: "alicloud_kms_instance",
		}
	})
}
