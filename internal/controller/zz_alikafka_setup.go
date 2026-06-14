// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	consumergroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alikafka/consumergroup"
	instance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alikafka/instance"
	instanceallowedipattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alikafka/instanceallowedipattachment"
	saslacl "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alikafka/saslacl"
	sasluser "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alikafka/sasluser"
	scheduledscalingrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alikafka/scheduledscalingrule"
	topic "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alikafka/topic"
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
