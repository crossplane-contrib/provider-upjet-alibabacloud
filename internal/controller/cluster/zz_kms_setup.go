// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	alias "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/kms/alias"
	instance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/kms/instance"
	key "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/kms/key"
	secret "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/kms/secret"
)

// Setup_kms creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.Setup,
		instance.Setup,
		key.Setup,
		secret.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_kms creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_kms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.SetupGated,
		instance.SetupGated,
		key.SetupGated,
		secret.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_kms registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_kms(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		alias.SetupWebhookWithManager,
		instance.SetupWebhookWithManager,
		key.SetupWebhookWithManager,
		secret.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
