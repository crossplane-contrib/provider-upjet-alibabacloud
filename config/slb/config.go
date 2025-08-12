package slb

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_slb_acl", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "Acl"
		delete(r.TerraformResource.Schema, "entry_list")
	})
	p.AddResourceConfigurator("alicloud_slb_acl_entry_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "AclEntryAttachment"
		r.References["acl_id"] = config.Reference{
			TerraformName: "alicloud_slb_acl",
		}
	})
}
