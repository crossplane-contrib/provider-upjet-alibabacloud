// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"

	// PathIdExtractor is the golang path to
	// IdExtractor function in this package.
	PathIdExtractor                    = SelfPackagePath + ".IdExtractor()"
	PathAccountNameExtractor           = SelfPackagePath + ".AccountNameExtractor()"
	PathDBNameExtractor                = SelfPackagePath + ".DBNameExtractor()"
	PathDBEndpointIdExtractor          = SelfPackagePath + ".DBEndpointIdExtractor()"
	PathAlarmContactGroupNameExtractor = SelfPackagePath + ".AlarmContactGroupNameExtractor()"
)

// IdExtractor extracts id of the
// resources from "status.atProvider.id".
func IdExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.id")
		if err != nil {
			return ""
		}
		return r
	}
}

// AccountNameExtractor extracts id of the
// resources from "status.atProvider.account_name".
func AccountNameExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.accountName")
		if err != nil {
			return ""
		}
		return r
	}
}

// DBNameExtractor extracts id of the
// resources from "status.atProvider.account_name".
func DBNameExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.dbName")
		if err != nil {
			return ""
		}
		return r
	}
}

// DBEndpointIdExtractor extracts id of the
// resources from "status.atProvider.account_name".
func DBEndpointIdExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.dbEndpointId")
		if err != nil {
			return ""
		}
		return r
	}
}

// AlarmContactGroupNameExtractor extracts id of the
// resources from "status.atProvider.account_name".
func AlarmContactGroupNameExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.alarmContactGroupNameExtractor")
		if err != nil {
			return ""
		}
		return r
	}
}
