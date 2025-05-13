package cs

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_cs_autoscaling_config", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cs"
		r.ShortGroup = "cs"
		r.Group = "ack.alibabacloud.crossplane.io"
		r.Kind = "AutoscalingConfig"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "VswitchRefs",
			SelectorFieldName: "VswitchSelector",
		}
		r.References["bind_vpcs.vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_edge_kubernetes", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cs"
		r.ShortGroup = "cs"
		r.Kind = "EdgeKubernetes"

		r.References["vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "VswitchRefs",
			SelectorFieldName: "VswitchSelector",
		}
		r.References["bind_vpcs.vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}

		delete(r.TerraformResource.Schema, "kube_config")
		delete(r.TerraformResource.Schema, "client_cert")
		delete(r.TerraformResource.Schema, "client_key")
		delete(r.TerraformResource.Schema, "cluster_ca_cert")
		delete(r.TerraformResource.Schema, "log_config")
		delete(r.TerraformResource.Schema, "force_update")
	})

	p.AddResourceConfigurator("alicloud_cs_kubernetes", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cs"
		r.ShortGroup = "cs"
		r.Kind = "Kubernetes"

		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["bind_vpcs.vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["master_vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "MasterVswitchRefs",
			SelectorFieldName: "MasterVswitchSelector",
		}
		r.References["bind_vpcs.master_vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["worker_vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "WorkerVswitchRefs",
			SelectorFieldName: "WorkerVswitchSelector",
		}
		r.References["bind_vpcs.worker_vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["pod_vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "PodVswitchRefs",
			SelectorFieldName: "PodVswitchSelector",
		}
		r.References["bind_vpcs.pod_vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}

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
		// this resource, which would be "cs"
		r.ShortGroup = "cs"
		r.Kind = "KubernetesAddon"

		r.References["cluster_id"] = config.Reference{
			TerraformName: "alicloud_cs_managed_kubernetes",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_kubernetes_node_pool", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cs"
		r.ShortGroup = "cs"
		r.Kind = "KubernetesNodePool"

		r.References["cluster_id"] = config.Reference{
			TerraformName: "alicloud_cs_managed_kubernetes",
		}
		r.References["key_name"] = config.Reference{
			TerraformName: "alicloud_key_pair",
		}
		r.References["vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "VswitchRefs",
			SelectorFieldName: "VswitchSelector",
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
		// this resource, which would be "cs"
		r.ShortGroup = "cs"
		r.Kind = "KubernetesPermissions"

		r.References["uid"] = config.Reference{
			TerraformName: "alicloud_cs_managed_kubernetes",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_managed_kubernetes", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "cs"
		r.ShortGroup = "cs"
		r.Kind = "ManagedKubernetes"

		r.References["vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "VswitchRefs",
			SelectorFieldName: "VswitchSelector",
		}
		r.References["bind_vpcs.vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}

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
		// this resource, which would be "cs"
		r.ShortGroup = "cs"
		r.Kind = "ServerlessKubernetes"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "VswitchRefs",
			SelectorFieldName: "VswitchSelector",
		}
		r.References["bind_vpcs.vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}

		delete(r.TerraformResource.Schema, "kube_config")
		delete(r.TerraformResource.Schema, "client_cert")
		delete(r.TerraformResource.Schema, "client_key")
		delete(r.TerraformResource.Schema, "cluster_ca_cert")
		delete(r.TerraformResource.Schema, "load_balancer_spec")
		delete(r.TerraformResource.Schema, "logging_type")
		delete(r.TerraformResource.Schema, "sls_project_name")
	})
}
