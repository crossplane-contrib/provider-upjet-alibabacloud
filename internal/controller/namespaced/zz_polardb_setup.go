// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/polardb/account"
	accountprivilege "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/polardb/accountprivilege"
	backuppolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/polardb/backuppolicy"
	cluster "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/polardb/cluster"
	clusterendpoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/polardb/clusterendpoint"
	database "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/polardb/database"
	endpoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/polardb/endpoint"
	endpointaddress "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/polardb/endpointaddress"
	globaldatabasenetwork "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/polardb/globaldatabasenetwork"
	parametergroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/polardb/parametergroup"
	primaryendpoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/polardb/primaryendpoint"
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

// SetupGated_polardb creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_polardb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		accountprivilege.SetupGated,
		backuppolicy.SetupGated,
		cluster.SetupGated,
		clusterendpoint.SetupGated,
		database.SetupGated,
		endpoint.SetupGated,
		endpointaddress.SetupGated,
		globaldatabasenetwork.SetupGated,
		parametergroup.SetupGated,
		primaryendpoint.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_polardb registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_polardb(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		account.SetupWebhookWithManager,
		accountprivilege.SetupWebhookWithManager,
		backuppolicy.SetupWebhookWithManager,
		cluster.SetupWebhookWithManager,
		clusterendpoint.SetupWebhookWithManager,
		database.SetupWebhookWithManager,
		endpoint.SetupWebhookWithManager,
		endpointaddress.SetupWebhookWithManager,
		globaldatabasenetwork.SetupWebhookWithManager,
		parametergroup.SetupWebhookWithManager,
		primaryendpoint.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
