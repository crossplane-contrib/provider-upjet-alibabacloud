// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	application "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oos/application"
	applicationgroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oos/applicationgroup"
	defaultpatchbaseline "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oos/defaultpatchbaseline"
	execution "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oos/execution"
	parameter "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oos/parameter"
	patchbaseline "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oos/patchbaseline"
	secretparameter "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oos/secretparameter"
	servicesetting "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oos/servicesetting"
	stateconfiguration "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oos/stateconfiguration"
	template "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oos/template"
)

// Setup_oos creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_oos(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		applicationgroup.Setup,
		defaultpatchbaseline.Setup,
		execution.Setup,
		parameter.Setup,
		patchbaseline.Setup,
		secretparameter.Setup,
		servicesetting.Setup,
		stateconfiguration.Setup,
		template.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_oos creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_oos(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.SetupGated,
		applicationgroup.SetupGated,
		defaultpatchbaseline.SetupGated,
		execution.SetupGated,
		parameter.SetupGated,
		patchbaseline.SetupGated,
		secretparameter.SetupGated,
		servicesetting.SetupGated,
		stateconfiguration.SetupGated,
		template.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_oos registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_oos(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		application.SetupWebhookWithManager,
		applicationgroup.SetupWebhookWithManager,
		defaultpatchbaseline.SetupWebhookWithManager,
		execution.SetupWebhookWithManager,
		parameter.SetupWebhookWithManager,
		patchbaseline.SetupWebhookWithManager,
		secretparameter.SetupWebhookWithManager,
		servicesetting.SetupWebhookWithManager,
		stateconfiguration.SetupWebhookWithManager,
		template.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
