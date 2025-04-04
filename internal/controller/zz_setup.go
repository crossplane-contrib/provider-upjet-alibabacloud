// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	addresspool "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/addresspool"
	customline "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/customline"
	domain "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/domain"
	domainattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/domainattachment"
	domaingroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/domaingroup"
	gtminstance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/gtminstance"
	instance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/instance"
	monitorconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/monitorconfig"
	record "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/alidns/record"
	domaincdn "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cdn/domain"
	domainconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cdn/domainconfig"
	fctrigger "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cdn/fctrigger"
	alarmcontactgroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/cloudmonitorservice/alarmcontactgroup"
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
	accounttair "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/account"
	auditlogconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/auditlogconfig"
	connection "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/connection"
	instancetair "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/instance"
	tairinstance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/tairinstance"
	vpc "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vpc"
	vswitch "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vswitch"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addresspool.Setup,
		customline.Setup,
		domain.Setup,
		domainattachment.Setup,
		domaingroup.Setup,
		gtminstance.Setup,
		instance.Setup,
		monitorconfig.Setup,
		record.Setup,
		domaincdn.Setup,
		domainconfig.Setup,
		fctrigger.Setup,
		alarmcontactgroup.Setup,
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
		accounttair.Setup,
		auditlogconfig.Setup,
		connection.Setup,
		instancetair.Setup,
		tairinstance.Setup,
		vpc.Setup,
		vswitch.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
