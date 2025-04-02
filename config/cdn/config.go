package cdn

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_cdn_domain_config", func(r *config.Resource) {
		r.ShortGroup = string(common.CDN)
		r.Kind = "DomainConfig"
		r.References["domain_name"] = config.Reference{
			TerraformName: "alicloud_cdn_domain_new",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_cdn_domain_new", func(r *config.Resource) {
		r.ShortGroup = string(common.CDN)
		r.Kind = "Domain"
	})
	p.AddResourceConfigurator("alicloud_cdn_fc_trigger", func(r *config.Resource) {
		r.ShortGroup = string(common.CDN)
		r.Kind = "FcTrigger"
	})
}
