// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesspoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/accesspoint"
	accountpublicaccessblock "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/accountpublicaccessblock"
	bucket "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucket"
	bucketaccessmonitor "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketaccessmonitor"
	bucketacl "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketacl"
	bucketcname "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketcname"
	bucketcnametoken "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketcnametoken"
	bucketcors "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketcors"
	bucketdataredundancytransition "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketdataredundancytransition"
	buckethttpsconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/buckethttpsconfig"
	bucketlogging "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketlogging"
	bucketmetaquery "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketmetaquery"
	bucketobject "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketobject"
	bucketpolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketpolicy"
	bucketpublicaccessblock "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketpublicaccessblock"
	bucketreferer "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketreferer"
	bucketreplication "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketreplication"
	bucketrequestpayment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketrequestpayment"
	bucketserversideencryption "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketserversideencryption"
	bucketstyle "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketstyle"
	buckettransferacceleration "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/buckettransferacceleration"
	bucketuserdefinedlogfields "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketuserdefinedlogfields"
	bucketversioning "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketversioning"
	bucketwebsite "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketwebsite"
	bucketworm "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/oss/bucketworm"
)

// Setup_oss creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_oss(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
		bucketstyle.Setup,
		buckettransferacceleration.Setup,
		bucketuserdefinedlogfields.Setup,
		bucketversioning.Setup,
		bucketwebsite.Setup,
		bucketworm.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_oss creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_oss(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesspoint.SetupGated,
		accountpublicaccessblock.SetupGated,
		bucket.SetupGated,
		bucketaccessmonitor.SetupGated,
		bucketacl.SetupGated,
		bucketcname.SetupGated,
		bucketcnametoken.SetupGated,
		bucketcors.SetupGated,
		bucketdataredundancytransition.SetupGated,
		buckethttpsconfig.SetupGated,
		bucketlogging.SetupGated,
		bucketmetaquery.SetupGated,
		bucketobject.SetupGated,
		bucketpolicy.SetupGated,
		bucketpublicaccessblock.SetupGated,
		bucketreferer.SetupGated,
		bucketreplication.SetupGated,
		bucketrequestpayment.SetupGated,
		bucketserversideencryption.SetupGated,
		bucketstyle.SetupGated,
		buckettransferacceleration.SetupGated,
		bucketuserdefinedlogfields.SetupGated,
		bucketversioning.SetupGated,
		bucketwebsite.SetupGated,
		bucketworm.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_oss registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_oss(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		accesspoint.SetupWebhookWithManager,
		accountpublicaccessblock.SetupWebhookWithManager,
		bucket.SetupWebhookWithManager,
		bucketaccessmonitor.SetupWebhookWithManager,
		bucketacl.SetupWebhookWithManager,
		bucketcname.SetupWebhookWithManager,
		bucketcnametoken.SetupWebhookWithManager,
		bucketcors.SetupWebhookWithManager,
		bucketdataredundancytransition.SetupWebhookWithManager,
		buckethttpsconfig.SetupWebhookWithManager,
		bucketlogging.SetupWebhookWithManager,
		bucketmetaquery.SetupWebhookWithManager,
		bucketobject.SetupWebhookWithManager,
		bucketpolicy.SetupWebhookWithManager,
		bucketpublicaccessblock.SetupWebhookWithManager,
		bucketreferer.SetupWebhookWithManager,
		bucketreplication.SetupWebhookWithManager,
		bucketrequestpayment.SetupWebhookWithManager,
		bucketserversideencryption.SetupWebhookWithManager,
		bucketstyle.SetupWebhookWithManager,
		buckettransferacceleration.SetupWebhookWithManager,
		bucketuserdefinedlogfields.SetupWebhookWithManager,
		bucketversioning.SetupWebhookWithManager,
		bucketwebsite.SetupWebhookWithManager,
		bucketworm.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
