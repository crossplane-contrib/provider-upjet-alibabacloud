package ecs

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_auto_provisioning_group", func(r *config.Resource) {
		r.Kind = "AutoProvisioningGroup"
		r.ShortGroup = string(common.ECS)
		r.References["launch_template_id"] = config.Reference{
			TerraformName: "alicloud_ecs_launch_template",
		}
	})
	p.AddResourceConfigurator("alicloud_ecs", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
	})
	p.AddResourceConfigurator("alicloud_ecs_auto_snapshot_policy", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		delete(r.TerraformResource.Schema, "name")
	})
	p.AddResourceConfigurator("alicloud_ecs_auto_snapshot_policy_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["auto_snapshot_policy_id"] = config.Reference{
			TerraformName: "alicloud_ecs_auto_snapshot_policy",
		}
		r.References["disk_id"] = config.Reference{
			TerraformName: "alicloud_ecs_disk",
		}
		delete(r.TerraformResource.Schema, "name")
	})
	p.AddResourceConfigurator("alicloud_ecs_deployment_set", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		delete(r.TerraformResource.Schema, "domain")
		delete(r.TerraformResource.Schema, "granularity")
	})
	p.AddResourceConfigurator("alicloud_ecs_disk", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_instance",
		}
		delete(r.TerraformResource.Schema, "availability_zone")
		delete(r.TerraformResource.Schema, "name")
	})
	p.AddResourceConfigurator("alicloud_ecs_disk_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_instance",
		}
		r.References["disk_id"] = config.Reference{
			TerraformName: "alicloud_ecs_disk",
		}
	})
	p.AddResourceConfigurator("alicloud_ecs_image_pipeline", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
	})
	p.AddResourceConfigurator("alicloud_ecs_image_pipeline_execution", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["image_pipeline_id"] = config.Reference{
			TerraformName: "alicloud_ecs_image_pipeline",
		}
	})
	p.AddResourceConfigurator("alicloud_ecs_instance_set", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["launch_template_id"] = config.Reference{
			TerraformName: "alicloud_ecs_launch_template",
		}
		r.References["deployment_set_id"] = config.Reference{
			TerraformName: "alicloud_ecs_deployment_set",
		}

	})
	p.AddResourceConfigurator("alicloud_ecs_invocation", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["command_id"] = config.Reference{
			TerraformName: "alicloud_ecs_command",
		}
	})
	p.AddResourceConfigurator("alicloud_ecs_key_pair", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		delete(r.TerraformResource.Schema, "key_name")
	})
	p.AddResourceConfigurator("alicloud_ecs_key_pair_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["key_pair_name"] = config.Reference{
			TerraformName: "alicloud_ecs_key_pair",
		}
		r.References["instance_ids"] = config.Reference{
			TerraformName:     "alicloud_instance",
			RefFieldName:      "InstanceRefs",
			SelectorFieldName: "InstanceSelector",
		}
		delete(r.TerraformResource.Schema, "key_name")
	})
	p.AddResourceConfigurator("alicloud_ecs_launch_template", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["deployment_set_id"] = config.Reference{
			TerraformName: "alicloud_ecs_deployment_set",
		}
		r.References["key_pair_name"] = config.Reference{
			TerraformName: "alicloud_ecs_key_pair",
		}
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["security_group_id"] = config.Reference{
			TerraformName: "alicloud_security_group",
		}
		r.References["security_group_ids"] = config.Reference{
			TerraformName:     "alicloud_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["networkInterfaces.security_group_id"] = config.Reference{
			TerraformName: "alicloud_security_group",
		}
		r.References["networkInterfaces.vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "system_disk_category")
		delete(r.TerraformResource.Schema, "system_disk_description")
		delete(r.TerraformResource.Schema, "system_disk_name")
		delete(r.TerraformResource.Schema, "system_disk_size")
		delete(r.TerraformResource.Schema, "userdata")
	})
	p.AddResourceConfigurator("alicloud_ecs_network_interface", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["security_group_ids"] = config.Reference{
			TerraformName:     "alicloud_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "security_groups")
		delete(r.TerraformResource.Schema, "private_ips")
		delete(r.TerraformResource.Schema, "private_ip")
		delete(r.TerraformResource.Schema, "private_ips_count")
	})
	p.AddResourceConfigurator("alicloud_ecs_network_interface_attachment", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["network_interface_id"] = config.Reference{
			TerraformName: "alicloud_ecs_network_interface",
		}
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_instance",
		}
	})
	p.AddResourceConfigurator("alicloud_ecs_network_interface_permission", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["network_interface_id"] = config.Reference{
			TerraformName: "alicloud_ecs_network_interface",
		}
	})
	p.AddResourceConfigurator("alicloud_ecs_snapshot", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["disk_id"] = config.Reference{
			TerraformName: "alicloud_ecs_disk",
		}
		delete(r.TerraformResource.Schema, "instant_access")
		delete(r.TerraformResource.Schema, "instant_access_retention_days")
		delete(r.TerraformResource.Schema, "name")
	})
	p.AddResourceConfigurator("alicloud_ecs_snapshot_group", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_instance",
		}
	})
	p.AddResourceConfigurator("alicloud_image", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.Kind = "Image"
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_instance",
		}
		r.References["snapshot_id"] = config.Reference{
			TerraformName: "alicloud_ecs_snapshot",
		}
		delete(r.TerraformResource.Schema, "name")
	})
	p.AddResourceConfigurator("alicloud_image_copy", func(r *config.Resource) {
		r.Kind = "ImageCopy"
		r.ShortGroup = string(common.ECS)
		r.References["source_image_id"] = config.Reference{
			TerraformName: "alicloud_image",
		}
		delete(r.TerraformResource.Schema, "name")
	})
	p.AddResourceConfigurator("alicloud_image_export", func(r *config.Resource) {
		r.Kind = "ImageExport"
		r.ShortGroup = string(common.ECS)
		r.References["image_id"] = config.Reference{
			TerraformName: "alicloud_image",
		}
	})
	p.AddResourceConfigurator("alicloud_image_import", func(r *config.Resource) {
		r.Kind = "ImageImport"
		r.ShortGroup = string(common.ECS)
	})
	p.AddResourceConfigurator("alicloud_image_share_permission", func(r *config.Resource) {
		r.Kind = "ImageSharePermission"
		r.ShortGroup = string(common.ECS)
		r.References["image_id"] = config.Reference{
			TerraformName: "alicloud_image",
		}
	})
	p.AddResourceConfigurator("alicloud_instance", func(r *config.Resource) {
		r.ShortGroup = string(common.ECS)
		r.Kind = "Instance"
		r.References["security_groups"] = config.Reference{
			TerraformName:     "alicloud_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["key_name"] = config.Reference{
			TerraformName: "alicloud_ecs_key_pair",
		}
		r.References["launch_template_id"] = config.Reference{
			TerraformName: "alicloud_ecs_launch_template",
		}
		r.References["image_id"] = config.Reference{
			TerraformName: "alicloud_image",
		}
		r.References["networkInterfaces.network_interface_id"] = config.Reference{
			TerraformName: "alicloud_ecs_network_interface",
		}
		r.References["networkInterfaces.security_group_ids"] = config.Reference{
			TerraformName:     "alicloud_security_group",
			RefFieldName:      "SecurityGroupRefs",
			SelectorFieldName: "SecurityGroupSelector",
		}
		r.References["networkInterfaces.vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		delete(r.TerraformResource.Schema, "allocate_public_ip")
		delete(r.TerraformResource.Schema, "internet_max_bandwidth_in")
		delete(r.TerraformResource.Schema, "io_optimized")
		delete(r.TerraformResource.Schema, "subnet_id")
	})
	p.AddResourceConfigurator("alicloud_reserved_instance", func(r *config.Resource) {
		r.Kind = "ReservedInstance"
		r.ShortGroup = string(common.ECS)
		delete(r.TerraformResource.Schema, "name")
	})
	p.AddResourceConfigurator("alicloud_security_group", func(r *config.Resource) {
		r.Kind = "SecurityGroup"
		r.ShortGroup = string(common.ECS)
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "inner_access")
	})
	p.AddResourceConfigurator("alicloud_security_group_rule", func(r *config.Resource) {
		r.Kind = "SecurityGroupRule"
		r.ShortGroup = string(common.ECS)
		r.References["security_group_id"] = config.Reference{
			TerraformName: "alicloud_security_group",
		}
		r.References["source_security_group_id"] = config.Reference{
			TerraformName: "alicloud_security_group",
		}
	})
}
