package quotas

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_quotas_quota_alarm", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "quotas"
		r.ShortGroup = string(common.QUOTAS)
		r.Kind = "QuotaAlarm"
	})

	p.AddResourceConfigurator("alicloud_quotas_quota_application", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "quotas"
		r.ShortGroup = string(common.QUOTAS)
		r.Kind = "QuotaApplication"
	})

	p.AddResourceConfigurator("alicloud_quotas_template_applications", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "quotas"
		r.ShortGroup = string(common.QUOTAS)
		r.Kind = "TemplateApplications"
	})

	p.AddResourceConfigurator("alicloud_quotas_template_quota", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "quotas"
		r.ShortGroup = string(common.QUOTAS)
		r.Kind = "TemplateQuota"
	})

	p.AddResourceConfigurator("alicloud_quotas_template_service", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "quotas"
		r.ShortGroup = string(common.QUOTAS)
		r.Kind = "TemplateService"
	})
}
