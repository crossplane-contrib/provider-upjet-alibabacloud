// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	logtailattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/sls/logtailattachment"
	logtailconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/sls/logtailconfig"
	machinegroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/sls/machinegroup"
	project "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/sls/project"
	store "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/sls/store"
	storeindex "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/sls/storeindex"
)

// Setup_sls creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sls(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		logtailattachment.Setup,
		logtailconfig.Setup,
		machinegroup.Setup,
		project.Setup,
		store.Setup,
		storeindex.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_sls creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_sls(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		logtailattachment.SetupGated,
		logtailconfig.SetupGated,
		machinegroup.SetupGated,
		project.SetupGated,
		store.SetupGated,
		storeindex.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_sls registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_sls(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		logtailattachment.SetupWebhookWithManager,
		logtailconfig.SetupWebhookWithManager,
		machinegroup.SetupWebhookWithManager,
		project.SetupWebhookWithManager,
		store.SetupWebhookWithManager,
		storeindex.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
