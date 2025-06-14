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

type SessionManagerStatusInitParameters struct {

	// The name of the Session Manager Status. Valid values: sessionManagerStatus.
	SessionManagerStatusName *string `json:"sessionManagerStatusName,omitempty" tf:"session_manager_status_name,omitempty"`

	// The status of the Session Manager Status. Valid values: Enabled, Disabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SessionManagerStatusObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Session Manager Status. Valid values: sessionManagerStatus.
	SessionManagerStatusName *string `json:"sessionManagerStatusName,omitempty" tf:"session_manager_status_name,omitempty"`

	// The status of the Session Manager Status. Valid values: Enabled, Disabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SessionManagerStatusParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// The name of the Session Manager Status. Valid values: sessionManagerStatus.
	// +kubebuilder:validation:Optional
	SessionManagerStatusName *string `json:"sessionManagerStatusName,omitempty" tf:"session_manager_status_name,omitempty"`

	// The status of the Session Manager Status. Valid values: Enabled, Disabled.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// SessionManagerStatusSpec defines the desired state of SessionManagerStatus
type SessionManagerStatusSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SessionManagerStatusParameters `json:"forProvider"`
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
	InitProvider SessionManagerStatusInitParameters `json:"initProvider,omitempty"`
}

// SessionManagerStatusStatus defines the observed state of SessionManagerStatus.
type SessionManagerStatusStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SessionManagerStatusObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SessionManagerStatus is the Schema for the SessionManagerStatuss API. Provides a Alicloud ECS Session Manager Status resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type SessionManagerStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sessionManagerStatusName) || (has(self.initProvider) && has(self.initProvider.sessionManagerStatusName))",message="spec.forProvider.sessionManagerStatusName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.status) || (has(self.initProvider) && has(self.initProvider.status))",message="spec.forProvider.status is a required parameter"
	Spec   SessionManagerStatusSpec   `json:"spec"`
	Status SessionManagerStatusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SessionManagerStatusList contains a list of SessionManagerStatuss
type SessionManagerStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SessionManagerStatus `json:"items"`
}

// Repository type metadata.
var (
	SessionManagerStatus_Kind             = "SessionManagerStatus"
	SessionManagerStatus_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SessionManagerStatus_Kind}.String()
	SessionManagerStatus_KindAPIVersion   = SessionManagerStatus_Kind + "." + CRDGroupVersion.String()
	SessionManagerStatus_GroupVersionKind = CRDGroupVersion.WithKind(SessionManagerStatus_Kind)
)

func init() {
	SchemeBuilder.Register(&SessionManagerStatus{}, &SessionManagerStatusList{})
}
