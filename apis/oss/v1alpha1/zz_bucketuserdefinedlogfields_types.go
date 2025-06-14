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

type BucketUserDefinedLogFieldsInitParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/oss/v1alpha1.Bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Container for custom request header configuration information.
	// +listType=set
	HeaderSet []*string `json:"headerSet,omitempty" tf:"header_set,omitempty"`

	// Container for custom request parameters configuration information.
	// +listType=set
	ParamSet []*string `json:"paramSet,omitempty" tf:"param_set,omitempty"`
}

type BucketUserDefinedLogFieldsObservation struct {

	// The name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Container for custom request header configuration information.
	// +listType=set
	HeaderSet []*string `json:"headerSet,omitempty" tf:"header_set,omitempty"`

	// The ID of the resource supplied above.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Container for custom request parameters configuration information.
	// +listType=set
	ParamSet []*string `json:"paramSet,omitempty" tf:"param_set,omitempty"`
}

type BucketUserDefinedLogFieldsParameters struct {

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

	// Container for custom request header configuration information.
	// +kubebuilder:validation:Optional
	// +listType=set
	HeaderSet []*string `json:"headerSet,omitempty" tf:"header_set,omitempty"`

	// Container for custom request parameters configuration information.
	// +kubebuilder:validation:Optional
	// +listType=set
	ParamSet []*string `json:"paramSet,omitempty" tf:"param_set,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`
}

// BucketUserDefinedLogFieldsSpec defines the desired state of BucketUserDefinedLogFields
type BucketUserDefinedLogFieldsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketUserDefinedLogFieldsParameters `json:"forProvider"`
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
	InitProvider BucketUserDefinedLogFieldsInitParameters `json:"initProvider,omitempty"`
}

// BucketUserDefinedLogFieldsStatus defines the observed state of BucketUserDefinedLogFields.
type BucketUserDefinedLogFieldsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketUserDefinedLogFieldsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketUserDefinedLogFields is the Schema for the BucketUserDefinedLogFieldss API. Provides a Alicloud OSS Bucket User Defined Log Fields resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type BucketUserDefinedLogFields struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketUserDefinedLogFieldsSpec   `json:"spec"`
	Status            BucketUserDefinedLogFieldsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketUserDefinedLogFieldsList contains a list of BucketUserDefinedLogFieldss
type BucketUserDefinedLogFieldsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketUserDefinedLogFields `json:"items"`
}

// Repository type metadata.
var (
	BucketUserDefinedLogFields_Kind             = "BucketUserDefinedLogFields"
	BucketUserDefinedLogFields_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketUserDefinedLogFields_Kind}.String()
	BucketUserDefinedLogFields_KindAPIVersion   = BucketUserDefinedLogFields_Kind + "." + CRDGroupVersion.String()
	BucketUserDefinedLogFields_GroupVersionKind = CRDGroupVersion.WithKind(BucketUserDefinedLogFields_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketUserDefinedLogFields{}, &BucketUserDefinedLogFieldsList{})
}
