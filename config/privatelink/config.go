package privatelink

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_privatelink_vpc_endpoint", func(r *config.Resource) {
		r.ShortGroup = string(common.PRIVATELINK)
		r.Kind = "VpcEndpoint"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["security_group_ids"] = config.Reference{
			TerraformName:     "alicloud_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["service_id"] = config.Reference{
			TerraformName: "alicloud_privatelink_vpc_endpoint_service",
		}
	})
	p.AddResourceConfigurator("alicloud_privatelink_vpc_endpoint_connection", func(r *config.Resource) {
		r.ShortGroup = string(common.PRIVATELINK)
		r.Kind = "VpcEndpointConnection"
		r.References["endpoint_id"] = config.Reference{
			TerraformName: "alicloud_privatelink_vpc_endpoint",
		}
		r.References["service_id"] = config.Reference{
			TerraformName: "alicloud_privatelink_vpc_endpoint_service",
		}
	})
	p.AddResourceConfigurator("alicloud_privatelink_vpc_endpoint_service_resource", func(r *config.Resource) {
		r.ShortGroup = string(common.PRIVATELINK)
		r.Kind = "VpcEndpointServiceResource"
		r.References["service_id"] = config.Reference{
			TerraformName: "alicloud_privatelink_vpc_endpoint_service",
		}
	})
	p.AddResourceConfigurator("alicloud_privatelink_vpc_endpoint_service_user", func(r *config.Resource) {
		r.ShortGroup = string(common.PRIVATELINK)
		r.Kind = "VpcEndpointServiceUser"
		r.References["service_id"] = config.Reference{
			TerraformName: "alicloud_privatelink_vpc_endpoint_service",
		}
		r.References["user_id"] = config.Reference{
			TerraformName: "alicloud_ram_user",
		}
	})
	p.AddResourceConfigurator("alicloud_privatelink_vpc_endpoint_zone", func(r *config.Resource) {
		r.ShortGroup = string(common.PRIVATELINK)
		r.Kind = "VpcEndpointZone"
		r.References["endpoint_id"] = config.Reference{
			TerraformName: "alicloud_privatelink_vpc_endpoint",
		}
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
	})
}
