package slb

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_slb_load_balancer", func(r *config.Resource) {
		r.References["master_zone_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathVSwitchZoneIdExtractor,
		}
		r.References["slave_zone_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathVSwitchZoneIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_slb_listener", func(r *config.Resource) {
		r.References["acl_ids"] = config.Reference{
			TerraformName: "alicloud_slb_acl",
			// RefFieldName:      "AclRefs",
			// SelectorFieldName: "AclSelector",
		}
	})
}
