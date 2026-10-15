package alidns

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// p.AddResourceConfigurator("alicloud_alidns_access_strategy", func(r *config.Resource) {
	//	r.ShortGroup = string(common.ALIDNS)
	//	r.Kind = "AccessStrategy"
	//	r.References["instance_id"] = config.Reference{
	//		TerraformName: "alicloud_alidns_gtm_instance",
	//	}
	//	r.References["failover_addr_pools.addr_pool_id"] = config.Reference{
	//		TerraformName: "alicloud_alidns_address_pool",
	//	}
	// })
	p.AddResourceConfigurator("alicloud_alidns_address_pool", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIDNS)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_alidns_gtm_instance",
		}
		r.References["alert_group.*"] = config.Reference{
			TerraformName: "alicloud_cms_alarm_contact_group",
			Extractor:     common.PathAlarmContactGroupNameExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_alidns_custom_line", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIDNS)
		r.References["domain_name"] = config.Reference{
			TerraformName: "alicloud_alidns_domain",
		}
	})
	p.AddResourceConfigurator("alicloud_alidns_domain", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIDNS)
		r.References["group_id"] = config.Reference{
			TerraformName: "alicloud_alidns_domain_group",
		}
	})
	p.AddResourceConfigurator("alicloud_alidns_domain_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIDNS)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_alidns_instance",
		}
	})
	p.AddResourceConfigurator("alicloud_alidns_domain_group", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIDNS)
	})
	p.AddResourceConfigurator("alicloud_alidns", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIDNS)
	})
	p.AddResourceConfigurator("alicloud_alidns_monitor_config", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIDNS)
		r.References["addr_pool_id"] = config.Reference{
			TerraformName: "alicloud_alidns_address_pool",
		}
	})
	p.AddResourceConfigurator("alicloud_alidns_record", func(r *config.Resource) {
		r.ShortGroup = string(common.ALIDNS)
		r.References["domain_name"] = config.Reference{
			TerraformName: "alicloud_alidns_domain",
		}
	})
}
