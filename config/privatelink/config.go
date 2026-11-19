package privatelink

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
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
	p.AddResourceConfigurator("alicloud_privatelink_vpc_endpoint_service", func(r *config.Resource) {
		// Group and Kind are left to upjet's defaults, which already give
		// privatelink/VPCEndpointService.
		//
		// Terraform provider 1.293.0 added an Optional+Computed "resource"
		// block that Read fills from ListVpcEndpointServiceResources and
		// Update reconciles against spec. Attachments owned by separate
		// VpcEndpointServiceResource managed resources would be late-inited
		// into this resource's spec, leaving two controllers writing the same
		// association. Keep it observation-only.
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"resource",
			},
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
		// Currently only alb is supported caused by the limitation of crossplane. More load balancer types if want to
		// be supported, there are two ways:
		// 1. waiting for crossplane supports querying and filtering the specified resource, see https://github.com/crossplane/crossplane/blob/main/design/design-doc-observe-only-resources.md#querying-and-filtering
		// 2. implement a new kind, like VpcEndpointServiceResourceSlb, VpcEndpointServiceResourceNlb and so on.
		r.References["resource_id"] = config.Reference{
			TerraformName: "alicloud_alb_load_balancer",
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
