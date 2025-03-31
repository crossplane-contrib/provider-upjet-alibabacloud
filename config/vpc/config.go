package vpc

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

const Vpc = "vpc"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_vpc", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = Vpc
		// Delete deprecated fields
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "router_table_id")
		delete(r.TerraformResource.Schema, "secondary_cidr_blocks")
	})
	p.AddResourceConfigurator("alicloud_havip", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Vpc
		r.ShortGroup = Vpc

		delete(r.TerraformResource.Schema, "havip_id")
	})
	p.AddResourceConfigurator("alicloud_haveip_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Vpc
		r.Kind = "HaveIpAttachment"
		r.ShortGroup = Vpc
	})
	p.AddResourceConfigurator("alicloud_network_acl", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Vpc
		r.ShortGroup = Vpc
	})
	p.AddResourceConfigurator("alicloud_network_acl_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Vpc
		r.ShortGroup = Vpc
	})
	p.AddResourceConfigurator("alicloud_network_acl_entries", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Vpc
		r.ShortGroup = Vpc
	})
	p.AddResourceConfigurator("alicloud_route_entry", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Vpc
		r.ShortGroup = Vpc
	})
	p.AddResourceConfigurator("alicloud_route_table", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Vpc
		r.ShortGroup = Vpc
	})
	p.AddResourceConfigurator("alicloud_route_table_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Vpc
		r.ShortGroup = Vpc
	})
	p.AddResourceConfigurator("alicloud_vswitch", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Vpc
		r.ShortGroup = Vpc
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
			Extractor:     common.PathIdExtractor,
		}
		// Delete deprecated fields
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "availability_zone")
	})
}
