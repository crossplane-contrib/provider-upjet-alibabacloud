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

type DomainAttachmentInitParameters struct {

	// The domain names bound to the DNS instance.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/alidns/v1alpha1.Domain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("domain_name",false)
	// +listType=set
	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`

	// References to Domain in alidns to populate domainNames.
	// +kubebuilder:validation:Optional
	DomainNamesRefs []v1.Reference `json:"domainNamesRefs,omitempty" tf:"-"`

	// Selector for a list of Domain in alidns to populate domainNames.
	// +kubebuilder:validation:Optional
	DomainNamesSelector *v1.Selector `json:"domainNamesSelector,omitempty" tf:"-"`

	// The id of the DNS instance.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/alidns/v1alpha1.Instance
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in alidns to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in alidns to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`
}

type DomainAttachmentObservation struct {

	// The domain names bound to the DNS instance.
	// +listType=set
	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`

	// This ID of this resource. The value is same as instance_id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The id of the DNS instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`
}

type DomainAttachmentParameters struct {

	// The domain names bound to the DNS instance.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/alidns/v1alpha1.Domain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("domain_name",false)
	// +kubebuilder:validation:Optional
	// +listType=set
	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`

	// References to Domain in alidns to populate domainNames.
	// +kubebuilder:validation:Optional
	DomainNamesRefs []v1.Reference `json:"domainNamesRefs,omitempty" tf:"-"`

	// Selector for a list of Domain in alidns to populate domainNames.
	// +kubebuilder:validation:Optional
	DomainNamesSelector *v1.Selector `json:"domainNamesSelector,omitempty" tf:"-"`

	// The id of the DNS instance.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/alidns/v1alpha1.Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in alidns to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in alidns to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`
}

// DomainAttachmentSpec defines the desired state of DomainAttachment
type DomainAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainAttachmentParameters `json:"forProvider"`
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
	InitProvider DomainAttachmentInitParameters `json:"initProvider,omitempty"`
}

// DomainAttachmentStatus defines the observed state of DomainAttachment.
type DomainAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DomainAttachment is the Schema for the DomainAttachments API. Provides bind the domain name to the Alidns instance resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type DomainAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainAttachmentSpec   `json:"spec"`
	Status            DomainAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainAttachmentList contains a list of DomainAttachments
type DomainAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainAttachment `json:"items"`
}

// Repository type metadata.
var (
	DomainAttachment_Kind             = "DomainAttachment"
	DomainAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainAttachment_Kind}.String()
	DomainAttachment_KindAPIVersion   = DomainAttachment_Kind + "." + CRDGroupVersion.String()
	DomainAttachment_GroupVersionKind = CRDGroupVersion.WithKind(DomainAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainAttachment{}, &DomainAttachmentList{})
}
