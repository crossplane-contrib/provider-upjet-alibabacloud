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
	accesspoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/accesspoint"
	accountpublicaccessblock "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/accountpublicaccessblock"
	bucket "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucket"
	bucketaccessmonitor "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketaccessmonitor"
	bucketacl "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketacl"
	bucketcname "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketcname"
	bucketcnametoken "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketcnametoken"
	bucketcors "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketcors"
	bucketdataredundancytransition "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketdataredundancytransition"
	buckethttpsconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/buckethttpsconfig"
	bucketlogging "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketlogging"
	bucketmetaquery "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketmetaquery"
	bucketobject "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketobject"
	bucketpolicy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketpolicy"
	bucketpublicaccessblock "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketpublicaccessblock"
	bucketreferer "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketreferer"
	bucketreplication "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketreplication"
	bucketrequestpayment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketrequestpayment"
	bucketserversideencryption "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketserversideencryption"
	buckettransferacceleration "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/buckettransferacceleration"
	bucketuserdefinedlogfields "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketuserdefinedlogfields"
	bucketversioning "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketversioning"
	bucketwebsite "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketwebsite"
	bucketworm "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketworm"
	providerconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/providerconfig"
	quotaalarm "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/quotas/quotaalarm"
	quotaapplication "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/quotas/quotaapplication"
	templateapplications "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/quotas/templateapplications"
	templatequota "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/quotas/templatequota"
	templateservice "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/quotas/templateservice"
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
	vswitch "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vswitch"
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
		accesspoint.Setup,
		accountpublicaccessblock.Setup,
		bucket.Setup,
		bucketaccessmonitor.Setup,
		bucketacl.Setup,
		bucketcname.Setup,
		bucketcnametoken.Setup,
		bucketcors.Setup,
		bucketdataredundancytransition.Setup,
		buckethttpsconfig.Setup,
		bucketlogging.Setup,
		bucketmetaquery.Setup,
		bucketobject.Setup,
		bucketpolicy.Setup,
		bucketpublicaccessblock.Setup,
		bucketreferer.Setup,
		bucketreplication.Setup,
		bucketrequestpayment.Setup,
		bucketserversideencryption.Setup,
		buckettransferacceleration.Setup,
		bucketuserdefinedlogfields.Setup,
		bucketversioning.Setup,
		bucketwebsite.Setup,
		bucketworm.Setup,
		providerconfig.Setup,
		quotaalarm.Setup,
		quotaapplication.Setup,
		templateapplications.Setup,
		templatequota.Setup,
		templateservice.Setup,
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
		vswitch.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
