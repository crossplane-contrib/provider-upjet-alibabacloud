// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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
	bucketstyle "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketstyle"
	buckettransferacceleration "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/buckettransferacceleration"
	bucketuserdefinedlogfields "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketuserdefinedlogfields"
	bucketversioning "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketversioning"
	bucketwebsite "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketwebsite"
	bucketworm "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/oss/bucketworm"
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
