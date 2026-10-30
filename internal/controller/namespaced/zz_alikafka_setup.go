// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	consumergroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alikafka/consumergroup"
	instance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alikafka/instance"
	instanceallowedipattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alikafka/instanceallowedipattachment"
	saslacl "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alikafka/saslacl"
	sasluser "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alikafka/sasluser"
	scheduledscalingrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alikafka/scheduledscalingrule"
	topic "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alikafka/topic"
)

// Setup_alikafka creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_alikafka(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		consumergroup.Setup,
		instance.Setup,
		instanceallowedipattachment.Setup,
		saslacl.Setup,
		sasluser.Setup,
		scheduledscalingrule.Setup,
		topic.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_alikafka creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_alikafka(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		consumergroup.SetupGated,
		instance.SetupGated,
		instanceallowedipattachment.SetupGated,
		saslacl.SetupGated,
		sasluser.SetupGated,
		scheduledscalingrule.SetupGated,
		topic.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_alikafka registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_alikafka(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		consumergroup.SetupWebhookWithManager,
		instance.SetupWebhookWithManager,
		instanceallowedipattachment.SetupWebhookWithManager,
		saslacl.SetupWebhookWithManager,
		sasluser.SetupWebhookWithManager,
		scheduledscalingrule.SetupWebhookWithManager,
		topic.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
