// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConnectionDrainConfigInitParameters struct {

	// Specifies whether to enable connection draining. Valid values:
	ConnectionDrainEnabled *bool `json:"connectionDrainEnabled,omitempty" tf:"connection_drain_enabled,omitempty"`

	// The timeout period of connection draining.
	ConnectionDrainTimeout *float64 `json:"connectionDrainTimeout,omitempty" tf:"connection_drain_timeout,omitempty"`
}

type ConnectionDrainConfigObservation struct {

	// Specifies whether to enable connection draining. Valid values:
	ConnectionDrainEnabled *bool `json:"connectionDrainEnabled,omitempty" tf:"connection_drain_enabled,omitempty"`

	// The timeout period of connection draining.
	ConnectionDrainTimeout *float64 `json:"connectionDrainTimeout,omitempty" tf:"connection_drain_timeout,omitempty"`
}

type ConnectionDrainConfigParameters struct {

	// Specifies whether to enable connection draining. Valid values:
	// +kubebuilder:validation:Optional
	ConnectionDrainEnabled *bool `json:"connectionDrainEnabled,omitempty" tf:"connection_drain_enabled,omitempty"`

	// The timeout period of connection draining.
	// +kubebuilder:validation:Optional
	ConnectionDrainTimeout *float64 `json:"connectionDrainTimeout,omitempty" tf:"connection_drain_timeout,omitempty"`
}

type HealthCheckConfigInitParameters struct {

	// The status code for a successful health check
	HealthCheckCodes []*string `json:"healthCheckCodes,omitempty" tf:"health_check_codes,omitempty"`

	// The backend port that is used for health checks.
	HealthCheckConnectPort *float64 `json:"healthCheckConnectPort,omitempty" tf:"health_check_connect_port,omitempty"`

	// Specifies whether to enable the health check feature. Valid values:
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty" tf:"health_check_enabled,omitempty"`

	// The HTTP version that is used for health checks. Valid values:
	HealthCheckHTTPVersion *string `json:"healthCheckHttpVersion,omitempty" tf:"health_check_http_version,omitempty"`

	// The domain name that is used for health checks.
	HealthCheckHost *string `json:"healthCheckHost,omitempty" tf:"health_check_host,omitempty"`

	// The interval at which health checks are performed. Unit: seconds.
	HealthCheckInterval *float64 `json:"healthCheckInterval,omitempty" tf:"health_check_interval,omitempty"`

	// The HTTP method that is used for health checks. Valid values:
	HealthCheckMethod *string `json:"healthCheckMethod,omitempty" tf:"health_check_method,omitempty"`

	// The URL that is used for health checks.
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// The protocol that is used for health checks. Valid values:
	HealthCheckProtocol *string `json:"healthCheckProtocol,omitempty" tf:"health_check_protocol,omitempty"`

	// The timeout period of a health check response. If a backend ECS instance does not respond within the specified timeout period, the ECS instance fails the health check. Unit: seconds.
	HealthCheckTimeout *float64 `json:"healthCheckTimeout,omitempty" tf:"health_check_timeout,omitempty"`

	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health check status of the backend server changes from fail to success.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health check status of the backend server changes from success to fail.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthCheckConfigObservation struct {

	// The status code for a successful health check
	HealthCheckCodes []*string `json:"healthCheckCodes,omitempty" tf:"health_check_codes,omitempty"`

	// The backend port that is used for health checks.
	HealthCheckConnectPort *float64 `json:"healthCheckConnectPort,omitempty" tf:"health_check_connect_port,omitempty"`

	// Specifies whether to enable the health check feature. Valid values:
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty" tf:"health_check_enabled,omitempty"`

	// The HTTP version that is used for health checks. Valid values:
	HealthCheckHTTPVersion *string `json:"healthCheckHttpVersion,omitempty" tf:"health_check_http_version,omitempty"`

	// The domain name that is used for health checks.
	HealthCheckHost *string `json:"healthCheckHost,omitempty" tf:"health_check_host,omitempty"`

	// The interval at which health checks are performed. Unit: seconds.
	HealthCheckInterval *float64 `json:"healthCheckInterval,omitempty" tf:"health_check_interval,omitempty"`

	// The HTTP method that is used for health checks. Valid values:
	HealthCheckMethod *string `json:"healthCheckMethod,omitempty" tf:"health_check_method,omitempty"`

	// The URL that is used for health checks.
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// The protocol that is used for health checks. Valid values:
	HealthCheckProtocol *string `json:"healthCheckProtocol,omitempty" tf:"health_check_protocol,omitempty"`

	// The timeout period of a health check response. If a backend ECS instance does not respond within the specified timeout period, the ECS instance fails the health check. Unit: seconds.
	HealthCheckTimeout *float64 `json:"healthCheckTimeout,omitempty" tf:"health_check_timeout,omitempty"`

	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health check status of the backend server changes from fail to success.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health check status of the backend server changes from success to fail.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthCheckConfigParameters struct {

	// The status code for a successful health check
	// +kubebuilder:validation:Optional
	HealthCheckCodes []*string `json:"healthCheckCodes,omitempty" tf:"health_check_codes,omitempty"`

	// The backend port that is used for health checks.
	// +kubebuilder:validation:Optional
	HealthCheckConnectPort *float64 `json:"healthCheckConnectPort,omitempty" tf:"health_check_connect_port,omitempty"`

	// Specifies whether to enable the health check feature. Valid values:
	// +kubebuilder:validation:Optional
	HealthCheckEnabled *bool `json:"healthCheckEnabled" tf:"health_check_enabled,omitempty"`

	// The HTTP version that is used for health checks. Valid values:
	// +kubebuilder:validation:Optional
	HealthCheckHTTPVersion *string `json:"healthCheckHttpVersion,omitempty" tf:"health_check_http_version,omitempty"`

	// The domain name that is used for health checks.
	// +kubebuilder:validation:Optional
	HealthCheckHost *string `json:"healthCheckHost,omitempty" tf:"health_check_host,omitempty"`

	// The interval at which health checks are performed. Unit: seconds.
	// +kubebuilder:validation:Optional
	HealthCheckInterval *float64 `json:"healthCheckInterval,omitempty" tf:"health_check_interval,omitempty"`

	// The HTTP method that is used for health checks. Valid values:
	// +kubebuilder:validation:Optional
	HealthCheckMethod *string `json:"healthCheckMethod,omitempty" tf:"health_check_method,omitempty"`

	// The URL that is used for health checks.
	// +kubebuilder:validation:Optional
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// The protocol that is used for health checks. Valid values:
	// +kubebuilder:validation:Optional
	HealthCheckProtocol *string `json:"healthCheckProtocol,omitempty" tf:"health_check_protocol,omitempty"`

	// The timeout period of a health check response. If a backend ECS instance does not respond within the specified timeout period, the ECS instance fails the health check. Unit: seconds.
	// +kubebuilder:validation:Optional
	HealthCheckTimeout *float64 `json:"healthCheckTimeout,omitempty" tf:"health_check_timeout,omitempty"`

	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health check status of the backend server changes from fail to success.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health check status of the backend server changes from success to fail.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type ServerGroupInitParameters struct {

	// Elegant interrupt configuration. See connection_drain_config below.
	ConnectionDrainConfig []ConnectionDrainConfigInitParameters `json:"connectionDrainConfig,omitempty" tf:"connection_drain_config,omitempty"`

	// Indicates whether cross-zone load balancing is enabled for the server group. Valid values:
	CrossZoneEnabled *bool `json:"crossZoneEnabled,omitempty" tf:"cross_zone_enabled,omitempty"`

	// Whether to PreCheck only this request. Value:
	// true: Send a check request,
	// false (default): Send a normal request.
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// The configuration of health checks See health_check_config below.
	HealthCheckConfig []HealthCheckConfigInitParameters `json:"healthCheckConfig,omitempty" tf:"health_check_config,omitempty"`

	// The ID of the resource group to which you want to transfer the cloud resource.
	HealthCheckTemplateID *string `json:"healthCheckTemplateId,omitempty" tf:"health_check_template_id,omitempty"`

	// Enable Ipv6
	IPv6Enabled *bool `json:"ipv6Enabled,omitempty" tf:"ipv6_enabled,omitempty"`

	// The backend protocol. Valid values:
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Elegant interrupt configuration.
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// The scheduling algorithm. Valid values:
	Scheduler *string `json:"scheduler,omitempty" tf:"scheduler,omitempty"`

	// The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	ServerGroupName *string `json:"serverGroupName,omitempty" tf:"server_group_name,omitempty"`

	// The type of server group. Valid values:
	ServerGroupType *string `json:"serverGroupType,omitempty" tf:"server_group_type,omitempty"`

	// List of servers. See servers below.
	Servers []ServersInitParameters `json:"servers,omitempty" tf:"servers,omitempty"`

	// Only applicable to the ALB Ingress scenario, indicating the K8s Service name corresponding to the server group.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Slow start configuration. See slow_start_config below.
	SlowStartConfig []SlowStartConfigInitParameters `json:"slowStartConfig,omitempty" tf:"slow_start_config,omitempty"`

	// The configuration of health checks See sticky_session_config below.
	StickySessionConfig []StickySessionConfigInitParameters `json:"stickySessionConfig,omitempty" tf:"sticky_session_config,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Url consistency hash parameter configuration See uch_config below.
	UchConfig []UchConfigInitParameters `json:"uchConfig,omitempty" tf:"uch_config,omitempty"`

	// Specifies whether to enable persistent TCP connections.
	UpstreamKeepaliveEnabled *bool `json:"upstreamKeepaliveEnabled,omitempty" tf:"upstream_keepalive_enabled,omitempty"`

	// The ID of the virtual private cloud (VPC). You can add only servers that are deployed in the specified VPC to the server group.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/vpc/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type ServerGroupObservation struct {

	// Elegant interrupt configuration. See connection_drain_config below.
	ConnectionDrainConfig []ConnectionDrainConfigObservation `json:"connectionDrainConfig,omitempty" tf:"connection_drain_config,omitempty"`

	// The creation time of the resource
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Indicates whether cross-zone load balancing is enabled for the server group. Valid values:
	CrossZoneEnabled *bool `json:"crossZoneEnabled,omitempty" tf:"cross_zone_enabled,omitempty"`

	// Whether to PreCheck only this request. Value:
	// true: Send a check request,
	// false (default): Send a normal request.
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// The configuration of health checks See health_check_config below.
	HealthCheckConfig []HealthCheckConfigObservation `json:"healthCheckConfig,omitempty" tf:"health_check_config,omitempty"`

	// The ID of the resource group to which you want to transfer the cloud resource.
	HealthCheckTemplateID *string `json:"healthCheckTemplateId,omitempty" tf:"health_check_template_id,omitempty"`

	// The ID of the resource supplied above.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Enable Ipv6
	IPv6Enabled *bool `json:"ipv6Enabled,omitempty" tf:"ipv6_enabled,omitempty"`

	// The backend protocol. Valid values:
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Elegant interrupt configuration.
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// The scheduling algorithm. Valid values:
	Scheduler *string `json:"scheduler,omitempty" tf:"scheduler,omitempty"`

	// The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	ServerGroupName *string `json:"serverGroupName,omitempty" tf:"server_group_name,omitempty"`

	// The type of server group. Valid values:
	ServerGroupType *string `json:"serverGroupType,omitempty" tf:"server_group_type,omitempty"`

	// List of servers. See servers below.
	Servers []ServersObservation `json:"servers,omitempty" tf:"servers,omitempty"`

	// Only applicable to the ALB Ingress scenario, indicating the K8s Service name corresponding to the server group.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Slow start configuration. See slow_start_config below.
	SlowStartConfig []SlowStartConfigObservation `json:"slowStartConfig,omitempty" tf:"slow_start_config,omitempty"`

	// The addition status of the backend server. Value:
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The configuration of health checks See sticky_session_config below.
	StickySessionConfig []StickySessionConfigObservation `json:"stickySessionConfig,omitempty" tf:"sticky_session_config,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Url consistency hash parameter configuration See uch_config below.
	UchConfig []UchConfigObservation `json:"uchConfig,omitempty" tf:"uch_config,omitempty"`

	// Specifies whether to enable persistent TCP connections.
	UpstreamKeepaliveEnabled *bool `json:"upstreamKeepaliveEnabled,omitempty" tf:"upstream_keepalive_enabled,omitempty"`

	// The ID of the virtual private cloud (VPC). You can add only servers that are deployed in the specified VPC to the server group.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type ServerGroupParameters struct {

	// Elegant interrupt configuration. See connection_drain_config below.
	// +kubebuilder:validation:Optional
	ConnectionDrainConfig []ConnectionDrainConfigParameters `json:"connectionDrainConfig,omitempty" tf:"connection_drain_config,omitempty"`

	// Indicates whether cross-zone load balancing is enabled for the server group. Valid values:
	// +kubebuilder:validation:Optional
	CrossZoneEnabled *bool `json:"crossZoneEnabled,omitempty" tf:"cross_zone_enabled,omitempty"`

	// Whether to PreCheck only this request. Value:
	// true: Send a check request,
	// false (default): Send a normal request.
	// +kubebuilder:validation:Optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// The configuration of health checks See health_check_config below.
	// +kubebuilder:validation:Optional
	HealthCheckConfig []HealthCheckConfigParameters `json:"healthCheckConfig,omitempty" tf:"health_check_config,omitempty"`

	// The ID of the resource group to which you want to transfer the cloud resource.
	// +kubebuilder:validation:Optional
	HealthCheckTemplateID *string `json:"healthCheckTemplateId,omitempty" tf:"health_check_template_id,omitempty"`

	// Enable Ipv6
	// +kubebuilder:validation:Optional
	IPv6Enabled *bool `json:"ipv6Enabled,omitempty" tf:"ipv6_enabled,omitempty"`

	// The backend protocol. Valid values:
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// Elegant interrupt configuration.
	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// The scheduling algorithm. Valid values:
	// +kubebuilder:validation:Optional
	Scheduler *string `json:"scheduler,omitempty" tf:"scheduler,omitempty"`

	// The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	// +kubebuilder:validation:Optional
	ServerGroupName *string `json:"serverGroupName,omitempty" tf:"server_group_name,omitempty"`

	// The type of server group. Valid values:
	// +kubebuilder:validation:Optional
	ServerGroupType *string `json:"serverGroupType,omitempty" tf:"server_group_type,omitempty"`

	// List of servers. See servers below.
	// +kubebuilder:validation:Optional
	Servers []ServersParameters `json:"servers,omitempty" tf:"servers,omitempty"`

	// Only applicable to the ALB Ingress scenario, indicating the K8s Service name corresponding to the server group.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Slow start configuration. See slow_start_config below.
	// +kubebuilder:validation:Optional
	SlowStartConfig []SlowStartConfigParameters `json:"slowStartConfig,omitempty" tf:"slow_start_config,omitempty"`

	// The configuration of health checks See sticky_session_config below.
	// +kubebuilder:validation:Optional
	StickySessionConfig []StickySessionConfigParameters `json:"stickySessionConfig,omitempty" tf:"sticky_session_config,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Url consistency hash parameter configuration See uch_config below.
	// +kubebuilder:validation:Optional
	UchConfig []UchConfigParameters `json:"uchConfig,omitempty" tf:"uch_config,omitempty"`

	// Specifies whether to enable persistent TCP connections.
	// +kubebuilder:validation:Optional
	UpstreamKeepaliveEnabled *bool `json:"upstreamKeepaliveEnabled,omitempty" tf:"upstream_keepalive_enabled,omitempty"`

	// The ID of the virtual private cloud (VPC). You can add only servers that are deployed in the specified VPC to the server group.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type ServersInitParameters struct {

	// The description of the backend server. The description must be 2 to 256 characters in length, and cannot start with http:// or https://.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The port that is used by the backend server. Valid values: 1 to 65535. You can specify at most 200 servers in each call.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies whether to enable the remote IP feature. You can specify at most 200 servers in each call. Default values:
	RemoteIPEnabled *bool `json:"remoteIpEnabled,omitempty" tf:"remote_ip_enabled,omitempty"`

	// The ID of the backend server. You can specify at most 200 servers in each call.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ecs/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Instance in ecs to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ecs to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// The IP address of the backend server. You can specify at most 200 servers in each call.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ecs/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("private_ip",false)
	ServerIP *string `json:"serverIp,omitempty" tf:"server_ip,omitempty"`

	// Reference to a Instance in ecs to populate serverIp.
	// +kubebuilder:validation:Optional
	ServerIPRef *v1.Reference `json:"serverIpRef,omitempty" tf:"-"`

	// Selector for a Instance in ecs to populate serverIp.
	// +kubebuilder:validation:Optional
	ServerIPSelector *v1.Selector `json:"serverIpSelector,omitempty" tf:"-"`

	// The type of the backend server. You can specify at most 200 servers in each call. Default values:
	ServerType *string `json:"serverType,omitempty" tf:"server_type,omitempty"`

	// The weight of the backend server. Valid values: 0 to 100. Default value: 100. If the value is set to 0, no requests are forwarded to the server. You can specify at most 200 servers in each call.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type ServersObservation struct {

	// The description of the backend server. The description must be 2 to 256 characters in length, and cannot start with http:// or https://.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The port that is used by the backend server. Valid values: 1 to 65535. You can specify at most 200 servers in each call.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies whether to enable the remote IP feature. You can specify at most 200 servers in each call. Default values:
	RemoteIPEnabled *bool `json:"remoteIpEnabled,omitempty" tf:"remote_ip_enabled,omitempty"`

	// The ID of the server group.
	ServerGroupID *string `json:"serverGroupId,omitempty" tf:"server_group_id,omitempty"`

	// The ID of the backend server. You can specify at most 200 servers in each call.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// The IP address of the backend server. You can specify at most 200 servers in each call.
	ServerIP *string `json:"serverIp,omitempty" tf:"server_ip,omitempty"`

	// The type of the backend server. You can specify at most 200 servers in each call. Default values:
	ServerType *string `json:"serverType,omitempty" tf:"server_type,omitempty"`

	// The addition status of the backend server. Value:
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The weight of the backend server. Valid values: 0 to 100. Default value: 100. If the value is set to 0, no requests are forwarded to the server. You can specify at most 200 servers in each call.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type ServersParameters struct {

	// The description of the backend server. The description must be 2 to 256 characters in length, and cannot start with http:// or https://.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The port that is used by the backend server. Valid values: 1 to 65535. You can specify at most 200 servers in each call.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies whether to enable the remote IP feature. You can specify at most 200 servers in each call. Default values:
	// +kubebuilder:validation:Optional
	RemoteIPEnabled *bool `json:"remoteIpEnabled,omitempty" tf:"remote_ip_enabled,omitempty"`

	// The ID of the backend server. You can specify at most 200 servers in each call.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ecs/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Instance in ecs to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ecs to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// The IP address of the backend server. You can specify at most 200 servers in each call.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ecs/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("private_ip",false)
	// +kubebuilder:validation:Optional
	ServerIP *string `json:"serverIp,omitempty" tf:"server_ip,omitempty"`

	// Reference to a Instance in ecs to populate serverIp.
	// +kubebuilder:validation:Optional
	ServerIPRef *v1.Reference `json:"serverIpRef,omitempty" tf:"-"`

	// Selector for a Instance in ecs to populate serverIp.
	// +kubebuilder:validation:Optional
	ServerIPSelector *v1.Selector `json:"serverIpSelector,omitempty" tf:"-"`

	// The type of the backend server. You can specify at most 200 servers in each call. Default values:
	// +kubebuilder:validation:Optional
	ServerType *string `json:"serverType" tf:"server_type,omitempty"`

	// The weight of the backend server. Valid values: 0 to 100. Default value: 100. If the value is set to 0, no requests are forwarded to the server. You can specify at most 200 servers in each call.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type SlowStartConfigInitParameters struct {

	// The duration of a slow start.
	SlowStartDuration *float64 `json:"slowStartDuration,omitempty" tf:"slow_start_duration,omitempty"`

	// Indicates whether slow starts are enabled. Valid values:
	SlowStartEnabled *bool `json:"slowStartEnabled,omitempty" tf:"slow_start_enabled,omitempty"`
}

type SlowStartConfigObservation struct {

	// The duration of a slow start.
	SlowStartDuration *float64 `json:"slowStartDuration,omitempty" tf:"slow_start_duration,omitempty"`

	// Indicates whether slow starts are enabled. Valid values:
	SlowStartEnabled *bool `json:"slowStartEnabled,omitempty" tf:"slow_start_enabled,omitempty"`
}

type SlowStartConfigParameters struct {

	// The duration of a slow start.
	// +kubebuilder:validation:Optional
	SlowStartDuration *float64 `json:"slowStartDuration,omitempty" tf:"slow_start_duration,omitempty"`

	// Indicates whether slow starts are enabled. Valid values:
	// +kubebuilder:validation:Optional
	SlowStartEnabled *bool `json:"slowStartEnabled,omitempty" tf:"slow_start_enabled,omitempty"`
}

type StickySessionConfigInitParameters struct {

	// The cookie to be configured on the server.
	Cookie *string `json:"cookie,omitempty" tf:"cookie,omitempty"`

	// The maximum amount of time to wait before the session cookie expires. Unit: seconds.
	CookieTimeout *float64 `json:"cookieTimeout,omitempty" tf:"cookie_timeout,omitempty"`

	// Specifies whether to enable session persistence. Valid values:
	StickySessionEnabled *bool `json:"stickySessionEnabled,omitempty" tf:"sticky_session_enabled,omitempty"`

	// The method that is used to handle a cookie. Valid values:
	StickySessionType *string `json:"stickySessionType,omitempty" tf:"sticky_session_type,omitempty"`
}

type StickySessionConfigObservation struct {

	// The cookie to be configured on the server.
	Cookie *string `json:"cookie,omitempty" tf:"cookie,omitempty"`

	// The maximum amount of time to wait before the session cookie expires. Unit: seconds.
	CookieTimeout *float64 `json:"cookieTimeout,omitempty" tf:"cookie_timeout,omitempty"`

	// Specifies whether to enable session persistence. Valid values:
	StickySessionEnabled *bool `json:"stickySessionEnabled,omitempty" tf:"sticky_session_enabled,omitempty"`

	// The method that is used to handle a cookie. Valid values:
	StickySessionType *string `json:"stickySessionType,omitempty" tf:"sticky_session_type,omitempty"`
}

type StickySessionConfigParameters struct {

	// The cookie to be configured on the server.
	// +kubebuilder:validation:Optional
	Cookie *string `json:"cookie,omitempty" tf:"cookie,omitempty"`

	// The maximum amount of time to wait before the session cookie expires. Unit: seconds.
	// +kubebuilder:validation:Optional
	CookieTimeout *float64 `json:"cookieTimeout,omitempty" tf:"cookie_timeout,omitempty"`

	// Specifies whether to enable session persistence. Valid values:
	// +kubebuilder:validation:Optional
	StickySessionEnabled *bool `json:"stickySessionEnabled,omitempty" tf:"sticky_session_enabled,omitempty"`

	// The method that is used to handle a cookie. Valid values:
	// +kubebuilder:validation:Optional
	StickySessionType *string `json:"stickySessionType,omitempty" tf:"sticky_session_type,omitempty"`
}

type UchConfigInitParameters struct {

	// The parameter type. Only QueryString can be filled.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Consistency hash parameter value
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type UchConfigObservation struct {

	// The parameter type. Only QueryString can be filled.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Consistency hash parameter value
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type UchConfigParameters struct {

	// The parameter type. Only QueryString can be filled.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Consistency hash parameter value
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// ServerGroupSpec defines the desired state of ServerGroup
type ServerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerGroupParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ServerGroupInitParameters `json:"initProvider,omitempty"`
}

// ServerGroupStatus defines the observed state of ServerGroup.
type ServerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServerGroup is the Schema for the ServerGroups API. Provides a Alicloud Application Load Balancer (ALB) Server Group resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type ServerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.healthCheckConfig) || (has(self.initProvider) && has(self.initProvider.healthCheckConfig))",message="spec.forProvider.healthCheckConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serverGroupName) || (has(self.initProvider) && has(self.initProvider.serverGroupName))",message="spec.forProvider.serverGroupName is a required parameter"
	Spec   ServerGroupSpec   `json:"spec"`
	Status ServerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerGroupList contains a list of ServerGroups
type ServerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerGroup `json:"items"`
}

// Repository type metadata.
var (
	ServerGroup_Kind             = "ServerGroup"
	ServerGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerGroup_Kind}.String()
	ServerGroup_KindAPIVersion   = ServerGroup_Kind + "." + CRDGroupVersion.String()
	ServerGroup_GroupVersionKind = CRDGroupVersion.WithKind(ServerGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerGroup{}, &ServerGroupList{})
}
