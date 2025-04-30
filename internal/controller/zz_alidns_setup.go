// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	addresspool "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/addresspool"
	customline "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/customline"
	domain "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/domain"
	domainattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/domainattachment"
	domaingroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/domaingroup"
	gtminstance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/gtminstance"
	instance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/instance"
	monitorconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/monitorconfig"
	record "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/record"
)

// Setup_alidns creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_alidns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addresspool.Setup,
		customline.Setup,
		domain.Setup,
		domainattachment.Setup,
		domaingroup.Setup,
		gtminstance.Setup,
		instance.Setup,
		monitorconfig.Setup,
		record.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
