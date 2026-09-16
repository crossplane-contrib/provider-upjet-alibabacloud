package ackone

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_ack_one_cluster", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ackone"
		r.ShortGroup = string(common.ACKONE)
		r.Kind = "Cluster"
		r.UseAsync = true
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["bind_vpcs.vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["network.vswitches"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
	})

	p.AddResourceConfigurator("alicloud_ack_one_membership_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ackone"
		r.ShortGroup = string(common.ACKONE)
		r.Kind = "MembershipAttachment"
	})
}
