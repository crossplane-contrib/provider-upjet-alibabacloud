// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/account"
	backup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/backup"
	backuppolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/backuppolicy"
	connection "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/connection"
	database "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/database"
	dbinstanceendpoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/dbinstanceendpoint"
	dbinstanceendpointaddress "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/dbinstanceendpointaddress"
	dbnode "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/dbnode"
	dbproxy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/dbproxy"
	dbproxypublic "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/dbproxypublic"
	instance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/instance"
	instancecrossbackuppolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/instancecrossbackuppolicy"
	parametergroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/parametergroup"
	readonlyinstance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/readonlyinstance"
	readwritesplittingconnection "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/readwritesplittingconnection"
	servicelinkedrole "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/servicelinkedrole"
	whitelisttemplate "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/rds/whitelisttemplate"
)

// Setup_rds creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_rds(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		backup.Setup,
		backuppolicy.Setup,
		connection.Setup,
		database.Setup,
		dbinstanceendpoint.Setup,
		dbinstanceendpointaddress.Setup,
		dbnode.Setup,
		dbproxy.Setup,
		dbproxypublic.Setup,
		instance.Setup,
		instancecrossbackuppolicy.Setup,
		parametergroup.Setup,
		readonlyinstance.Setup,
		readwritesplittingconnection.Setup,
		servicelinkedrole.Setup,
		whitelisttemplate.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
