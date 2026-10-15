// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/account"
	backup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/backup"
	backuppolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/backuppolicy"
	connection "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/connection"
	database "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/database"
	dbinstanceendpoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/dbinstanceendpoint"
	dbinstanceendpointaddress "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/dbinstanceendpointaddress"
	dbnode "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/dbnode"
	dbproxy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/dbproxy"
	dbproxypublic "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/dbproxypublic"
	instance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/instance"
	instancecrossbackuppolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/instancecrossbackuppolicy"
	parametergroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/parametergroup"
	readonlyinstance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/readonlyinstance"
	readwritesplittingconnection "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/readwritesplittingconnection"
	servicelinkedrole "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/servicelinkedrole"
	whitelisttemplate "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/rds/whitelisttemplate"
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

// SetupGated_rds creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_rds(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		backup.SetupGated,
		backuppolicy.SetupGated,
		connection.SetupGated,
		database.SetupGated,
		dbinstanceendpoint.SetupGated,
		dbinstanceendpointaddress.SetupGated,
		dbnode.SetupGated,
		dbproxy.SetupGated,
		dbproxypublic.SetupGated,
		instance.SetupGated,
		instancecrossbackuppolicy.SetupGated,
		parametergroup.SetupGated,
		readonlyinstance.SetupGated,
		readwritesplittingconnection.SetupGated,
		servicelinkedrole.SetupGated,
		whitelisttemplate.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_rds registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_rds(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		account.SetupWebhookWithManager,
		backup.SetupWebhookWithManager,
		backuppolicy.SetupWebhookWithManager,
		connection.SetupWebhookWithManager,
		database.SetupWebhookWithManager,
		dbinstanceendpoint.SetupWebhookWithManager,
		dbinstanceendpointaddress.SetupWebhookWithManager,
		dbnode.SetupWebhookWithManager,
		dbproxy.SetupWebhookWithManager,
		dbproxypublic.SetupWebhookWithManager,
		instance.SetupWebhookWithManager,
		instancecrossbackuppolicy.SetupWebhookWithManager,
		parametergroup.SetupWebhookWithManager,
		readonlyinstance.SetupWebhookWithManager,
		readwritesplittingconnection.SetupWebhookWithManager,
		servicelinkedrole.SetupWebhookWithManager,
		whitelisttemplate.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
