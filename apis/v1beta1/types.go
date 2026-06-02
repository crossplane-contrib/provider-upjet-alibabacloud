/*
Copyright 2022 Upbound Inc.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// OIDCOptions defines the options for OIDC/WebIdentity authentication
// using Alibaba Cloud RAM AssumeRoleWithOIDC.
type OIDCOptions struct {
	// RoleARN is the ARN of the RAM role to assume via OIDC.
	// Example: acs:ram::1234567890123456:role/example-oidc-role
	// +kubebuilder:validation:Required
	RoleARN string `json:"roleArn"`

	// ProviderARN is the ARN of the OIDC provider registered in Alibaba Cloud RAM.
	// Example: acs:ram::1234567890123456:oidc-provider/example-oidc-provider
	// +kubebuilder:validation:Required
	ProviderARN string `json:"providerArn"`

	// Region is the Alibaba Cloud region used for the STS AssumeRoleWithOIDC API call.
	// This is only required for the credential exchange and does not affect the region
	// of managed resources (which is taken from spec.forProvider.region).
	// If not set, defaults to cn-hangzhou.
	// +optional
	Region string `json:"region,omitempty"`

	// RoleSessionName is the identifier for the assumed role session.
	// Defaults to "crossplane-oidc-session".
	// +optional
	RoleSessionName string `json:"roleSessionName,omitempty"`
}

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	Credentials ProviderCredentials `json:"credentials"`
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=None;Secret;InjectedIdentity;Environment;Filesystem;WebIdentity
	Source xpv1.CredentialsSource `json:"source"`

	// OIDC defines the options for OIDC/WebIdentity-based authentication using
	// Alibaba Cloud RAM AssumeRoleWithOIDC. Required when source is WebIdentity.
	// +optional
	OIDC *OIDCOptions `json:"oidc,omitempty"`

	xpv1.CommonCredentialSelectors `json:",inline"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures a AlibabaCloud provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,template}
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// +kubebuilder:object:root=true

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CONFIG-NAME",type="string",JSONPath=".providerConfigRef.name"
// +kubebuilder:printcolumn:name="RESOURCE-KIND",type="string",JSONPath=".resourceRef.kind"
// +kubebuilder:printcolumn:name="RESOURCE-NAME",type="string",JSONPath=".resourceRef.name"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,template}
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv1.ProviderConfigUsage `json:",inline"`
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}
