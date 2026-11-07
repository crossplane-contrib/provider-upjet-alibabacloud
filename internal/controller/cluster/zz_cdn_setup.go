// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	domain "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cdn/domain"
	domainconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cdn/domainconfig"
	fctrigger "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cdn/fctrigger"
)

// Setup_cdn creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cdn(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.Setup,
		domainconfig.Setup,
		fctrigger.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cdn creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cdn(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.SetupGated,
		domainconfig.SetupGated,
		fctrigger.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_cdn registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_cdn(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		domain.SetupWebhookWithManager,
		domainconfig.SetupWebhookWithManager,
		fctrigger.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
