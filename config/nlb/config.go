package nlb

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_nlb_load_balancer", func(r *config.Resource) {
		r.ShortGroup = string(common.NLB)
		r.Kind = "LoadBalancer"
		// vpc_id, security_group_ids and zone_mappings.vswitch_id are resolved
		// automatically (KnownReferences() in config/overrides.go and upjet's
		// reference injector, which matches field names at any depth).
		// zone_id is not inferrable by name, so it is wired explicitly here to
		// the vswitch's observed zoneId -- same as alicloud_alb_load_balancer.
		r.References["zone_mappings.zone_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathVSwitchZoneIdExtractor,
		}
	})
}
