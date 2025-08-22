package slb

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_slb_acl", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "entry_list")
	})
	p.AddResourceConfigurator("alicloud_slb_load_balancer", func(r *config.Resource) {
		r.References["master_zone_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathVSwitchZoneIdExtractor,
		}
		r.References["slave_zone_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathVSwitchZoneIdExtractor,
		}
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "specification")
	})
}
