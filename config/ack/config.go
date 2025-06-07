package ack

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_cs_autoscaling_config", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = string(common.ACK)

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
			TerraformName: "alicloud_vswitch",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_edge_kubernetes", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = string(common.ACK)
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

				"new_nat_gateway",

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
			RefFieldName:      "WorkerVswitchIDsRefs",
			SelectorFieldName: "WorkerVswitchIDsSelector",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_kubernetes", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = string(common.ACK)
		r.Kind = "Kubernetes"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				// Deprecated
				"availability_zone",
				"cpu_policy",
				"exclude_autoscaler_nodes",
				"kube_config",
				"node_port_range",
				"taints",
				"user_data",
				"worker_auto_renew",
				"worker_auto_renew_period",
				"worker_data_disks",
				"worker_disk_category",
				"worker_disk_performance_level",
				"worker_disk_size",
				"worker_disk_snapshot_policy_id",
				"worker_instance_types",
				"worker_instance_charge_type",
				"worker_number",
				"worker_period",
				"worker_period_unit",
				"worker_vswitch_ids",
			},
		}

		r.References["vpc_id"] = config.Reference{
			TerraformName: "alicloud_vpc",
		}
		r.References["master_vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["pod_vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
		r.References["bind_vpcs.pod_vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_kubernetes_addon", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = string(common.ACK)
		r.Kind = "KubernetesAddon"

		r.References["cluster_id"] = config.Reference{
			TerraformName:     "alicloud_cs_managed_kubernetes",
			RefFieldName:      "ClusterIDRefs",
			SelectorFieldName: "ClusterIDSelector",
		}
		r.References["worker_vswitch_ids"] = config.Reference{
			TerraformName:     "alicloud_vswitch",
			RefFieldName:      "WorkerVswitchIDsRefs",
			SelectorFieldName: "WorkerVswitchIDsSelector",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_kubernetes_node_pool", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = string(common.ACK)
		r.Kind = "KubernetesNodePool"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				// Deprecated
				"cis_enabled",
				"platform",
				"node_count",
				"rollout_policy",
				"name",
			},
		}

		r.References["cluster_id"] = config.Reference{
			TerraformName:     "alicloud_cs_managed_kubernetes",
			RefFieldName:      "ClusterIDRefs",
			SelectorFieldName: "ClusterIDSelector",
		}
		r.References["key_name"] = config.Reference{
			TerraformName: "alicloud_ecs_key_pair",
		}
		r.References["vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}

		delete(r.TerraformResource.Schema, "security_group_id")
	})

	p.AddResourceConfigurator("alicloud_cs_kubernetes_permissions", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = string(common.ACK)
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
		r.ShortGroup = string(common.ACK)
		r.Kind = "ManagedKubernetes"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				// name_prefix - (Optional) The kubernetes cluster name's prefix.
				// It is conflict with name. If it is specified, terraform will
				// use it to build the only cluster name.
				// Default to "Terraform-Creation".
				"name_prefix",

				"new_nat_gateway",

				// Deprecated
				"runtime",
				"enable_ssh",
				"rds_instances",
				"exclude_autoscaler_nodes",
				"worker_number",
				"worker_instance_types",
				"password",
				"key_name",
				"kms_encrypted_password",
				"kms_encryption_context",
				"worker_instance_charge_type",
				"worker_period",
				"worker_period_unit",
				"worker_auto_renew",
				"worker_auto_renew_period",
				"worker_disk_category",
				"worker_disk_size",
				"worker_data_disks",
				"worker_vswitch_ids",
				"node_name_mode",
				"node_port_range",
				"os_type",
				"platform",
				"image_id",
				"cpu_policy",
				"user_data",
				"taints",
				"worker_disk_performance_level",
				"worker_disk_snapshot_policy_id",
				"install_cloud_monitor",
				"kube_config",
				"availability_zone",
			},
		}

		r.References["vswitch_ids"] = config.Reference{
			TerraformName: "alicloud_vswitch",
		}
	})

	p.AddResourceConfigurator("alicloud_cs_serverless_kubernetes", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "ack"
		r.ShortGroup = string(common.ACK)
		r.Kind = "ServerlessKubernetes"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				// Ensure deletion is not blocked
				"new_nat_gateway",

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
			TerraformName: "alicloud_vswitch",
		}
	})
}
