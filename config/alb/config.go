package alb

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_alb_acl", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "Acl"
		delete(r.TerraformResource.Schema, "acl_entries")
	})
	p.AddResourceConfigurator("alicloud_alb_acl_entry_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "AclEntryAttachment"
		r.References["acl_id"] = config.Reference{
			TerraformName: "alicloud_alb_acl",
		}
	})
	p.AddResourceConfigurator("alicloud_alb_ascript", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "Ascript"
		r.References["listener_id"] = config.Reference{
			TerraformName: "alicloud_alb_listener",
		}
	})
	p.AddResourceConfigurator("alicloud_alb_load_balancer", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "LoadBalancer"
		r.References["zone_mappings.zone_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathVSwitchZoneIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_alb_listener", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "Listener"
		r.References["load_balancer_id"] = config.Reference{
			TerraformName: "alicloud_alb_load_balancer",
		}
		r.References["security_policy_id"] = config.Reference{
			TerraformName: "alicloud_alb_security_policy",
		}
		// The listener's default certificate. ALB expects the composed
		// "<casCertId>-<region>" form here, same as the additional
		// certificate attachment below.
		r.References["certificates.certificate_id"] = config.Reference{
			TerraformName: "alicloud_ssl_certificates_service_certificate",
			Extractor:     common.PathAlbCertificateIdExtractor,
		}
		delete(r.TerraformResource.Schema, "acl_config")
		delete(r.TerraformResource.Schema, "xforwarded_for_config")
	})
	p.AddResourceConfigurator("alicloud_alb_listener_acl_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "ListenerAclAttachment"
		r.References["acl_id"] = config.Reference{
			TerraformName: "alicloud_alb_acl",
		}
		r.References["listener_id"] = config.Reference{
			TerraformName: "alicloud_alb_listener",
		}
	})
	//  p.AddResourceConfigurator("alicloud_alb_listener_additional_certificate_attachment", func(r *config.Resource) {
	//	r.ShortGroup = string(common.ALB)
	//	r.Kind = "ListenerAdditionalCertificateAttachment"
	//	r.References["certificate_id"] = config.Reference{
	//		TerraformName: "alicloud_ssl_certificate",
	//	}
	//	r.References["listener_id"] = config.Reference{
	//		TerraformName: "alicloud_alb_listener",
	//	}
	//  })
	//  p.AddResourceConfigurator("alicloud_alb_load_balancer_access_log_config_attachment", func(r *config.Resource) {
	//	r.ShortGroup = string(common.ALB)
	//	r.Kind = "LoadBalancerSecurityGroupAttachment"
	//	r.References["load_balancer_id"] = config.Reference{
	//		TerraformName: "alicloud_alb_load_balancer",
	//	}
	//	r.References["bandwidth_package_id"] = config.Reference{
	//		TerraformName: "alicloud_common_bandwidth_package",
	//	}
	//  })
	//  p.AddResourceConfigurator("alicloud_alb_load_balancer_common_bandwidth_package_attachment", func(r *config.Resource) {
	//	r.ShortGroup = string(common.ALB)
	//	r.Kind = "LoadBalancerSecurityGroupAttachment"
	//	r.References["load_balancer_id"] = config.Reference{
	//		TerraformName: "alicloud_alb_load_balancer",
	//	}
	//	r.References["bandwidth_package_id"] = config.Reference{
	//		TerraformName: "alicloud_common_bandwidth_package",
	//	}
	// })
	p.AddResourceConfigurator("alicloud_alb_load_balancer_security_group_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "LoadBalancerSecurityGroupAttachment"
		r.References["load_balancer_id"] = config.Reference{
			TerraformName: "alicloud_alb_load_balancer",
		}
		r.References["security_group_id"] = config.Reference{
			TerraformName: "alicloud_security_group",
		}
	})
	p.AddResourceConfigurator("alicloud_alb_load_balancer_zone_shifted_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "LoadBalancerZoneShiftedAttachment"
		r.References["load_balancer_id"] = config.Reference{
			TerraformName: "alicloud_alb_load_balancer",
		}
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
	})
	p.AddResourceConfigurator("alicloud_alb_rule", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "Rule"
		r.References["listener_id"] = config.Reference{
			TerraformName: "alicloud_alb_listener",
		}
	})
	p.AddResourceConfigurator("alicloud_alb_listener_additional_certificate_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ALB)
		r.Kind = "ListenerAdditionalCertificateAttachment"
		// ALB expects the composed "<casCertId>-<region>" form here, not the
		// bare certificate id that the Certificate resource reports.
		r.References["certificate_id"] = config.Reference{
			TerraformName: "alicloud_ssl_certificates_service_certificate",
			Extractor:     common.PathAlbCertificateIdExtractor,
		}
		r.References["listener_id"] = config.Reference{
			TerraformName: "alicloud_alb_listener",
		}
	})
}
