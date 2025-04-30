// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/account"
	accountprivilege "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/accountprivilege"
	backuppolicy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/backuppolicy"
	cluster "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/cluster"
	clusterendpoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/clusterendpoint"
	database "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/database"
	endpoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/endpoint"
	endpointaddress "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/endpointaddress"
	globaldatabasenetwork "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/globaldatabasenetwork"
	parametergroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/parametergroup"
	primaryendpoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/polardb/primaryendpoint"
)

// Setup_polardb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_polardb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		accountprivilege.Setup,
		backuppolicy.Setup,
		cluster.Setup,
		clusterendpoint.Setup,
		database.Setup,
		endpoint.Setup,
		endpointaddress.Setup,
		globaldatabasenetwork.Setup,
		parametergroup.Setup,
		primaryendpoint.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
