package ack

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_cs_autoscaling_config", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = "ack"

		r.Kind = "AutoscalingConfig"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["cluster_id"] = config.Reference{
			TerraformName:     "alicloud_cs_managed_kubernetes",
			RefFieldName:      "ClusterIDRefs",
			SelectorFieldName: "ClusterIDSelector",
		}
		r.References["vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "VswitchIDRefs",
			SelectorFieldName: "VswitchIDSelector",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_edge_kubernetes", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = "ack"
		r.Kind = "EdgeKubernetes"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				// name_prefix - (Optional) The kubernetes cluster name's prefix.
				// It is conflict with name. If it is specified, terraform will
				// using it to build the only cluster name.
				// Default to "Terraform-Creation".
				"name_prefix",

				// Conflicts with is_enterprise_security_group
				"security_group_id",

				// Test if this causes lifecycle prevent_destroy: true
				"is_enterprise_security_group",

				// Deprecated
				"kube_config",
				"client_cert",
				"client_key",
				"cluster_ca_cert",
				"log_config",
				"force_update",
			},
		}

		r.References["worker_vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "WorkerVswitchIDRefs",
			SelectorFieldName: "WorkerVswitchIDSelector",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_kubernetes", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = "ack"
		r.Kind = "Kubernetes"

		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		// TODO: Try without bind, preference to remove
		// r.References["bind_vpcs.vswitch_id"] = config.Reference{
		// TerraformName: "alicloud_vswitch",
		// }
		r.References["master_vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["pod_vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["bind_vpcs.pod_vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}

		// TODO: Use for deprecated fields ...
		//  r.LateInitializer = config.LateInitializer{
		//		IgnoredFields: []string{
		//			"name",
		//		},
		//}

		delete(r.TerraformResource.Schema, "availability_zone")
		delete(r.TerraformResource.Schema, "cpu_policy")
		delete(r.TerraformResource.Schema, "exclude_autoscaler_nodes")
		delete(r.TerraformResource.Schema, "kube_config")
		delete(r.TerraformResource.Schema, "node_port_range")
		delete(r.TerraformResource.Schema, "taints")
		delete(r.TerraformResource.Schema, "user_data")
		delete(r.TerraformResource.Schema, "worker_auto_renew")
		delete(r.TerraformResource.Schema, "worker_auto_renew_period")
		delete(r.TerraformResource.Schema, "worker_data_disks")
		delete(r.TerraformResource.Schema, "worker_disk_category")
		delete(r.TerraformResource.Schema, "worker_disk_performance_level")
		delete(r.TerraformResource.Schema, "worker_disk_size")
		delete(r.TerraformResource.Schema, "worker_disk_snapshot_policy_id")
		delete(r.TerraformResource.Schema, "worker_instance_types")
		delete(r.TerraformResource.Schema, "worker_instance_charge_type")
		delete(r.TerraformResource.Schema, "worker_number")
		delete(r.TerraformResource.Schema, "worker_period")
		delete(r.TerraformResource.Schema, "worker_period_unit")
		delete(r.TerraformResource.Schema, "worker_vswitch_ids")
	})

	p.AddResourceConfigurator("alicloud_cs_kubernetes_addon", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = "ack"
		r.Kind = "KubernetesAddon"

		r.References["cluster_id"] = config.Reference{
			TerraformName:     "alicloud_cs_managed_kubernetes",
			RefFieldName:      "ClusterIDRefs",
			SelectorFieldName: "ClusterIDSelector",
		}
		r.References["worker_vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "WorkerVswitchIDRefs",
			SelectorFieldName: "WorkerVswitchIDSelector",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_kubernetes_node_pool", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = "ack"
		r.Kind = "KubernetesNodePool"

		r.References["cluster_id"] = config.Reference{
			TerraformName:     "alicloud_cs_managed_kubernetes",
			RefFieldName:      "ClusterIDRefs",
			SelectorFieldName: "ClusterIDSelector",
		}
		r.References["key_name"] = config.Reference{
			TerraformName: "alicloud_ecs_key_pair",
		}
		r.References["vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "VswitchIDRefs",
			SelectorFieldName: "VswitchIDSelector",
		}

		delete(r.TerraformResource.Schema, "cis_enabled")
		delete(r.TerraformResource.Schema, "platform")
		delete(r.TerraformResource.Schema, "security_group_id")
		delete(r.TerraformResource.Schema, "node_count")
		delete(r.TerraformResource.Schema, "rollout_policy")
		delete(r.TerraformResource.Schema, "name")
	})

	p.AddResourceConfigurator("alicloud_cs_kubernetes_permissions", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = "ack"
		r.Kind = "KubernetesPermissions"

		r.References["permissions.cluster"] = config.Reference{
			TerraformName:     "alicloud_cs_managed_kubernetes",
			RefFieldName:      "ClusterIDRefs",
			SelectorFieldName: "ClusterIDSelector",
		}
		r.References["permissions.role_name"] = config.Reference{
			TerraformName:     "alicloud_ram_role",
			RefFieldName:      "RoleIDRefs",
			SelectorFieldName: "RoleIDSelector",
		}
		r.References["uid"] = config.Reference{
			TerraformName: "alicloud_ram_user",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_managed_kubernetes", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = "ack"
		r.Kind = "ManagedKubernetes"

		r.References["vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "VswitchIDRefs",
			SelectorFieldName: "VswitchIDSelector",
		}
		r.References["worker_vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "WorkerVswitchIDRefs",
			SelectorFieldName: "WorkerVswitchIDSelector",
		}

		// name_prefix - (Optional) The kubernetes cluster name's prefix.
		// It is conflict with name. If it is specified, terraform will
		// use it to build the only cluster name.
		// Default to "Terraform-Creation".
		delete(r.TerraformResource.Schema, "name_prefix")

		delete(r.TerraformResource.Schema, "runtime")
		delete(r.TerraformResource.Schema, "enable_ssh")
		delete(r.TerraformResource.Schema, "rds_instances")
		delete(r.TerraformResource.Schema, "exclude_autoscaler_nodes")
		delete(r.TerraformResource.Schema, "worker_number")
		delete(r.TerraformResource.Schema, "worker_instance_types")
		delete(r.TerraformResource.Schema, "password")
		delete(r.TerraformResource.Schema, "key_name")
		delete(r.TerraformResource.Schema, "kms_encrypted_password")
		delete(r.TerraformResource.Schema, "kms_encryption_context")
		delete(r.TerraformResource.Schema, "worker_instance_charge_type")
		delete(r.TerraformResource.Schema, "worker_period")
		delete(r.TerraformResource.Schema, "worker_period_unit")
		delete(r.TerraformResource.Schema, "worker_auto_renew")
		delete(r.TerraformResource.Schema, "worker_auto_renew_period")
		delete(r.TerraformResource.Schema, "worker_disk_category")
		delete(r.TerraformResource.Schema, "worker_disk_size")
		delete(r.TerraformResource.Schema, "worker_data_disks")
		delete(r.TerraformResource.Schema, "node_name_mode")
		delete(r.TerraformResource.Schema, "node_port_range")
		delete(r.TerraformResource.Schema, "os_type")
		delete(r.TerraformResource.Schema, "platform")
		delete(r.TerraformResource.Schema, "image_id")
		delete(r.TerraformResource.Schema, "cpu_policy")
		delete(r.TerraformResource.Schema, "user_data")
		delete(r.TerraformResource.Schema, "taints")
		delete(r.TerraformResource.Schema, "worker_disk_performance_level")
		delete(r.TerraformResource.Schema, "worker_disk_snapshot_policy_id")
		delete(r.TerraformResource.Schema, "install_cloud_monitor")
		delete(r.TerraformResource.Schema, "kube_config")
		delete(r.TerraformResource.Schema, "availability_zone")
	})

	p.AddResourceConfigurator("alicloud_cs_serverless_kubernetes", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = "ack"
		r.Kind = "ServerlessKubernetes"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				// Deprecated
				"kube_config",
				"client_cert",
				"client_key",
				"cluster_ca_cert",
				"load_balancer_spec",
				"logging_type",
				"sls_project_name",
			},
		}

		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "VswitchIDRefs",
			SelectorFieldName: "VswitchIDSelector",
		}
	})
}
