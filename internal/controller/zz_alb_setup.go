// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	acl "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/acl"
	aclentryattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/aclentryattachment"
	ascript "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/ascript"
	healthchecktemplate "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/healthchecktemplate"
	listener "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/listener"
	listeneraclattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/listeneraclattachment"
	listeneradditionalcertificateattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/listeneradditionalcertificateattachment"
	loadbalancer "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/loadbalancer"
	loadbalancersecuritygroupattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/loadbalancersecuritygroupattachment"
	loadbalancerzoneshiftedattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/loadbalancerzoneshiftedattachment"
	rule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/rule"
	securitypolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/securitypolicy"
	servergroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/alb/servergroup"
)

// Setup_alb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_alb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.Setup,
		aclentryattachment.Setup,
		ascript.Setup,
		healthchecktemplate.Setup,
		listener.Setup,
		listeneraclattachment.Setup,
		listeneradditionalcertificateattachment.Setup,
		loadbalancer.Setup,
		loadbalancersecuritygroupattachment.Setup,
		loadbalancerzoneshiftedattachment.Setup,
		rule.Setup,
		securitypolicy.Setup,
		servergroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
