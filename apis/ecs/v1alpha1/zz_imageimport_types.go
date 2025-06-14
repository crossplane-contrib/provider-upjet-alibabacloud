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

type ImageImportDiskDeviceMappingInitParameters struct {

	// The device name of the disk.
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	// The size of the disk. Default value: 5.
	DiskImageSize *float64 `json:"diskImageSize,omitempty" tf:"disk_image_size,omitempty"`

	// The format of the image. Valid values: RAW, VHD, qcow2.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The OSS bucket where the image file is stored.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/oss/v1alpha1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	OssBucket *string `json:"ossBucket,omitempty" tf:"oss_bucket,omitempty"`

	// Reference to a Bucket in oss to populate ossBucket.
	// +kubebuilder:validation:Optional
	OssBucketRef *v1.Reference `json:"ossBucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in oss to populate ossBucket.
	// +kubebuilder:validation:Optional
	OssBucketSelector *v1.Selector `json:"ossBucketSelector,omitempty" tf:"-"`

	// The name (key) of the object that the uploaded image is stored as in the OSS bucket.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/oss/v1alpha1.BucketObject
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	OssObject *string `json:"ossObject,omitempty" tf:"oss_object,omitempty"`

	// Reference to a BucketObject in oss to populate ossObject.
	// +kubebuilder:validation:Optional
	OssObjectRef *v1.Reference `json:"ossObjectRef,omitempty" tf:"-"`

	// Selector for a BucketObject in oss to populate ossObject.
	// +kubebuilder:validation:Optional
	OssObjectSelector *v1.Selector `json:"ossObjectSelector,omitempty" tf:"-"`
}

type ImageImportDiskDeviceMappingObservation struct {

	// The device name of the disk.
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	// The size of the disk. Default value: 5.
	DiskImageSize *float64 `json:"diskImageSize,omitempty" tf:"disk_image_size,omitempty"`

	// The format of the image. Valid values: RAW, VHD, qcow2.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The OSS bucket where the image file is stored.
	OssBucket *string `json:"ossBucket,omitempty" tf:"oss_bucket,omitempty"`

	// The name (key) of the object that the uploaded image is stored as in the OSS bucket.
	OssObject *string `json:"ossObject,omitempty" tf:"oss_object,omitempty"`
}

type ImageImportDiskDeviceMappingParameters struct {

	// The device name of the disk.
	// +kubebuilder:validation:Optional
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	// The size of the disk. Default value: 5.
	// +kubebuilder:validation:Optional
	DiskImageSize *float64 `json:"diskImageSize,omitempty" tf:"disk_image_size,omitempty"`

	// The format of the image. Valid values: RAW, VHD, qcow2.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The OSS bucket where the image file is stored.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/oss/v1alpha1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	OssBucket *string `json:"ossBucket,omitempty" tf:"oss_bucket,omitempty"`

	// Reference to a Bucket in oss to populate ossBucket.
	// +kubebuilder:validation:Optional
	OssBucketRef *v1.Reference `json:"ossBucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in oss to populate ossBucket.
	// +kubebuilder:validation:Optional
	OssBucketSelector *v1.Selector `json:"ossBucketSelector,omitempty" tf:"-"`

	// The name (key) of the object that the uploaded image is stored as in the OSS bucket.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/oss/v1alpha1.BucketObject
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	OssObject *string `json:"ossObject,omitempty" tf:"oss_object,omitempty"`

	// Reference to a BucketObject in oss to populate ossObject.
	// +kubebuilder:validation:Optional
	OssObjectRef *v1.Reference `json:"ossObjectRef,omitempty" tf:"-"`

	// Selector for a BucketObject in oss to populate ossObject.
	// +kubebuilder:validation:Optional
	OssObjectSelector *v1.Selector `json:"ossObjectSelector,omitempty" tf:"-"`
}

type ImageImportInitParameters struct {

	// The architecture of the image. Default value: x86_64. Valid values: x86_64, i386.
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// The boot mode of the image. Valid values: BIOS, UEFI.
	BootMode *string `json:"bootMode,omitempty" tf:"boot_mode,omitempty"`

	// The description of the image. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The information about the custom image. See disk_device_mapping below.
	DiskDeviceMapping []ImageImportDiskDeviceMappingInitParameters `json:"diskDeviceMapping,omitempty" tf:"disk_device_mapping,omitempty"`

	// The name of the image. The image_name must be 2 to 128 characters in length. The image_name must start with a letter and cannot start with acs: or aliyun. The image_name cannot contain http:// or https://. The image_name can contain letters, digits, periods (.), colons (:), underscores (_), and hyphens (-).
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// The type of the license used to activate the operating system after the image is imported. Default value: Auto. Valid values: Auto, Aliyun, BYOL.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// The type of the operating system. Default value: linux. Valid values: windows, linux.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The operating system platform. More valid values refer to ImportImage OpenAPI.
	// -> NOTE: Before provider version 1.197.0, the default value of platform is Ubuntu.
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`
}

type ImageImportObservation struct {

	// The architecture of the image. Default value: x86_64. Valid values: x86_64, i386.
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// The boot mode of the image. Valid values: BIOS, UEFI.
	BootMode *string `json:"bootMode,omitempty" tf:"boot_mode,omitempty"`

	// The description of the image. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The information about the custom image. See disk_device_mapping below.
	DiskDeviceMapping []ImageImportDiskDeviceMappingObservation `json:"diskDeviceMapping,omitempty" tf:"disk_device_mapping,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the image. The image_name must be 2 to 128 characters in length. The image_name must start with a letter and cannot start with acs: or aliyun. The image_name cannot contain http:// or https://. The image_name can contain letters, digits, periods (.), colons (:), underscores (_), and hyphens (-).
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// The type of the license used to activate the operating system after the image is imported. Default value: Auto. Valid values: Auto, Aliyun, BYOL.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// The type of the operating system. Default value: linux. Valid values: windows, linux.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The operating system platform. More valid values refer to ImportImage OpenAPI.
	// -> NOTE: Before provider version 1.197.0, the default value of platform is Ubuntu.
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`
}

type ImageImportParameters struct {

	// The architecture of the image. Default value: x86_64. Valid values: x86_64, i386.
	// +kubebuilder:validation:Optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// The boot mode of the image. Valid values: BIOS, UEFI.
	// +kubebuilder:validation:Optional
	BootMode *string `json:"bootMode,omitempty" tf:"boot_mode,omitempty"`

	// The description of the image. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The information about the custom image. See disk_device_mapping below.
	// +kubebuilder:validation:Optional
	DiskDeviceMapping []ImageImportDiskDeviceMappingParameters `json:"diskDeviceMapping,omitempty" tf:"disk_device_mapping,omitempty"`

	// The name of the image. The image_name must be 2 to 128 characters in length. The image_name must start with a letter and cannot start with acs: or aliyun. The image_name cannot contain http:// or https://. The image_name can contain letters, digits, periods (.), colons (:), underscores (_), and hyphens (-).
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// The type of the license used to activate the operating system after the image is imported. Default value: Auto. Valid values: Auto, Aliyun, BYOL.
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// The type of the operating system. Default value: linux. Valid values: windows, linux.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The operating system platform. More valid values refer to ImportImage OpenAPI.
	// -> NOTE: Before provider version 1.197.0, the default value of platform is Ubuntu.
	// +kubebuilder:validation:Optional
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`
}

// ImageImportSpec defines the desired state of ImageImport
type ImageImportSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageImportParameters `json:"forProvider"`
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
	InitProvider ImageImportInitParameters `json:"initProvider,omitempty"`
}

// ImageImportStatus defines the observed state of ImageImport.
type ImageImportStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageImportObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ImageImport is the Schema for the ImageImports API. Provides a Alicloud ECS Image Import resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type ImageImport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.diskDeviceMapping) || (has(self.initProvider) && has(self.initProvider.diskDeviceMapping))",message="spec.forProvider.diskDeviceMapping is a required parameter"
	Spec   ImageImportSpec   `json:"spec"`
	Status ImageImportStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageImportList contains a list of ImageImports
type ImageImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageImport `json:"items"`
}

// Repository type metadata.
var (
	ImageImport_Kind             = "ImageImport"
	ImageImport_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageImport_Kind}.String()
	ImageImport_KindAPIVersion   = ImageImport_Kind + "." + CRDGroupVersion.String()
	ImageImport_GroupVersionKind = CRDGroupVersion.WithKind(ImageImport_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageImport{}, &ImageImportList{})
}
