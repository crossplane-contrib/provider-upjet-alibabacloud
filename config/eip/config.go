package eip

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_eip_address", func(r *config.Resource) {
		r.ShortGroup = string(common.EIP)
		r.Kind = "EIPAddress"
		// Delete deprecated fields
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "instance_charge_type")
	})
	p.AddResourceConfigurator("alicloud_eip_association", func(r *config.Resource) {
		r.ShortGroup = string(common.EIP)
		r.Kind = "EIPAssociation"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
	})
}
