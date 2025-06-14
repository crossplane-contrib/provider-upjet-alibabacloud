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

type KeyPairAttachmentInitParameters struct {

	// Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// The list of ECS instance's IDs.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ecs/v1alpha1.Instance
	// +crossplane:generate:reference:refFieldName=InstanceRefs
	// +crossplane:generate:reference:selectorFieldName=InstanceSelector
	// +listType=set
	InstanceIds []*string `json:"instanceIds,omitempty" tf:"instance_ids,omitempty"`

	// References to Instance in ecs to populate instanceIds.
	// +kubebuilder:validation:Optional
	InstanceRefs []v1.Reference `json:"instanceRefs,omitempty" tf:"-"`

	// Selector for a list of Instance in ecs to populate instanceIds.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// The name of key pair used to bind.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ecs/v1alpha1.KeyPair
	KeyPairName *string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`

	// Reference to a KeyPair in ecs to populate keyPairName.
	// +kubebuilder:validation:Optional
	KeyPairNameRef *v1.Reference `json:"keyPairNameRef,omitempty" tf:"-"`

	// Selector for a KeyPair in ecs to populate keyPairName.
	// +kubebuilder:validation:Optional
	KeyPairNameSelector *v1.Selector `json:"keyPairNameSelector,omitempty" tf:"-"`
}

type KeyPairAttachmentObservation struct {

	// Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// The resource ID of Key Pair Attachment. The value is formatted <key_pair_name>:<instance_ids>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The list of ECS instance's IDs.
	// +listType=set
	InstanceIds []*string `json:"instanceIds,omitempty" tf:"instance_ids,omitempty"`

	// The name of key pair used to bind.
	KeyPairName *string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`
}

type KeyPairAttachmentParameters struct {

	// Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// The list of ECS instance's IDs.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ecs/v1alpha1.Instance
	// +crossplane:generate:reference:refFieldName=InstanceRefs
	// +crossplane:generate:reference:selectorFieldName=InstanceSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	InstanceIds []*string `json:"instanceIds,omitempty" tf:"instance_ids,omitempty"`

	// References to Instance in ecs to populate instanceIds.
	// +kubebuilder:validation:Optional
	InstanceRefs []v1.Reference `json:"instanceRefs,omitempty" tf:"-"`

	// Selector for a list of Instance in ecs to populate instanceIds.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// The name of key pair used to bind.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ecs/v1alpha1.KeyPair
	// +kubebuilder:validation:Optional
	KeyPairName *string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`

	// Reference to a KeyPair in ecs to populate keyPairName.
	// +kubebuilder:validation:Optional
	KeyPairNameRef *v1.Reference `json:"keyPairNameRef,omitempty" tf:"-"`

	// Selector for a KeyPair in ecs to populate keyPairName.
	// +kubebuilder:validation:Optional
	KeyPairNameSelector *v1.Selector `json:"keyPairNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`
}

// KeyPairAttachmentSpec defines the desired state of KeyPairAttachment
type KeyPairAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyPairAttachmentParameters `json:"forProvider"`
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
	InitProvider KeyPairAttachmentInitParameters `json:"initProvider,omitempty"`
}

// KeyPairAttachmentStatus defines the observed state of KeyPairAttachment.
type KeyPairAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyPairAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// KeyPairAttachment is the Schema for the KeyPairAttachments API. Provides a Alicloud ECS Key Pair Attachment resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type KeyPairAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyPairAttachmentSpec   `json:"spec"`
	Status            KeyPairAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyPairAttachmentList contains a list of KeyPairAttachments
type KeyPairAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyPairAttachment `json:"items"`
}

// Repository type metadata.
var (
	KeyPairAttachment_Kind             = "KeyPairAttachment"
	KeyPairAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeyPairAttachment_Kind}.String()
	KeyPairAttachment_KindAPIVersion   = KeyPairAttachment_Kind + "." + CRDGroupVersion.String()
	KeyPairAttachment_GroupVersionKind = CRDGroupVersion.WithKind(KeyPairAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&KeyPairAttachment{}, &KeyPairAttachmentList{})
}
