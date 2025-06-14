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

type AccessLogConfigInitParameters struct {

	// The project to which the access log is shipped.
	LogProject *string `json:"logProject,omitempty" tf:"log_project,omitempty"`

	// The Logstore to which the access log is shipped.
	LogStore *string `json:"logStore,omitempty" tf:"log_store,omitempty"`
}

type AccessLogConfigObservation struct {

	// The project to which the access log is shipped.
	LogProject *string `json:"logProject,omitempty" tf:"log_project,omitempty"`

	// The Logstore to which the access log is shipped.
	LogStore *string `json:"logStore,omitempty" tf:"log_store,omitempty"`
}

type AccessLogConfigParameters struct {

	// The project to which the access log is shipped.
	// +kubebuilder:validation:Optional
	LogProject *string `json:"logProject,omitempty" tf:"log_project,omitempty"`

	// The Logstore to which the access log is shipped.
	// +kubebuilder:validation:Optional
	LogStore *string `json:"logStore,omitempty" tf:"log_store,omitempty"`
}

type DeletionProtectionConfigInitParameters struct {

	// Remove the Protection Status
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DeletionProtectionConfigObservation struct {

	// Remove the Protection Status
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Deletion Protection Turn-on Time Use Greenwich Mean Time, in the Format of Yyyy-MM-ddTHH: mm: SSZ
	EnabledTime *string `json:"enabledTime,omitempty" tf:"enabled_time,omitempty"`
}

type DeletionProtectionConfigParameters struct {

	// Remove the Protection Status
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LoadBalancerAddressesInitParameters struct {
}

type LoadBalancerAddressesObservation struct {

	// IP Address. The Public IP Address, and Private IP Address from the Address Type
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The ID of the EIP instance.
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// The type of the EIP instance.
	EIPType *string `json:"eipType,omitempty" tf:"eip_type,omitempty"`

	// IPv4 Local address list. The list of addresses used by ALB to interact with the backend service.
	IPv4LocalAddresses []*string `json:"ipv4LocalAddresses,omitempty" tf:"ipv4_local_addresses,omitempty"`

	// Ipv6 address
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// The IPv6 address detection status of the application-based load balancing instance.
	IPv6AddressHcStatus *string `json:"ipv6AddressHcStatus,omitempty" tf:"ipv6_address_hc_status,omitempty"`

	// IPv6 Local address list. The list of addresses used by ALB to interact with the backend service.
	IPv6LocalAddresses []*string `json:"ipv6LocalAddresses,omitempty" tf:"ipv6_local_addresses,omitempty"`

	// IPv4 private network address.
	IntranetAddress *string `json:"intranetAddress,omitempty" tf:"intranet_address,omitempty"`

	// The private network IPv4 address detection status of the application-oriented load balancing instance.
	IntranetAddressHcStatus *string `json:"intranetAddressHcStatus,omitempty" tf:"intranet_address_hc_status,omitempty"`
}

type LoadBalancerAddressesParameters struct {
}

type LoadBalancerBillingConfigInitParameters struct {

	// Pay Type
	PayType *string `json:"payType,omitempty" tf:"pay_type,omitempty"`
}

type LoadBalancerBillingConfigObservation struct {

	// Pay Type
	PayType *string `json:"payType,omitempty" tf:"pay_type,omitempty"`
}

type LoadBalancerBillingConfigParameters struct {

	// Pay Type
	// +kubebuilder:validation:Optional
	PayType *string `json:"payType" tf:"pay_type,omitempty"`
}

type LoadBalancerInitParameters struct {

	// The configuration of the access log. See access_log_config below.
	AccessLogConfig []AccessLogConfigInitParameters `json:"accessLogConfig,omitempty" tf:"access_log_config,omitempty"`

	// The method in which IP addresses are assigned. Valid values:  Fixed: The ALB instance uses a fixed IP address. Dynamic (default): An IP address is dynamically assigned to each zone of the ALB instance.
	AddressAllocatedMode *string `json:"addressAllocatedMode,omitempty" tf:"address_allocated_mode,omitempty"`

	// The protocol version. Value:
	AddressIPVersion *string `json:"addressIpVersion,omitempty" tf:"address_ip_version,omitempty"`

	// The type of IP address that the SLB instance uses to provide services.
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// The ID of the EIP bandwidth plan which is associated with an ALB instance that uses a public IP address.
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// The Protection Configuration See deletion_protection_config below.
	DeletionProtectionConfig []DeletionProtectionConfigInitParameters `json:"deletionProtectionConfig,omitempty" tf:"deletion_protection_config,omitempty"`

	// Specifies whether to enable deletion protection. Default value: false. Valid values:
	DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled,omitempty" tf:"deletion_protection_enabled,omitempty"`

	// Whether to PreCheck only this request, value:
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// The address type of Ipv6
	IPv6AddressType *string `json:"ipv6AddressType,omitempty" tf:"ipv6_address_type,omitempty"`

	// The configuration of the billing method. See load_balancer_billing_config below.
	LoadBalancerBillingConfig []LoadBalancerBillingConfigInitParameters `json:"loadBalancerBillingConfig,omitempty" tf:"load_balancer_billing_config,omitempty"`

	// The edition of the ALB instance.
	LoadBalancerEdition *string `json:"loadBalancerEdition,omitempty" tf:"load_balancer_edition,omitempty"`

	// The name of the resource
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// Modify the Protection Configuration See modification_protection_config below.
	ModificationProtectionConfig []ModificationProtectionConfigInitParameters `json:"modificationProtectionConfig,omitempty" tf:"modification_protection_config,omitempty"`

	// The ID of the resource group
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the virtual private cloud (VPC) where the SLB instance is deployed.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/vpc/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The zones and vSwitches. You must specify at least two zones. See zone_mappings below.
	ZoneMappings []ZoneMappingsInitParameters `json:"zoneMappings,omitempty" tf:"zone_mappings,omitempty"`
}

type LoadBalancerObservation struct {

	// The configuration of the access log. See access_log_config below.
	AccessLogConfig []AccessLogConfigObservation `json:"accessLogConfig,omitempty" tf:"access_log_config,omitempty"`

	// The method in which IP addresses are assigned. Valid values:  Fixed: The ALB instance uses a fixed IP address. Dynamic (default): An IP address is dynamically assigned to each zone of the ALB instance.
	AddressAllocatedMode *string `json:"addressAllocatedMode,omitempty" tf:"address_allocated_mode,omitempty"`

	// The protocol version. Value:
	AddressIPVersion *string `json:"addressIpVersion,omitempty" tf:"address_ip_version,omitempty"`

	// The type of IP address that the SLB instance uses to provide services.
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// The ID of the EIP bandwidth plan which is associated with an ALB instance that uses a public IP address.
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// The creation time of the resource
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// DNS Domain Name
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// The Protection Configuration See deletion_protection_config below.
	DeletionProtectionConfig []DeletionProtectionConfigObservation `json:"deletionProtectionConfig,omitempty" tf:"deletion_protection_config,omitempty"`

	// Specifies whether to enable deletion protection. Default value: false. Valid values:
	DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled,omitempty" tf:"deletion_protection_enabled,omitempty"`

	// Whether to PreCheck only this request, value:
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// The ID of the resource supplied above.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The address type of Ipv6
	IPv6AddressType *string `json:"ipv6AddressType,omitempty" tf:"ipv6_address_type,omitempty"`

	// The configuration of the billing method. See load_balancer_billing_config below.
	LoadBalancerBillingConfig []LoadBalancerBillingConfigObservation `json:"loadBalancerBillingConfig,omitempty" tf:"load_balancer_billing_config,omitempty"`

	// The edition of the ALB instance.
	LoadBalancerEdition *string `json:"loadBalancerEdition,omitempty" tf:"load_balancer_edition,omitempty"`

	// The name of the resource
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// Modify the Protection Configuration See modification_protection_config below.
	ModificationProtectionConfig []ModificationProtectionConfigObservation `json:"modificationProtectionConfig,omitempty" tf:"modification_protection_config,omitempty"`

	// The region ID of the resource
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// The ID of the resource group
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// Server Load Balancer Instance Status:, indicating that the instance listener will no longer forward traffic.(default).
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the virtual private cloud (VPC) where the SLB instance is deployed.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The zones and vSwitches. You must specify at least two zones. See zone_mappings below.
	ZoneMappings []ZoneMappingsObservation `json:"zoneMappings,omitempty" tf:"zone_mappings,omitempty"`
}

type LoadBalancerParameters struct {

	// The configuration of the access log. See access_log_config below.
	// +kubebuilder:validation:Optional
	AccessLogConfig []AccessLogConfigParameters `json:"accessLogConfig,omitempty" tf:"access_log_config,omitempty"`

	// The method in which IP addresses are assigned. Valid values:  Fixed: The ALB instance uses a fixed IP address. Dynamic (default): An IP address is dynamically assigned to each zone of the ALB instance.
	// +kubebuilder:validation:Optional
	AddressAllocatedMode *string `json:"addressAllocatedMode,omitempty" tf:"address_allocated_mode,omitempty"`

	// The protocol version. Value:
	// +kubebuilder:validation:Optional
	AddressIPVersion *string `json:"addressIpVersion,omitempty" tf:"address_ip_version,omitempty"`

	// The type of IP address that the SLB instance uses to provide services.
	// +kubebuilder:validation:Optional
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// The ID of the EIP bandwidth plan which is associated with an ALB instance that uses a public IP address.
	// +kubebuilder:validation:Optional
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// The Protection Configuration See deletion_protection_config below.
	// +kubebuilder:validation:Optional
	DeletionProtectionConfig []DeletionProtectionConfigParameters `json:"deletionProtectionConfig,omitempty" tf:"deletion_protection_config,omitempty"`

	// Specifies whether to enable deletion protection. Default value: false. Valid values:
	// +kubebuilder:validation:Optional
	DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled,omitempty" tf:"deletion_protection_enabled,omitempty"`

	// Whether to PreCheck only this request, value:
	// +kubebuilder:validation:Optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run,omitempty"`

	// The address type of Ipv6
	// +kubebuilder:validation:Optional
	IPv6AddressType *string `json:"ipv6AddressType,omitempty" tf:"ipv6_address_type,omitempty"`

	// The configuration of the billing method. See load_balancer_billing_config below.
	// +kubebuilder:validation:Optional
	LoadBalancerBillingConfig []LoadBalancerBillingConfigParameters `json:"loadBalancerBillingConfig,omitempty" tf:"load_balancer_billing_config,omitempty"`

	// The edition of the ALB instance.
	// +kubebuilder:validation:Optional
	LoadBalancerEdition *string `json:"loadBalancerEdition,omitempty" tf:"load_balancer_edition,omitempty"`

	// The name of the resource
	// +kubebuilder:validation:Optional
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// Modify the Protection Configuration See modification_protection_config below.
	// +kubebuilder:validation:Optional
	ModificationProtectionConfig []ModificationProtectionConfigParameters `json:"modificationProtectionConfig,omitempty" tf:"modification_protection_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// The ID of the resource group
	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the virtual private cloud (VPC) where the SLB instance is deployed.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The zones and vSwitches. You must specify at least two zones. See zone_mappings below.
	// +kubebuilder:validation:Optional
	ZoneMappings []ZoneMappingsParameters `json:"zoneMappings,omitempty" tf:"zone_mappings,omitempty"`
}

type ModificationProtectionConfigInitParameters struct {

	// Managed Instance
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	// Load Balancing Modify the Protection Status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ModificationProtectionConfigObservation struct {

	// Managed Instance
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	// Load Balancing Modify the Protection Status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ModificationProtectionConfigParameters struct {

	// Managed Instance
	// +kubebuilder:validation:Optional
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	// Load Balancing Modify the Protection Status
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ZoneMappingsInitParameters struct {

	// The ID of the EIP instance.
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// The type of the EIP instance.
	EIPType *string `json:"eipType,omitempty" tf:"eip_type,omitempty"`

	// IPv4 private network address.
	IntranetAddress *string `json:"intranetAddress,omitempty" tf:"intranet_address,omitempty"`

	// The ID of the vSwitch that corresponds to the zone. Each zone can use only one vSwitch and subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/vpc/v1alpha1.Vswitch
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VswitchID *string `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`

	// Reference to a Vswitch in vpc to populate vswitchId.
	// +kubebuilder:validation:Optional
	VswitchIDRef *v1.Reference `json:"vswitchIdRef,omitempty" tf:"-"`

	// Selector for a Vswitch in vpc to populate vswitchId.
	// +kubebuilder:validation:Optional
	VswitchIDSelector *v1.Selector `json:"vswitchIdSelector,omitempty" tf:"-"`

	// The ID of the zone to which the SLB instance belongs.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ZoneMappingsObservation struct {

	// IP Address. The Public IP Address, and Private IP Address from the Address Type
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The ID of the EIP instance.
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// The type of the EIP instance.
	EIPType *string `json:"eipType,omitempty" tf:"eip_type,omitempty"`

	// Ipv6 address
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// IPv4 private network address.
	IntranetAddress *string `json:"intranetAddress,omitempty" tf:"intranet_address,omitempty"`

	// The instance address.
	LoadBalancerAddresses []LoadBalancerAddressesObservation `json:"loadBalancerAddresses,omitempty" tf:"load_balancer_addresses,omitempty"`

	// The ID of the vSwitch that corresponds to the zone. Each zone can use only one vSwitch and subnet.
	VswitchID *string `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`

	// The ID of the zone to which the SLB instance belongs.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ZoneMappingsParameters struct {

	// The ID of the EIP instance.
	// +kubebuilder:validation:Optional
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// The type of the EIP instance.
	// +kubebuilder:validation:Optional
	EIPType *string `json:"eipType,omitempty" tf:"eip_type,omitempty"`

	// IPv4 private network address.
	// +kubebuilder:validation:Optional
	IntranetAddress *string `json:"intranetAddress,omitempty" tf:"intranet_address,omitempty"`

	// The ID of the vSwitch that corresponds to the zone. Each zone can use only one vSwitch and subnet.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/vpc/v1alpha1.Vswitch
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VswitchID *string `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`

	// Reference to a Vswitch in vpc to populate vswitchId.
	// +kubebuilder:validation:Optional
	VswitchIDRef *v1.Reference `json:"vswitchIdRef,omitempty" tf:"-"`

	// Selector for a Vswitch in vpc to populate vswitchId.
	// +kubebuilder:validation:Optional
	VswitchIDSelector *v1.Selector `json:"vswitchIdSelector,omitempty" tf:"-"`

	// The ID of the zone to which the SLB instance belongs.
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId" tf:"zone_id,omitempty"`
}

// LoadBalancerSpec defines the desired state of LoadBalancer
type LoadBalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerParameters `json:"forProvider"`
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
	InitProvider LoadBalancerInitParameters `json:"initProvider,omitempty"`
}

// LoadBalancerStatus defines the observed state of LoadBalancer.
type LoadBalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LoadBalancer is the Schema for the LoadBalancers API. Provides a Alicloud Application Load Balancer (ALB) Load Balancer resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addressType) || (has(self.initProvider) && has(self.initProvider.addressType))",message="spec.forProvider.addressType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.loadBalancerBillingConfig) || (has(self.initProvider) && has(self.initProvider.loadBalancerBillingConfig))",message="spec.forProvider.loadBalancerBillingConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.loadBalancerEdition) || (has(self.initProvider) && has(self.initProvider.loadBalancerEdition))",message="spec.forProvider.loadBalancerEdition is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zoneMappings) || (has(self.initProvider) && has(self.initProvider.zoneMappings))",message="spec.forProvider.zoneMappings is a required parameter"
	Spec   LoadBalancerSpec   `json:"spec"`
	Status LoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerList contains a list of LoadBalancers
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancer_Kind             = "LoadBalancer"
	LoadBalancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancer_Kind}.String()
	LoadBalancer_KindAPIVersion   = LoadBalancer_Kind + "." + CRDGroupVersion.String()
	LoadBalancer_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancer_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
