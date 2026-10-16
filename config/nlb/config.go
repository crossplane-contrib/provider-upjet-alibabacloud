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
		// vpc_id and security_group_ids are top-level fields, so KnownReferences()
		// in config/overrides.go wires them. It only walks the top-level schema
		// map, though, so the two fields nested under zone_mappings are declared
		// explicitly here: otherwise their refs would come only from the scraped
		// registry example in config/provider-metadata.yaml and would silently
		// disappear if a provider bump changed that example.
		r.References["zone_mappings.vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		// zone_id is not inferrable by name, so it resolves to the vswitch's
		// observed zoneId -- same as alicloud_alb_load_balancer.
		r.References["zone_mappings.zone_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathVSwitchZoneIdExtractor,
		}
		// Both the flat protection fields and their replacement blocks
		// (deletion_protection_config, modification_protection_config) are
		// Optional+Computed, so upjet would late-initialize both forms into
		// spec.forProvider and they would then compete with each other. The
		// flat ones are dropped during code generation (see
		// generationOnlySchemaOmissions in config/schema_omissions.go) so only
		// the canonical blocks are managed: upstream appends the flat fields
		// after zone_mappings, outside the alphabetical argument list, and the
		// blocks cover every field they do (enabled/reason, status/reason)
		// plus a read-only enabled_time.
	})
}
