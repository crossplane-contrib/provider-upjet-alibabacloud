package messageservice

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_message_service_endpoint", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "messageservice"
		r.ShortGroup = string(common.MessageService)
		r.Kind = "Endpoint"
	})
	p.AddResourceConfigurator("alicloud_message_service_endpoint_acl", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "messageservice"
		r.ShortGroup = string(common.MessageService)
		r.Kind = "EndpointAcl"
		r.References["endpoint_type"] = config.Reference{
			TerraformName: "alicloud_message_service_endpoint",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_message_service_queue", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "messageservice"
		r.ShortGroup = string(common.MessageService)
		r.Kind = "Queue"
	})
	p.AddResourceConfigurator("alicloud_message_service_subscription", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "messageservice"
		r.ShortGroup = string(common.MessageService)
		r.Kind = "Subscription"
		r.References["topic_name"] = config.Reference{
			TerraformName: "alicloud_message_service_topic",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_message_service_topic", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "messageservice"
		r.ShortGroup = string(common.MessageService)
		r.Kind = "Topic"
		delete(r.TerraformResource.Schema, "logging_enabled")
	})
}
