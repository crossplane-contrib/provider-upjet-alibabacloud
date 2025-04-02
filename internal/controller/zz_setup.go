// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	domain "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cdn/domain"
	domainconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cdn/domainconfig"
	fctrigger "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cdn/fctrigger"
	group "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/group"
	endpoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/endpoint"
	endpointacl "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/endpointacl"
	queue "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/queue"
	subscription "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/subscription"
	topic "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/messageservice/topic"
	account "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/account"
	accountprivilege "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/accountprivilege"
	backuppolicy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/backuppolicy"
	cluster "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/cluster"
	clusterendpoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/clusterendpoint"
	database "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/database"
	endpointpolardb "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/endpoint"
	endpointaddress "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/endpointaddress"
	globaldatabasenetwork "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/globaldatabasenetwork"
	parametergroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/parametergroup"
	primaryendpoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/primaryendpoint"
	providerconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/providerconfig"
	vpc "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vpc"
	vswitch "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vswitch"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.Setup,
		domainconfig.Setup,
		fctrigger.Setup,
		group.Setup,
		endpoint.Setup,
		endpointacl.Setup,
		queue.Setup,
		subscription.Setup,
		topic.Setup,
		account.Setup,
		accountprivilege.Setup,
		backuppolicy.Setup,
		cluster.Setup,
		clusterendpoint.Setup,
		database.Setup,
		endpointpolardb.Setup,
		endpointaddress.Setup,
		globaldatabasenetwork.Setup,
		parametergroup.Setup,
		primaryendpoint.Setup,
		providerconfig.Setup,
		vpc.Setup,
		vswitch.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
