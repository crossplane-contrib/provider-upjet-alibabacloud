// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cluster "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ackone/cluster"
	membershipattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ackone/membershipattachment"
)

// Setup_ackone creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ackone(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		membershipattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ackone creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ackone(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.SetupGated,
		membershipattachment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_ackone registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_ackone(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		cluster.SetupWebhookWithManager,
		membershipattachment.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
