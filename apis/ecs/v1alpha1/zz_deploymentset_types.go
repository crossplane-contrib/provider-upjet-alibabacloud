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

type DeploymentSetInitParameters struct {

	// The name of the deployment set. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
	DeploymentSetName *string `json:"deploymentSetName,omitempty" tf:"deployment_set_name,omitempty"`

	// The description of the deployment set. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The emergency solution to use in the situation where instances in the deployment set cannot be evenly distributed to different zones due to resource insufficiency after the instances failover. Valid values:
	OnUnableToRedeployFailedInstance *string `json:"onUnableToRedeployFailedInstance,omitempty" tf:"on_unable_to_redeploy_failed_instance,omitempty"`

	// The deployment strategy. Default value: Availability. Valid values: Availability, AvailabilityGroup, LowLatency.
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`
}

type DeploymentSetObservation struct {

	// The name of the deployment set. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
	DeploymentSetName *string `json:"deploymentSetName,omitempty" tf:"deployment_set_name,omitempty"`

	// The description of the deployment set. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The emergency solution to use in the situation where instances in the deployment set cannot be evenly distributed to different zones due to resource insufficiency after the instances failover. Valid values:
	OnUnableToRedeployFailedInstance *string `json:"onUnableToRedeployFailedInstance,omitempty" tf:"on_unable_to_redeploy_failed_instance,omitempty"`

	// The deployment strategy. Default value: Availability. Valid values: Availability, AvailabilityGroup, LowLatency.
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`
}

type DeploymentSetParameters struct {

	// The name of the deployment set. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
	// +kubebuilder:validation:Optional
	DeploymentSetName *string `json:"deploymentSetName,omitempty" tf:"deployment_set_name,omitempty"`

	// The description of the deployment set. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The emergency solution to use in the situation where instances in the deployment set cannot be evenly distributed to different zones due to resource insufficiency after the instances failover. Valid values:
	// +kubebuilder:validation:Optional
	OnUnableToRedeployFailedInstance *string `json:"onUnableToRedeployFailedInstance,omitempty" tf:"on_unable_to_redeploy_failed_instance,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// The deployment strategy. Default value: Availability. Valid values: Availability, AvailabilityGroup, LowLatency.
	// +kubebuilder:validation:Optional
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`
}

// DeploymentSetSpec defines the desired state of DeploymentSet
type DeploymentSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentSetParameters `json:"forProvider"`
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
	InitProvider DeploymentSetInitParameters `json:"initProvider,omitempty"`
}

// DeploymentSetStatus defines the observed state of DeploymentSet.
type DeploymentSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DeploymentSet is the Schema for the DeploymentSets API. Provides a Alicloud ECS Deployment Set resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type DeploymentSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentSetSpec   `json:"spec"`
	Status            DeploymentSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentSetList contains a list of DeploymentSets
type DeploymentSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentSet `json:"items"`
}

// Repository type metadata.
var (
	DeploymentSet_Kind             = "DeploymentSet"
	DeploymentSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeploymentSet_Kind}.String()
	DeploymentSet_KindAPIVersion   = DeploymentSet_Kind + "." + CRDGroupVersion.String()
	DeploymentSet_GroupVersionKind = CRDGroupVersion.WithKind(DeploymentSet_Kind)
)

func init() {
	SchemeBuilder.Register(&DeploymentSet{}, &DeploymentSetList{})
}
