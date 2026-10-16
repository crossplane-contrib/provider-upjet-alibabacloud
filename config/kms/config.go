package kms

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
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
			TerraformName: "alicloud_vswitch",
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
