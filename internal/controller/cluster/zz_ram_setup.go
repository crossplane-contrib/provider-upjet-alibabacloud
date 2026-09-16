// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesskey "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/accesskey"
	accountalias "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/accountalias"
	accountpasswordpolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/accountpasswordpolicy"
	group "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/group"
	groupmembership "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/groupmembership"
	grouppolicyattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/grouppolicyattachment"
	loginprofile "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/loginprofile"
	passwordpolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/passwordpolicy"
	policy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/policy"
	role "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/role"
	rolepolicyattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/rolepolicyattachment"
	samlprovider "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/samlprovider"
	securitypreference "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/securitypreference"
	user "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/user"
	usergroupattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/usergroupattachment"
	userpolicyattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ram/userpolicyattachment"
)

// Setup_ram creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ram(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		group.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		loginprofile.Setup,
		passwordpolicy.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		securitypreference.Setup,
		user.Setup,
		usergroupattachment.Setup,
		userpolicyattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ram creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ram(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesskey.SetupGated,
		accountalias.SetupGated,
		accountpasswordpolicy.SetupGated,
		group.SetupGated,
		groupmembership.SetupGated,
		grouppolicyattachment.SetupGated,
		loginprofile.SetupGated,
		passwordpolicy.SetupGated,
		policy.SetupGated,
		role.SetupGated,
		rolepolicyattachment.SetupGated,
		samlprovider.SetupGated,
		securitypreference.SetupGated,
		user.SetupGated,
		usergroupattachment.SetupGated,
		userpolicyattachment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_ram registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_ram(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		accesskey.SetupWebhookWithManager,
		accountalias.SetupWebhookWithManager,
		accountpasswordpolicy.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		groupmembership.SetupWebhookWithManager,
		grouppolicyattachment.SetupWebhookWithManager,
		loginprofile.SetupWebhookWithManager,
		passwordpolicy.SetupWebhookWithManager,
		policy.SetupWebhookWithManager,
		role.SetupWebhookWithManager,
		rolepolicyattachment.SetupWebhookWithManager,
		samlprovider.SetupWebhookWithManager,
		securitypreference.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		usergroupattachment.SetupWebhookWithManager,
		userpolicyattachment.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
