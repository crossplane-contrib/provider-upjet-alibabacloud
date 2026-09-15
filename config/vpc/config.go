package vpc

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_route_table", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = string(common.VPC)
		r.Kind = "RouteTable"
		// Delete deprecated fields
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
	p.AddResourceConfigurator("alicloud_vswitch", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = string(common.VPC)
		// Delete deprecated fields
	})
}
