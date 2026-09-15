// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	activation "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/activation"
	autoprovisioninggroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/autoprovisioninggroup"
	autosnapshotpolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/autosnapshotpolicy"
	autosnapshotpolicyattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/autosnapshotpolicyattachment"
	capacityreservation "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/capacityreservation"
	command "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/command"
	dedicatedhost "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/dedicatedhost"
	dedicatedhostcluster "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/dedicatedhostcluster"
	deploymentset "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/deploymentset"
	disk "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/disk"
	diskattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/diskattachment"
	elasticityassurance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/elasticityassurance"
	hpccluster "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/hpccluster"
	image "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/image"
	imagecomponent "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/imagecomponent"
	imagecopy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/imagecopy"
	imageexport "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/imageexport"
	imageimport "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/imageimport"
	imagepipeline "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/imagepipeline"
	imagepipelineexecution "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/imagepipelineexecution"
	imagesharepermission "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/imagesharepermission"
	instance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/instance"
	instanceset "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/instanceset"
	invocation "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/invocation"
	keypair "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/keypair"
	keypairattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/keypairattachment"
	launchtemplate "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/launchtemplate"
	networkinterface "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/networkinterface"
	networkinterfaceattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/networkinterfaceattachment"
	networkinterfacepermission "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/networkinterfacepermission"
	prefixlist "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/prefixlist"
	reservedinstance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/reservedinstance"
	securitygroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/securitygrouprule"
	sessionmanagerstatus "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/sessionmanagerstatus"
	snapshot "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/snapshot"
	snapshotgroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/snapshotgroup"
	storagecapacityunit "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/ecs/storagecapacityunit"
)

// Setup_ecs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ecs(mgr ctrl.Manager, o controller.Options) error {
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
		keypairattachment.Setup,
		launchtemplate.Setup,
		networkinterface.Setup,
		networkinterfaceattachment.Setup,
		networkinterfacepermission.Setup,
		prefixlist.Setup,
		reservedinstance.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		sessionmanagerstatus.Setup,
		snapshot.Setup,
		snapshotgroup.Setup,
		storagecapacityunit.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ecs creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ecs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activation.SetupGated,
		autoprovisioninggroup.SetupGated,
		autosnapshotpolicy.SetupGated,
		autosnapshotpolicyattachment.SetupGated,
		capacityreservation.SetupGated,
		command.SetupGated,
		dedicatedhost.SetupGated,
		dedicatedhostcluster.SetupGated,
		deploymentset.SetupGated,
		disk.SetupGated,
		diskattachment.SetupGated,
		elasticityassurance.SetupGated,
		hpccluster.SetupGated,
		image.SetupGated,
		imagecomponent.SetupGated,
		imagecopy.SetupGated,
		imageexport.SetupGated,
		imageimport.SetupGated,
		imagepipeline.SetupGated,
		imagepipelineexecution.SetupGated,
		imagesharepermission.SetupGated,
		instance.SetupGated,
		instanceset.SetupGated,
		invocation.SetupGated,
		keypair.SetupGated,
		keypairattachment.SetupGated,
		launchtemplate.SetupGated,
		networkinterface.SetupGated,
		networkinterfaceattachment.SetupGated,
		networkinterfacepermission.SetupGated,
		prefixlist.SetupGated,
		reservedinstance.SetupGated,
		securitygroup.SetupGated,
		securitygrouprule.SetupGated,
		sessionmanagerstatus.SetupGated,
		snapshot.SetupGated,
		snapshotgroup.SetupGated,
		storagecapacityunit.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_ecs registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_ecs(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		activation.SetupWebhookWithManager,
		autoprovisioninggroup.SetupWebhookWithManager,
		autosnapshotpolicy.SetupWebhookWithManager,
		autosnapshotpolicyattachment.SetupWebhookWithManager,
		capacityreservation.SetupWebhookWithManager,
		command.SetupWebhookWithManager,
		dedicatedhost.SetupWebhookWithManager,
		dedicatedhostcluster.SetupWebhookWithManager,
		deploymentset.SetupWebhookWithManager,
		disk.SetupWebhookWithManager,
		diskattachment.SetupWebhookWithManager,
		elasticityassurance.SetupWebhookWithManager,
		hpccluster.SetupWebhookWithManager,
		image.SetupWebhookWithManager,
		imagecomponent.SetupWebhookWithManager,
		imagecopy.SetupWebhookWithManager,
		imageexport.SetupWebhookWithManager,
		imageimport.SetupWebhookWithManager,
		imagepipeline.SetupWebhookWithManager,
		imagepipelineexecution.SetupWebhookWithManager,
		imagesharepermission.SetupWebhookWithManager,
		instance.SetupWebhookWithManager,
		instanceset.SetupWebhookWithManager,
		invocation.SetupWebhookWithManager,
		keypair.SetupWebhookWithManager,
		keypairattachment.SetupWebhookWithManager,
		launchtemplate.SetupWebhookWithManager,
		networkinterface.SetupWebhookWithManager,
		networkinterfaceattachment.SetupWebhookWithManager,
		networkinterfacepermission.SetupWebhookWithManager,
		prefixlist.SetupWebhookWithManager,
		reservedinstance.SetupWebhookWithManager,
		securitygroup.SetupWebhookWithManager,
		securitygrouprule.SetupWebhookWithManager,
		sessionmanagerstatus.SetupWebhookWithManager,
		snapshot.SetupWebhookWithManager,
		snapshotgroup.SetupWebhookWithManager,
		storagecapacityunit.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
