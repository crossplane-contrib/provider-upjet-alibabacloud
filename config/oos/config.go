package oos

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_oos_application", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_parameter", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_template", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_application_group", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_default_patch_baseline", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_execution", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
		// Terraform provider 1.293.0 added "tags" here as tagsSchemaForceNew.
		// AddExternalTagsField would otherwise write the crossplane-* tags into
		// spec, and on an Execution created before that field existed the plan
		// wants to add them, which is a replacement. Upjet renders
		// prevent_destroy for resources that are not being deleted, so the
		// apply fails on every reconcile instead. Users can still set tags
		// themselves; we just do not inject them.
		r.InitializerFns = nil
	})
	p.AddResourceConfigurator("alicloud_oos_patch_baseline", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_secret_parameter", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
	p.AddResourceConfigurator("alicloud_oos_service_setting", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
		// Upstream's DiffSuppressFunc on these fields checks
		// delivery_sls_enabled instead of delivery_oss_enabled, so they are
		// never sent unless SLS delivery is on too:
		// https://github.com/aliyun/terraform-provider-alicloud/issues/11154
		next := r.TerraformCustomDiff
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, cfg *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if next != nil {
				var err error
				if diff, err = next(diff, state, cfg); err != nil {
					return nil, err
				}
			}
			return restoreSuppressedStrings(diff, state, cfg, "delivery_oss_bucket_name", "delivery_oss_key_prefix"), nil
		}
	})
	p.AddResourceConfigurator("alicloud_oos_state_configuration", func(r *config.Resource) {
		r.ShortGroup = string(common.OOS)
	})
}

// restoreSuppressedStrings puts back the diff for each string field that is
// set in the configuration, differs from the prior state, and is missing from
// diff -- which, for these fields, only an upstream DiffSuppressFunc does.
func restoreSuppressedStrings(diff *terraform.InstanceDiff, state *terraform.InstanceState, cfg *terraform.ResourceConfig, fields ...string) *terraform.InstanceDiff {
	if cfg == nil {
		return diff
	}
	for _, f := range fields {
		raw, ok := cfg.Get(f)
		if !ok {
			continue
		}
		want, ok := raw.(string)
		if !ok || want == "" {
			continue
		}
		old := ""
		if state != nil {
			old = state.Attributes[f]
		}
		if old == want {
			continue
		}
		if diff == nil {
			diff = terraform.NewInstanceDiff()
		}
		if _, ok := diff.Attributes[f]; ok {
			continue
		}
		diff.Attributes[f] = &terraform.ResourceAttrDiff{Old: old, New: want}
	}
	return diff
}
