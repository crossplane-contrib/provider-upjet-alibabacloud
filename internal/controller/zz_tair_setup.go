// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/tair/account"
	auditlogconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/tair/auditlogconfig"
	connection "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/tair/connection"
	instance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/tair/instance"
	tairinstance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/tair/tairinstance"
)

// Setup_tair creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_tair(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		auditlogconfig.Setup,
		connection.Setup,
		instance.Setup,
		tairinstance.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_tair creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_tair(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		auditlogconfig.SetupGated,
		connection.SetupGated,
		instance.SetupGated,
		tairinstance.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_tair registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_tair(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		account.SetupWebhookWithManager,
		auditlogconfig.SetupWebhookWithManager,
		connection.SetupWebhookWithManager,
		instance.SetupWebhookWithManager,
		tairinstance.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
