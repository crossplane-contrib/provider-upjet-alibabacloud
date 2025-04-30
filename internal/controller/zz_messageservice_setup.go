// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	endpoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/endpoint"
	endpointacl "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/endpointacl"
	queue "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/queue"
	subscription "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/subscription"
	topic "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/topic"
)

// Setup_messageservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_messageservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpoint.Setup,
		endpointacl.Setup,
		queue.Setup,
		subscription.Setup,
		topic.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
