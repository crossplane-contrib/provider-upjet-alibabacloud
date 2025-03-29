package ecs

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_auto_provisioning_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "AutoProvisioningGroup"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_disk", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_disk_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "DiskAttachment"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_image", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_image_copy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "ImageCopy"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_image_export", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "ImageExport"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_image_import", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "ImageImport"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_image_share_permission", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "ImageSharePermission"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_instance", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_key_pair", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "KeyPair"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_key_pair_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "KeyPairAttachment"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_launch_template", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "LaunchTemplate"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_network_interface", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "NetworkInterface"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_network_interface_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "NetworkInterfaceAttachment"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_ram_role_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "RamRoleAttachment"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_reserved_instance", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "ReservedInstance"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_security_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "SecurityGroup"
		r.ShortGroup = "ecs"
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
		// this resource, which would be "ecs"
		r.Kind = "SecurityGroupRule"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_snapshot", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.ShortGroup = "ecs"
	})
	p.AddResourceConfigurator("alicloud_snapshot_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ecs"
		r.Kind = "SnapshotPolicy"
		r.ShortGroup = "ecs"
	})
}
