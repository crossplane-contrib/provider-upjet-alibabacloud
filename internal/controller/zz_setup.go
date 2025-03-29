// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	group "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/group"
	alias "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/alias"
	applicationaccesspoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/applicationaccesspoint"
	ciphertext "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/ciphertext"
	clientkey "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/clientkey"
	instance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/instance"
	key "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/key"
	keyversion "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/keyversion"
	networkrule "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/networkrule"
	policy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/policy"
	secret "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/secret"
	providerconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/providerconfig"
	accesskey "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/accesskey"
	accountalias "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/accountalias"
	accountpasswordpolicy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/accountpasswordpolicy"
	groupram "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/group"
	groupmembership "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/groupmembership"
	grouppolicyattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/grouppolicyattachment"
	loginprofile "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/loginprofile"
	policyram "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/policy"
	role "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/role"
	rolepolicyattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/rolepolicyattachment"
	samlprovider "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/samlprovider"
	securitypreference "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/securitypreference"
	user "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/user"
	userpolicyattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/userpolicyattachment"
	vpc "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.Setup,
		alias.Setup,
		applicationaccesspoint.Setup,
		ciphertext.Setup,
		clientkey.Setup,
		instance.Setup,
		key.Setup,
		keyversion.Setup,
		networkrule.Setup,
		policy.Setup,
		secret.Setup,
		providerconfig.Setup,
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		groupram.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		loginprofile.Setup,
		policyram.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		securitypreference.Setup,
		user.Setup,
		userpolicyattachment.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
