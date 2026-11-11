// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	acl "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/slb/acl"
	listener "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/slb/listener"
	loadbalancer "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/slb/loadbalancer"
)

// Setup_slb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_slb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.Setup,
		listener.Setup,
		loadbalancer.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_slb creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_slb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.SetupGated,
		listener.SetupGated,
		loadbalancer.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_slb registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_slb(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		acl.SetupWebhookWithManager,
		listener.SetupWebhookWithManager,
		loadbalancer.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
