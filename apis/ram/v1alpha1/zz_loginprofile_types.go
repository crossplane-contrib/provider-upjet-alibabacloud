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

type LoginProfileInitParameters struct {

	// Specifies whether to forcefully enable multi-factor authentication (MFA) for the RAM user. Valid values:
	MfaBindRequired *bool `json:"mfaBindRequired,omitempty" tf:"mfa_bind_required,omitempty"`

	// The password must meet the Password strength requirements. For more information about password strength setting requirements, see GetPasswordPolicy.
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// Whether the user must reset the password at the next logon. Value:
	PasswordResetRequired *bool `json:"passwordResetRequired,omitempty" tf:"password_reset_required,omitempty"`

	// The user name.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ram/v1alpha1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Reference to a User in ram to populate userName.
	// +kubebuilder:validation:Optional
	UserNameRef *v1.Reference `json:"userNameRef,omitempty" tf:"-"`

	// Selector for a User in ram to populate userName.
	// +kubebuilder:validation:Optional
	UserNameSelector *v1.Selector `json:"userNameSelector,omitempty" tf:"-"`
}

type LoginProfileObservation struct {

	// Creation time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The ID of the resource supplied above.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether to forcefully enable multi-factor authentication (MFA) for the RAM user. Valid values:
	MfaBindRequired *bool `json:"mfaBindRequired,omitempty" tf:"mfa_bind_required,omitempty"`

	// The password must meet the Password strength requirements. For more information about password strength setting requirements, see GetPasswordPolicy.
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// Whether the user must reset the password at the next logon. Value:
	PasswordResetRequired *bool `json:"passwordResetRequired,omitempty" tf:"password_reset_required,omitempty"`

	// The user name.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type LoginProfileParameters struct {

	// Specifies whether to forcefully enable multi-factor authentication (MFA) for the RAM user. Valid values:
	// +kubebuilder:validation:Optional
	MfaBindRequired *bool `json:"mfaBindRequired,omitempty" tf:"mfa_bind_required,omitempty"`

	// The password must meet the Password strength requirements. For more information about password strength setting requirements, see GetPasswordPolicy.
	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// Whether the user must reset the password at the next logon. Value:
	// +kubebuilder:validation:Optional
	PasswordResetRequired *bool `json:"passwordResetRequired,omitempty" tf:"password_reset_required,omitempty"`

	// The user name.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ram/v1alpha1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Reference to a User in ram to populate userName.
	// +kubebuilder:validation:Optional
	UserNameRef *v1.Reference `json:"userNameRef,omitempty" tf:"-"`

	// Selector for a User in ram to populate userName.
	// +kubebuilder:validation:Optional
	UserNameSelector *v1.Selector `json:"userNameSelector,omitempty" tf:"-"`
}

// LoginProfileSpec defines the desired state of LoginProfile
type LoginProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoginProfileParameters `json:"forProvider"`
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
	InitProvider LoginProfileInitParameters `json:"initProvider,omitempty"`
}

// LoginProfileStatus defines the observed state of LoginProfile.
type LoginProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoginProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LoginProfile is the Schema for the LoginProfiles API. Provides a Alicloud RAM Login Profile resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type LoginProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.password) || (has(self.initProvider) && has(self.initProvider.password))",message="spec.forProvider.password is a required parameter"
	Spec   LoginProfileSpec   `json:"spec"`
	Status LoginProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoginProfileList contains a list of LoginProfiles
type LoginProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoginProfile `json:"items"`
}

// Repository type metadata.
var (
	LoginProfile_Kind             = "LoginProfile"
	LoginProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoginProfile_Kind}.String()
	LoginProfile_KindAPIVersion   = LoginProfile_Kind + "." + CRDGroupVersion.String()
	LoginProfile_GroupVersionKind = CRDGroupVersion.WithKind(LoginProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&LoginProfile{}, &LoginProfileList{})
}
