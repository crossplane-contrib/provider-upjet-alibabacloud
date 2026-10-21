// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	chain "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/chain"
	chartnamespace "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/chartnamespace"
	chartrepository "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/chartrepository"
	eeinstance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/eeinstance"
	eenamespace "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/eenamespace"
	eerepo "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/eerepo"
	eesyncrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/eesyncrule"
	endpointaclpolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/endpointaclpolicy"
	registrynamespace "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/registrynamespace"
	repo "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/repo"
	scanrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/scanrule"
	storagedomainroutingrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/storagedomainroutingrule"
	vpcendpointlinkedvpc "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/vpcendpointlinkedvpc"
)

// Setup_cr creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cr(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		chain.Setup,
		chartnamespace.Setup,
		chartrepository.Setup,
		eeinstance.Setup,
		eenamespace.Setup,
		eerepo.Setup,
		eesyncrule.Setup,
		endpointaclpolicy.Setup,
		registrynamespace.Setup,
		repo.Setup,
		scanrule.Setup,
		storagedomainroutingrule.Setup,
		vpcendpointlinkedvpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cr creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cr(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		chain.SetupGated,
		chartnamespace.SetupGated,
		chartrepository.SetupGated,
		eeinstance.SetupGated,
		eenamespace.SetupGated,
		eerepo.SetupGated,
		eesyncrule.SetupGated,
		endpointaclpolicy.SetupGated,
		registrynamespace.SetupGated,
		repo.SetupGated,
		scanrule.SetupGated,
		storagedomainroutingrule.SetupGated,
		vpcendpointlinkedvpc.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_cr registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_cr(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		chain.SetupWebhookWithManager,
		chartnamespace.SetupWebhookWithManager,
		chartrepository.SetupWebhookWithManager,
		eeinstance.SetupWebhookWithManager,
		eenamespace.SetupWebhookWithManager,
		eerepo.SetupWebhookWithManager,
		eesyncrule.SetupWebhookWithManager,
		endpointaclpolicy.SetupWebhookWithManager,
		registrynamespace.SetupWebhookWithManager,
		repo.SetupWebhookWithManager,
		scanrule.SetupWebhookWithManager,
		storagedomainroutingrule.SetupWebhookWithManager,
		vpcendpointlinkedvpc.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
