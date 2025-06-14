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

type BucketVersioningInitParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/oss/v1alpha1.Bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BucketVersioningObservation struct {

	// The name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The ID of the resource supplied above.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BucketVersioningParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/oss/v1alpha1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// BucketVersioningSpec defines the desired state of BucketVersioning
type BucketVersioningSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketVersioningParameters `json:"forProvider"`
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
	InitProvider BucketVersioningInitParameters `json:"initProvider,omitempty"`
}

// BucketVersioningStatus defines the observed state of BucketVersioning.
type BucketVersioningStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketVersioningObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketVersioning is the Schema for the BucketVersionings API. Provides a Alicloud OSS Bucket Versioning resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type BucketVersioning struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketVersioningSpec   `json:"spec"`
	Status            BucketVersioningStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketVersioningList contains a list of BucketVersionings
type BucketVersioningList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketVersioning `json:"items"`
}

// Repository type metadata.
var (
	BucketVersioning_Kind             = "BucketVersioning"
	BucketVersioning_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketVersioning_Kind}.String()
	BucketVersioning_KindAPIVersion   = BucketVersioning_Kind + "." + CRDGroupVersion.String()
	BucketVersioning_GroupVersionKind = CRDGroupVersion.WithKind(BucketVersioning_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketVersioning{}, &BucketVersioningList{})
}
