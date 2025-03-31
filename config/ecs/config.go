package ecs

import "github.com/crossplane/upjet/pkg/config"

const Ecs = "ecs"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_auto_provisioning_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "AutoProvisioningGroup"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_disk", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_disk_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "DiskAttachment"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_image", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_image_copy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "ImageCopy"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_image_export", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "ImageExport"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_image_import", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "ImageImport"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_image_share_permission", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "ImageSharePermission"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_instance", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_key_pair", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "KeyPair"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_key_pair_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "KeyPairAttachment"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_launch_template", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "LaunchTemplate"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_network_interface", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "NetworkInterface"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_network_interface_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "NetworkInterfaceAttachment"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_ram_role_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "RamRoleAttachment"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_reserved_instance", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "ReservedInstance"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_security_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "SecurityGroup"
		r.ShortGroup = Ecs
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"name",
				"inner_access",
			},
		}
	})
	p.AddResourceConfigurator("alicloud_security_group_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "SecurityGroupRule"
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_snapshot", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.ShortGroup = Ecs
	})
	p.AddResourceConfigurator("alicloud_snapshot_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be Ecs
		r.Kind = "SnapshotPolicy"
		r.ShortGroup = Ecs
	})
}
