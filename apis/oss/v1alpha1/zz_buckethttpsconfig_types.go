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

type BucketHTTPSConfigInitParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/oss/v1alpha1.Bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Specifies whether to enable TLS version management for the bucket. Valid values: true, false.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the TLS versions allowed to access this buckets.
	// +listType=set
	TLSVersions []*string `json:"tlsVersions,omitempty" tf:"tls_versions,omitempty"`
}

type BucketHTTPSConfigObservation struct {

	// The name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Specifies whether to enable TLS version management for the bucket. Valid values: true, false.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// The ID of the resource supplied above.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the TLS versions allowed to access this buckets.
	// +listType=set
	TLSVersions []*string `json:"tlsVersions,omitempty" tf:"tls_versions,omitempty"`
}

type BucketHTTPSConfigParameters struct {

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

	// Specifies whether to enable TLS version management for the bucket. Valid values: true, false.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// Specifies the TLS versions allowed to access this buckets.
	// +kubebuilder:validation:Optional
	// +listType=set
	TLSVersions []*string `json:"tlsVersions,omitempty" tf:"tls_versions,omitempty"`
}

// BucketHTTPSConfigSpec defines the desired state of BucketHTTPSConfig
type BucketHTTPSConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketHTTPSConfigParameters `json:"forProvider"`
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
	InitProvider BucketHTTPSConfigInitParameters `json:"initProvider,omitempty"`
}

// BucketHTTPSConfigStatus defines the observed state of BucketHTTPSConfig.
type BucketHTTPSConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketHTTPSConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketHTTPSConfig is the Schema for the BucketHTTPSConfigs API. Provides a Alicloud OSS Bucket Https Config resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type BucketHTTPSConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enable) || (has(self.initProvider) && has(self.initProvider.enable))",message="spec.forProvider.enable is a required parameter"
	Spec   BucketHTTPSConfigSpec   `json:"spec"`
	Status BucketHTTPSConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketHTTPSConfigList contains a list of BucketHTTPSConfigs
type BucketHTTPSConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketHTTPSConfig `json:"items"`
}

// Repository type metadata.
var (
	BucketHTTPSConfig_Kind             = "BucketHTTPSConfig"
	BucketHTTPSConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketHTTPSConfig_Kind}.String()
	BucketHTTPSConfig_KindAPIVersion   = BucketHTTPSConfig_Kind + "." + CRDGroupVersion.String()
	BucketHTTPSConfig_GroupVersionKind = CRDGroupVersion.WithKind(BucketHTTPSConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketHTTPSConfig{}, &BucketHTTPSConfigList{})
}
