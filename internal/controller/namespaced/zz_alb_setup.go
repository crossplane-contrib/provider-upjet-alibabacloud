// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	acl "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/acl"
	aclentryattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/aclentryattachment"
	ascript "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/ascript"
	healthchecktemplate "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/healthchecktemplate"
	listener "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/listener"
	listeneraclattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/listeneraclattachment"
	loadbalancer "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/loadbalancer"
	loadbalancersecuritygroupattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/loadbalancersecuritygroupattachment"
	loadbalancerzoneshiftedattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/loadbalancerzoneshiftedattachment"
	rule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/rule"
	securitypolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/securitypolicy"
	servergroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/namespaced/alb/servergroup"
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

// SetupGated_alb creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_alb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_alb registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_alb(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
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
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
