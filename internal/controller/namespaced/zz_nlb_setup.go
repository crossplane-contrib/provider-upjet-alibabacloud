// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	loadbalancer "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/nlb/loadbalancer"
)

// Setup_nlb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_nlb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		loadbalancer.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_nlb creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_nlb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		loadbalancer.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_nlb registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_nlb(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		loadbalancer.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
