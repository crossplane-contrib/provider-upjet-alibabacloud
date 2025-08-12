package vpc

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_route_entry", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = string(common.VPC)
		r.Kind = "RouteEntry"
		// nexthop_id will auto resolve to ecs instance which cause cycle import
		delete(r.References, "nexthop_id")
	})
	p.AddResourceConfigurator("alicloud_route_table", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = string(common.VPC)
		r.Kind = "RouteTable"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		// Delete deprecated fields
		delete(r.TerraformResource.Schema, "name")
	})
	p.AddResourceConfigurator("alicloud_route_table_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = string(common.VPC)
		r.Kind = "RouteTableAttachment"
		r.References["route_table_id"] = config.Reference{
			TerraformName: "alicloud_route_table",
		}
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
	})
	p.AddResourceConfigurator("alicloud_vpc", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = string(common.VPC)
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"name",
			},
		}
	})
	p.AddResourceConfigurator("alicloud_vpc_gateway_endpoint", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = string(common.VPC)
		r.Kind = "GatewayEndpoint"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
	})
	p.AddResourceConfigurator("alicloud_vpc_gateway_endpoint_route_table_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = string(common.VPC)
		r.Kind = "GatewayEndpointRouteTableAttachment"
		r.References["route_table_id"] = config.Reference{
			TerraformName: "alicloud_route_table",
		}
		r.References["vpc_gateway_endpoint_id"] = config.Reference{
			TerraformName: "alicloud_vpc_gateway_endpoint",
		}
	})
	p.AddResourceConfigurator("alicloud_vpc_ipv4_gateway", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = string(common.VPC)
		r.Kind = "IPv4Gateway"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
	})
	p.AddResourceConfigurator("alicloud_vswitch", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = string(common.VPC)
		// Delete deprecated fields
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "availability_zone")
	})
}
