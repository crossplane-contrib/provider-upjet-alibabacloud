/*
Copyright 2021 Upbound Inc.
*/

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "alibabacloud.crossplane.io"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	// This group version holds no types: the StoreConfig API it used to serve
	// was part of the alpha External Secret Store feature, which Crossplane v2
	// and upjet v2 have dropped. The package is retained because upjet's
	// DefaultBasePackages registers "v1alpha1" as a base API version.
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)
