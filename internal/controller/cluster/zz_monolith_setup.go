// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	autoscalingconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ack/autoscalingconfig"
	edgekubernetes "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ack/edgekubernetes"
	kubernetes "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ack/kubernetes"
	kubernetesaddon "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ack/kubernetesaddon"
	kubernetesnodepool "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ack/kubernetesnodepool"
	kubernetespermissions "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ack/kubernetespermissions"
	managedkubernetes "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ack/managedkubernetes"
	serverlesskubernetes "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ack/serverlesskubernetes"
	cluster "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ackone/cluster"
	membershipattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ackone/membershipattachment"
	acl "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/acl"
	aclentryattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/aclentryattachment"
	ascript "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/ascript"
	healthchecktemplate "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/healthchecktemplate"
	listener "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/listener"
	listeneraclattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/listeneraclattachment"
	loadbalancer "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/loadbalancer"
	loadbalancersecuritygroupattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/loadbalancersecuritygroupattachment"
	loadbalancerzoneshiftedattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/loadbalancerzoneshiftedattachment"
	rule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/rule"
	securitypolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/securitypolicy"
	servergroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alb/servergroup"
	addresspool "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/addresspool"
	customline "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/customline"
	domain "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/domain"
	domainattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/domainattachment"
	domaingroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/domaingroup"
	gtminstance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/gtminstance"
	instance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/instance"
	monitorconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/monitorconfig"
	record "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alidns/record"
	consumergroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alikafka/consumergroup"
	instancealikafka "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alikafka/instance"
	instanceallowedipattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alikafka/instanceallowedipattachment"
	saslacl "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alikafka/saslacl"
	sasluser "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alikafka/sasluser"
	scheduledscalingrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alikafka/scheduledscalingrule"
	topic "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/alikafka/topic"
	domaincdn "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cdn/domain"
	domainconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cdn/domainconfig"
	fctrigger "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cdn/fctrigger"
	alarmcontactgroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cloudmonitorservice/alarmcontactgroup"
	chain "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/chain"
	chartnamespace "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/chartnamespace"
	chartrepository "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/chartrepository"
	eeinstance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/eeinstance"
	eenamespace "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/eenamespace"
	eerepo "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/eerepo"
	eesyncrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/eesyncrule"
	endpointaclpolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/endpointaclpolicy"
	registrynamespace "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/registrynamespace"
	repo "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/repo"
	scanrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/scanrule"
	storagedomainroutingrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/storagedomainroutingrule"
	vpcendpointlinkedvpc "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/cr/vpcendpointlinkedvpc"
	activation "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/activation"
	autoprovisioninggroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/autoprovisioninggroup"
	autosnapshotpolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/autosnapshotpolicy"
	autosnapshotpolicyattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/autosnapshotpolicyattachment"
	capacityreservation "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/capacityreservation"
	command "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/command"
	dedicatedhost "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/dedicatedhost"
	dedicatedhostcluster "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/dedicatedhostcluster"
	deploymentset "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/deploymentset"
	disk "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/disk"
	diskattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/diskattachment"
	elasticityassurance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/elasticityassurance"
	hpccluster "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/hpccluster"
	image "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/image"
	imagecomponent "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/imagecomponent"
	imagecopy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/imagecopy"
	imageexport "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/imageexport"
	imageimport "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/imageimport"
	imagepipeline "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/imagepipeline"
	imagepipelineexecution "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/imagepipelineexecution"
	imagesharepermission "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/imagesharepermission"
	instanceecs "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/instance"
	instanceset "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/instanceset"
	invocation "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/invocation"
	keypair "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/keypair"
	keypairattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/keypairattachment"
	launchtemplate "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/launchtemplate"
	networkinterface "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/networkinterface"
	networkinterfaceattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/networkinterfaceattachment"
	networkinterfacepermission "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/networkinterfacepermission"
	prefixlist "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/prefixlist"
	reservedinstance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/reservedinstance"
	securitygroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/securitygrouprule"
	sessionmanagerstatus "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/sessionmanagerstatus"
	snapshot "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/snapshot"
	snapshotgroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/snapshotgroup"
	storagecapacityunit "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/ecs/storagecapacityunit"
	alias "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/fcv3/alias"
	asyncinvokeconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/fcv3/asyncinvokeconfig"
	concurrencyconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/fcv3/concurrencyconfig"
	customdomain "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/fcv3/customdomain"
	function "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/fcv3/function"
	functionversion "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/fcv3/functionversion"
	layerversion "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/fcv3/layerversion"
	provisionconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/fcv3/provisionconfig"
	trigger "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/fcv3/trigger"
	vpcbinding "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/fcv3/vpcbinding"
	aliaskms "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/kms/alias"
	instancekms "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/kms/instance"
	key "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/kms/key"
	secret "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/kms/secret"
	endpoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/messageservice/endpoint"
	endpointacl "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/messageservice/endpointacl"
	queue "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/messageservice/queue"
	subscription "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/messageservice/subscription"
	topicmessageservice "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/messageservice/topic"
	application "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oos/application"
	applicationgroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oos/applicationgroup"
	defaultpatchbaseline "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oos/defaultpatchbaseline"
	execution "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oos/execution"
	parameter "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oos/parameter"
	patchbaseline "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oos/patchbaseline"
	secretparameter "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oos/secretparameter"
	servicesetting "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oos/servicesetting"
	stateconfiguration "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oos/stateconfiguration"
	template "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oos/template"
	accesspoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/accesspoint"
	accountpublicaccessblock "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/accountpublicaccessblock"
	bucket "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucket"
	bucketaccessmonitor "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketaccessmonitor"
	bucketacl "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketacl"
	bucketcname "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketcname"
	bucketcnametoken "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketcnametoken"
	bucketcors "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketcors"
	bucketdataredundancytransition "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketdataredundancytransition"
	buckethttpsconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/buckethttpsconfig"
	bucketlogging "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketlogging"
	bucketmetaquery "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketmetaquery"
	bucketobject "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketobject"
	bucketpolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketpolicy"
	bucketpublicaccessblock "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketpublicaccessblock"
	bucketreferer "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketreferer"
	bucketreplication "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketreplication"
	bucketrequestpayment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketrequestpayment"
	bucketserversideencryption "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketserversideencryption"
	bucketstyle "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketstyle"
	buckettransferacceleration "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/buckettransferacceleration"
	bucketuserdefinedlogfields "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketuserdefinedlogfields"
	bucketversioning "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketversioning"
	bucketwebsite "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketwebsite"
	bucketworm "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/oss/bucketworm"
	account "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/polardb/account"
	accountprivilege "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/polardb/accountprivilege"
	backuppolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/polardb/backuppolicy"
	clusterpolardb "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/polardb/cluster"
	clusterendpoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/polardb/clusterendpoint"
	database "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/polardb/database"
	endpointpolardb "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/polardb/endpoint"
	endpointaddress "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/polardb/endpointaddress"
	globaldatabasenetwork "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/polardb/globaldatabasenetwork"
	parametergroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/polardb/parametergroup"
	primaryendpoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/polardb/primaryendpoint"
	vpcendpoint "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpoint"
	vpcendpointconnection "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpointconnection"
	vpcendpointservice "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpointservice"
	vpcendpointserviceresource "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpointserviceresource"
	vpcendpointserviceuser "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpointserviceuser"
	vpcendpointzone "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/privatelink/vpcendpointzone"
	providerconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/providerconfig"
	quotaalarm "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/quotas/quotaalarm"
	quotaapplication "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/quotas/quotaapplication"
	templateapplications "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/quotas/templateapplications"
	templatequota "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/quotas/templatequota"
	templateservice "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/quotas/templateservice"
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
	aclslb "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/slb/acl"
	listenerslb "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/slb/listener"
	loadbalancerslb "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/slb/loadbalancer"
	accounttair "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/tair/account"
	auditlogconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/tair/auditlogconfig"
	connection "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/tair/connection"
	instancetair "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/tair/instance"
	tairinstance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/tair/tairinstance"
	routetable "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/vpc/routetable"
	vpc "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/vpc/vpc"
	vswitch "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cluster/vpc/vswitch"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfig.Setup,
		edgekubernetes.Setup,
		kubernetes.Setup,
		kubernetesaddon.Setup,
		kubernetesnodepool.Setup,
		kubernetespermissions.Setup,
		managedkubernetes.Setup,
		serverlesskubernetes.Setup,
		cluster.Setup,
		membershipattachment.Setup,
		acl.Setup,
		aclentryattachment.Setup,
		ascript.Setup,
		healthchecktemplate.Setup,
		listener.Setup,
		listeneraclattachment.Setup,
		loadbalancer.Setup,
		loadbalancersecuritygroupattachment.Setup,
		loadbalancerzoneshiftedattachment.Setup,
		rule.Setup,
		securitypolicy.Setup,
		servergroup.Setup,
		addresspool.Setup,
		customline.Setup,
		domain.Setup,
		domainattachment.Setup,
		domaingroup.Setup,
		gtminstance.Setup,
		instance.Setup,
		monitorconfig.Setup,
		record.Setup,
		consumergroup.Setup,
		instancealikafka.Setup,
		instanceallowedipattachment.Setup,
		saslacl.Setup,
		sasluser.Setup,
		scheduledscalingrule.Setup,
		topic.Setup,
		domaincdn.Setup,
		domainconfig.Setup,
		fctrigger.Setup,
		alarmcontactgroup.Setup,
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
		instanceecs.Setup,
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
		alias.Setup,
		asyncinvokeconfig.Setup,
		concurrencyconfig.Setup,
		customdomain.Setup,
		function.Setup,
		functionversion.Setup,
		layerversion.Setup,
		provisionconfig.Setup,
		trigger.Setup,
		vpcbinding.Setup,
		aliaskms.Setup,
		instancekms.Setup,
		key.Setup,
		secret.Setup,
		endpoint.Setup,
		endpointacl.Setup,
		queue.Setup,
		subscription.Setup,
		topicmessageservice.Setup,
		application.Setup,
		applicationgroup.Setup,
		defaultpatchbaseline.Setup,
		execution.Setup,
		parameter.Setup,
		patchbaseline.Setup,
		secretparameter.Setup,
		servicesetting.Setup,
		stateconfiguration.Setup,
		template.Setup,
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
		account.Setup,
		accountprivilege.Setup,
		backuppolicy.Setup,
		clusterpolardb.Setup,
		clusterendpoint.Setup,
		database.Setup,
		endpointpolardb.Setup,
		endpointaddress.Setup,
		globaldatabasenetwork.Setup,
		parametergroup.Setup,
		primaryendpoint.Setup,
		vpcendpoint.Setup,
		vpcendpointconnection.Setup,
		vpcendpointservice.Setup,
		vpcendpointserviceresource.Setup,
		vpcendpointserviceuser.Setup,
		vpcendpointzone.Setup,
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
		passwordpolicy.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		securitypreference.Setup,
		user.Setup,
		usergroupattachment.Setup,
		userpolicyattachment.Setup,
		aclslb.Setup,
		listenerslb.Setup,
		loadbalancerslb.Setup,
		accounttair.Setup,
		auditlogconfig.Setup,
		connection.Setup,
		instancetair.Setup,
		tairinstance.Setup,
		routetable.Setup,
		vpc.Setup,
		vswitch.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfig.SetupGated,
		edgekubernetes.SetupGated,
		kubernetes.SetupGated,
		kubernetesaddon.SetupGated,
		kubernetesnodepool.SetupGated,
		kubernetespermissions.SetupGated,
		managedkubernetes.SetupGated,
		serverlesskubernetes.SetupGated,
		cluster.SetupGated,
		membershipattachment.SetupGated,
		acl.SetupGated,
		aclentryattachment.SetupGated,
		ascript.SetupGated,
		healthchecktemplate.SetupGated,
		listener.SetupGated,
		listeneraclattachment.SetupGated,
		loadbalancer.SetupGated,
		loadbalancersecuritygroupattachment.SetupGated,
		loadbalancerzoneshiftedattachment.SetupGated,
		rule.SetupGated,
		securitypolicy.SetupGated,
		servergroup.SetupGated,
		addresspool.SetupGated,
		customline.SetupGated,
		domain.SetupGated,
		domainattachment.SetupGated,
		domaingroup.SetupGated,
		gtminstance.SetupGated,
		instance.SetupGated,
		monitorconfig.SetupGated,
		record.SetupGated,
		consumergroup.SetupGated,
		instancealikafka.SetupGated,
		instanceallowedipattachment.SetupGated,
		saslacl.SetupGated,
		sasluser.SetupGated,
		scheduledscalingrule.SetupGated,
		topic.SetupGated,
		domaincdn.SetupGated,
		domainconfig.SetupGated,
		fctrigger.SetupGated,
		alarmcontactgroup.SetupGated,
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
		instanceecs.SetupGated,
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
		alias.SetupGated,
		asyncinvokeconfig.SetupGated,
		concurrencyconfig.SetupGated,
		customdomain.SetupGated,
		function.SetupGated,
		functionversion.SetupGated,
		layerversion.SetupGated,
		provisionconfig.SetupGated,
		trigger.SetupGated,
		vpcbinding.SetupGated,
		aliaskms.SetupGated,
		instancekms.SetupGated,
		key.SetupGated,
		secret.SetupGated,
		endpoint.SetupGated,
		endpointacl.SetupGated,
		queue.SetupGated,
		subscription.SetupGated,
		topicmessageservice.SetupGated,
		application.SetupGated,
		applicationgroup.SetupGated,
		defaultpatchbaseline.SetupGated,
		execution.SetupGated,
		parameter.SetupGated,
		patchbaseline.SetupGated,
		secretparameter.SetupGated,
		servicesetting.SetupGated,
		stateconfiguration.SetupGated,
		template.SetupGated,
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
		account.SetupGated,
		accountprivilege.SetupGated,
		backuppolicy.SetupGated,
		clusterpolardb.SetupGated,
		clusterendpoint.SetupGated,
		database.SetupGated,
		endpointpolardb.SetupGated,
		endpointaddress.SetupGated,
		globaldatabasenetwork.SetupGated,
		parametergroup.SetupGated,
		primaryendpoint.SetupGated,
		vpcendpoint.SetupGated,
		vpcendpointconnection.SetupGated,
		vpcendpointservice.SetupGated,
		vpcendpointserviceresource.SetupGated,
		vpcendpointserviceuser.SetupGated,
		vpcendpointzone.SetupGated,
		providerconfig.SetupGated,
		quotaalarm.SetupGated,
		quotaapplication.SetupGated,
		templateapplications.SetupGated,
		templatequota.SetupGated,
		templateservice.SetupGated,
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
		aclslb.SetupGated,
		listenerslb.SetupGated,
		loadbalancerslb.SetupGated,
		accounttair.SetupGated,
		auditlogconfig.SetupGated,
		connection.SetupGated,
		instancetair.SetupGated,
		tairinstance.SetupGated,
		routetable.SetupGated,
		vpc.SetupGated,
		vswitch.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_monolith registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_monolith(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		autoscalingconfig.SetupWebhookWithManager,
		edgekubernetes.SetupWebhookWithManager,
		kubernetes.SetupWebhookWithManager,
		kubernetesaddon.SetupWebhookWithManager,
		kubernetesnodepool.SetupWebhookWithManager,
		kubernetespermissions.SetupWebhookWithManager,
		managedkubernetes.SetupWebhookWithManager,
		serverlesskubernetes.SetupWebhookWithManager,
		cluster.SetupWebhookWithManager,
		membershipattachment.SetupWebhookWithManager,
		acl.SetupWebhookWithManager,
		aclentryattachment.SetupWebhookWithManager,
		ascript.SetupWebhookWithManager,
		healthchecktemplate.SetupWebhookWithManager,
		listener.SetupWebhookWithManager,
		listeneraclattachment.SetupWebhookWithManager,
		loadbalancer.SetupWebhookWithManager,
		loadbalancersecuritygroupattachment.SetupWebhookWithManager,
		loadbalancerzoneshiftedattachment.SetupWebhookWithManager,
		rule.SetupWebhookWithManager,
		securitypolicy.SetupWebhookWithManager,
		servergroup.SetupWebhookWithManager,
		addresspool.SetupWebhookWithManager,
		customline.SetupWebhookWithManager,
		domain.SetupWebhookWithManager,
		domainattachment.SetupWebhookWithManager,
		domaingroup.SetupWebhookWithManager,
		gtminstance.SetupWebhookWithManager,
		instance.SetupWebhookWithManager,
		monitorconfig.SetupWebhookWithManager,
		record.SetupWebhookWithManager,
		consumergroup.SetupWebhookWithManager,
		instancealikafka.SetupWebhookWithManager,
		instanceallowedipattachment.SetupWebhookWithManager,
		saslacl.SetupWebhookWithManager,
		sasluser.SetupWebhookWithManager,
		scheduledscalingrule.SetupWebhookWithManager,
		topic.SetupWebhookWithManager,
		domaincdn.SetupWebhookWithManager,
		domainconfig.SetupWebhookWithManager,
		fctrigger.SetupWebhookWithManager,
		alarmcontactgroup.SetupWebhookWithManager,
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
		instanceecs.SetupWebhookWithManager,
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
		alias.SetupWebhookWithManager,
		asyncinvokeconfig.SetupWebhookWithManager,
		concurrencyconfig.SetupWebhookWithManager,
		customdomain.SetupWebhookWithManager,
		function.SetupWebhookWithManager,
		functionversion.SetupWebhookWithManager,
		layerversion.SetupWebhookWithManager,
		provisionconfig.SetupWebhookWithManager,
		trigger.SetupWebhookWithManager,
		vpcbinding.SetupWebhookWithManager,
		aliaskms.SetupWebhookWithManager,
		instancekms.SetupWebhookWithManager,
		key.SetupWebhookWithManager,
		secret.SetupWebhookWithManager,
		endpoint.SetupWebhookWithManager,
		endpointacl.SetupWebhookWithManager,
		queue.SetupWebhookWithManager,
		subscription.SetupWebhookWithManager,
		topicmessageservice.SetupWebhookWithManager,
		application.SetupWebhookWithManager,
		applicationgroup.SetupWebhookWithManager,
		defaultpatchbaseline.SetupWebhookWithManager,
		execution.SetupWebhookWithManager,
		parameter.SetupWebhookWithManager,
		patchbaseline.SetupWebhookWithManager,
		secretparameter.SetupWebhookWithManager,
		servicesetting.SetupWebhookWithManager,
		stateconfiguration.SetupWebhookWithManager,
		template.SetupWebhookWithManager,
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
		account.SetupWebhookWithManager,
		accountprivilege.SetupWebhookWithManager,
		backuppolicy.SetupWebhookWithManager,
		clusterpolardb.SetupWebhookWithManager,
		clusterendpoint.SetupWebhookWithManager,
		database.SetupWebhookWithManager,
		endpointpolardb.SetupWebhookWithManager,
		endpointaddress.SetupWebhookWithManager,
		globaldatabasenetwork.SetupWebhookWithManager,
		parametergroup.SetupWebhookWithManager,
		primaryendpoint.SetupWebhookWithManager,
		vpcendpoint.SetupWebhookWithManager,
		vpcendpointconnection.SetupWebhookWithManager,
		vpcendpointservice.SetupWebhookWithManager,
		vpcendpointserviceresource.SetupWebhookWithManager,
		vpcendpointserviceuser.SetupWebhookWithManager,
		vpcendpointzone.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		quotaalarm.SetupWebhookWithManager,
		quotaapplication.SetupWebhookWithManager,
		templateapplications.SetupWebhookWithManager,
		templatequota.SetupWebhookWithManager,
		templateservice.SetupWebhookWithManager,
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
		aclslb.SetupWebhookWithManager,
		listenerslb.SetupWebhookWithManager,
		loadbalancerslb.SetupWebhookWithManager,
		accounttair.SetupWebhookWithManager,
		auditlogconfig.SetupWebhookWithManager,
		connection.SetupWebhookWithManager,
		instancetair.SetupWebhookWithManager,
		tairinstance.SetupWebhookWithManager,
		routetable.SetupWebhookWithManager,
		vpc.SetupWebhookWithManager,
		vswitch.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
