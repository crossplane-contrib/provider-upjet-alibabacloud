// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	quotaalarm "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/quotas/quotaalarm"
	quotaapplication "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/quotas/quotaapplication"
	templateapplications "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/quotas/templateapplications"
	templatequota "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/quotas/templatequota"
	templateservice "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/quotas/templateservice"
)

// Setup_quotas creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_quotas(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		quotaalarm.Setup,
		quotaapplication.Setup,
		templateapplications.Setup,
		templatequota.Setup,
		templateservice.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_quotas creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_quotas(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		quotaalarm.SetupGated,
		quotaapplication.SetupGated,
		templateapplications.SetupGated,
		templatequota.SetupGated,
		templateservice.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_quotas registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_quotas(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		quotaalarm.SetupWebhookWithManager,
		quotaapplication.SetupWebhookWithManager,
		templateapplications.SetupWebhookWithManager,
		templatequota.SetupWebhookWithManager,
		templateservice.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
