package ecs

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_security_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.ShortGroup = string(common.ECS)
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		// Delete deprecated fields
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "inner_access")
	})
}
