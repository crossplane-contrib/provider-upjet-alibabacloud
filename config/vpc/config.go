package vpc

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_vpc", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "vpc"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"name",
			},
		}
		// Delete deprecated fields
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "router_table_id")
		delete(r.TerraformResource.Schema, "secondary_cidr_blocks")
	})

	p.AddResourceConfigurator("alicloud_vswitch", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = "vpc"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
			Extractor:     common.PathIdExtractor,
		}
		// Delete deprecated fields
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "availability_zone")
	})
}
