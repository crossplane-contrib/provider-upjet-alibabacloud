// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	activation "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/activation"
	autoprovisioninggroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/autoprovisioninggroup"
	autosnapshotpolicy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/autosnapshotpolicy"
	autosnapshotpolicyattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/autosnapshotpolicyattachment"
	capacityreservation "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/capacityreservation"
	command "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/command"
	dedicatedhost "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/dedicatedhost"
	dedicatedhostcluster "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/dedicatedhostcluster"
	deploymentset "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/deploymentset"
	disk "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/disk"
	diskattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/diskattachment"
	elasticityassurance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/elasticityassurance"
	hpccluster "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/hpccluster"
	image "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/image"
	imagecomponent "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imagecomponent"
	imagecopy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imagecopy"
	imageexport "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imageexport"
	imageimport "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imageimport"
	imagepipeline "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imagepipeline"
	imagepipelineexecution "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imagepipelineexecution"
	imagesharepermission "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/imagesharepermission"
	instance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/instance"
	instanceset "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/instanceset"
	invocation "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/invocation"
	keypair "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/keypair"
	keypairattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/keypairattachment"
	launchtemplate "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/launchtemplate"
	networkinterface "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/networkinterface"
	networkinterfaceattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/networkinterfaceattachment"
	networkinterfacepermission "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/networkinterfacepermission"
	prefixlist "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/prefixlist"
	ramroleattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/ramroleattachment"
	reservedinstance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/reservedinstance"
	securitygroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/securitygrouprule"
	sessionmanagerstatus "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/sessionmanagerstatus"
	snapshot "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/snapshot"
	snapshotgroup "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/snapshotgroup"
	snapshotpolicy "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/snapshotpolicy"
	storagecapacityunit "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/storagecapacityunit"
	alias "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/alias"
	applicationaccesspoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/applicationaccesspoint"
	ciphertext "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/ciphertext"
	clientkey "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/clientkey"
	instancekms "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/instance"
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
	group "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ram/group"
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
	dhcpoptionsset "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/dhcpoptionsset"
	dhcpoptionssetattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/dhcpoptionssetattachment"
	flowlog "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/flowlog"
	gatewayendpoint "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/gatewayendpoint"
	gatewayendpointroutetableattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/gatewayendpointroutetableattachment"
	gatewayroutetableattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/gatewayroutetableattachment"
	havip "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/havip"
	ipv4cidrblock "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/ipv4cidrblock"
	ipv4gateway "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/ipv4gateway"
	ipv6address "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/ipv6address"
	ipv6egressrule "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/ipv6egressrule"
	ipv6gateway "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/ipv6gateway"
	ipv6internetbandwidth "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/ipv6internetbandwidth"
	networkaclattachment "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/networkaclattachment"
	peerconnection "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/peerconnection"
	peerconnectionaccepter "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/peerconnectionaccepter"
	prefixlistvpc "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/prefixlist"
	publicipaddresspool "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/publicipaddresspool"
	publicipaddresspoolcidrblock "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/publicipaddresspoolcidrblock"
	trafficmirrorfilter "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/trafficmirrorfilter"
	trafficmirrorfilteregressrule "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/trafficmirrorfilteregressrule"
	trafficmirrorfilteringressrule "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/trafficmirrorfilteringressrule"
	trafficmirrorsession "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/trafficmirrorsession"
	vpc "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vpc"
	vswitch "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vswitch"
	vswitchcidrreservation "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vswitchcidrreservation"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activation.Setup,
		autoprovisioninggroup.Setup,
		autosnapshotpolicy.Setup,
		autosnapshotpolicyattachment.Setup,
		capacityreservation.Setup,
		command.Setup,
		dedicatedhost.Setup,
		dedicatedhostcluster.Setup,
		deploymentset.Setup,
		disk.Setup,
		disk.Setup,
		diskattachment.Setup,
		diskattachment.Setup,
		elasticityassurance.Setup,
		hpccluster.Setup,
		image.Setup,
		imagecomponent.Setup,
		imagecopy.Setup,
		imageexport.Setup,
		imageimport.Setup,
		imagepipeline.Setup,
		imagepipelineexecution.Setup,
		imagesharepermission.Setup,
		instance.Setup,
		instanceset.Setup,
		invocation.Setup,
		keypair.Setup,
		keypair.Setup,
		keypairattachment.Setup,
		keypairattachment.Setup,
		launchtemplate.Setup,
		launchtemplate.Setup,
		networkinterface.Setup,
		networkinterface.Setup,
		networkinterfaceattachment.Setup,
		networkinterfaceattachment.Setup,
		networkinterfacepermission.Setup,
		prefixlist.Setup,
		ramroleattachment.Setup,
		reservedinstance.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		sessionmanagerstatus.Setup,
		snapshot.Setup,
		snapshot.Setup,
		snapshotgroup.Setup,
		snapshotpolicy.Setup,
		storagecapacityunit.Setup,
		alias.Setup,
		applicationaccesspoint.Setup,
		ciphertext.Setup,
		clientkey.Setup,
		instancekms.Setup,
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
		group.Setup,
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
		dhcpoptionsset.Setup,
		dhcpoptionssetattachment.Setup,
		flowlog.Setup,
		gatewayendpoint.Setup,
		gatewayendpointroutetableattachment.Setup,
		gatewayroutetableattachment.Setup,
		havip.Setup,
		ipv4cidrblock.Setup,
		ipv4gateway.Setup,
		ipv6address.Setup,
		ipv6egressrule.Setup,
		ipv6gateway.Setup,
		ipv6internetbandwidth.Setup,
		networkaclattachment.Setup,
		peerconnection.Setup,
		peerconnectionaccepter.Setup,
		prefixlistvpc.Setup,
		publicipaddresspool.Setup,
		publicipaddresspoolcidrblock.Setup,
		trafficmirrorfilter.Setup,
		trafficmirrorfilteregressrule.Setup,
		trafficmirrorfilteringressrule.Setup,
		trafficmirrorsession.Setup,
		vpc.Setup,
		vswitch.Setup,
		vswitchcidrreservation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
