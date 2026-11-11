package privatezone

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_pvtz_endpoint", func(r *config.Resource) {
		r.ShortGroup = string(common.PrivateZone)
		r.Kind = "Endpoint"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["security_group_id"] = config.Reference{
			TerraformName: "alicloud_security_group",
		}
	})
	p.AddResourceConfigurator("alicloud_pvtz_rule", func(r *config.Resource) {
		r.ShortGroup = string(common.PrivateZone)
		r.Kind = "Rule"
		r.References["endpoint_id"] = config.Reference{
			TerraformName: "alicloud_pvtz_endpoint",
		}
	})
	p.AddResourceConfigurator("alicloud_pvtz_rule_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.PrivateZone)
		r.Kind = "RuleAttachment"
		r.References["rule_id"] = config.Reference{
			TerraformName: "alicloud_pvtz_rule",
		}
	})
	p.AddResourceConfigurator("alicloud_pvtz_user_vpc_authorization", func(r *config.Resource) {
		r.ShortGroup = string(common.PrivateZone)
		r.Kind = "UserVpcAuthorization"
	})
	p.AddResourceConfigurator("alicloud_pvtz_zone", func(r *config.Resource) {
		r.ShortGroup = string(common.PrivateZone)
		r.Kind = "Zone"
	})
	p.AddResourceConfigurator("alicloud_pvtz_zone_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.PrivateZone)
		r.Kind = "ZoneAttachment"
		r.References["zone_id"] = config.Reference{
			TerraformName: "alicloud_pvtz_zone",
		}
	})
	p.AddResourceConfigurator("alicloud_pvtz_zone_record", func(r *config.Resource) {
		r.ShortGroup = string(common.PrivateZone)
		r.Kind = "ZoneRecord"
		r.References["zone_id"] = config.Reference{
			TerraformName: "alicloud_pvtz_zone",
		}
	})
}
