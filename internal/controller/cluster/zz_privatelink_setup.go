// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	vpcendpoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpoint"
	vpcendpointconnection "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpointconnection"
	vpcendpointservice "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpointservice"
	vpcendpointserviceresource "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpointserviceresource"
	vpcendpointserviceuser "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpointserviceuser"
	vpcendpointzone "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpointzone"
)

// Setup_privatelink creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_privatelink(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		vpcendpoint.Setup,
		vpcendpointconnection.Setup,
		vpcendpointservice.Setup,
		vpcendpointserviceresource.Setup,
		vpcendpointserviceuser.Setup,
		vpcendpointzone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_privatelink creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_privatelink(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		vpcendpoint.SetupGated,
		vpcendpointconnection.SetupGated,
		vpcendpointservice.SetupGated,
		vpcendpointserviceresource.SetupGated,
		vpcendpointserviceuser.SetupGated,
		vpcendpointzone.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_privatelink registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_privatelink(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		vpcendpoint.SetupWebhookWithManager,
		vpcendpointconnection.SetupWebhookWithManager,
		vpcendpointservice.SetupWebhookWithManager,
		vpcendpointserviceresource.SetupWebhookWithManager,
		vpcendpointserviceuser.SetupWebhookWithManager,
		vpcendpointzone.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
