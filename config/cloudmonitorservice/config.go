package cloudmonitorservice

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_cms_alarm_contact_group", func(r *config.Resource) {
		r.ShortGroup = string(common.CLOUDMONITORSERVICE)
		r.Kind = "AlarmContactGroup"
	})
}
