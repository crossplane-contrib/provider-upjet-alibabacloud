// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	addresspool "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/addresspool"
	customline "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/customline"
	domain "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/domain"
	domainattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/domainattachment"
	domaingroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/domaingroup"
	gtminstance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/gtminstance"
	instance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/instance"
	monitorconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/monitorconfig"
	record "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/record"
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

// SetupGated_alidns creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_alidns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addresspool.SetupGated,
		customline.SetupGated,
		domain.SetupGated,
		domainattachment.SetupGated,
		domaingroup.SetupGated,
		gtminstance.SetupGated,
		instance.SetupGated,
		monitorconfig.SetupGated,
		record.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_alidns registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_alidns(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		addresspool.SetupWebhookWithManager,
		customline.SetupWebhookWithManager,
		domain.SetupWebhookWithManager,
		domainattachment.SetupWebhookWithManager,
		domaingroup.SetupWebhookWithManager,
		gtminstance.SetupWebhookWithManager,
		instance.SetupWebhookWithManager,
		monitorconfig.SetupWebhookWithManager,
		record.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
