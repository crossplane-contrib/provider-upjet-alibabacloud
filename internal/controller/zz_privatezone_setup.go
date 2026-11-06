// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	endpoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/privatezone/endpoint"
	rule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/privatezone/rule"
	ruleattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/privatezone/ruleattachment"
	uservpcauthorization "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/privatezone/uservpcauthorization"
	zone "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/privatezone/zone"
	zoneattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/privatezone/zoneattachment"
	zonerecord "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/privatezone/zonerecord"
)

// Setup_privatezone creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_privatezone(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpoint.Setup,
		rule.Setup,
		ruleattachment.Setup,
		uservpcauthorization.Setup,
		zone.Setup,
		zoneattachment.Setup,
		zonerecord.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
