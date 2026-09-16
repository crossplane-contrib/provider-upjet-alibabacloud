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
)

// Setup_ack creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ack(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfig.Setup,
		edgekubernetes.Setup,
		kubernetes.Setup,
		kubernetesaddon.Setup,
		kubernetesnodepool.Setup,
		kubernetespermissions.Setup,
		managedkubernetes.Setup,
		serverlesskubernetes.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ack creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ack(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfig.SetupGated,
		edgekubernetes.SetupGated,
		kubernetes.SetupGated,
		kubernetesaddon.SetupGated,
		kubernetesnodepool.SetupGated,
		kubernetespermissions.SetupGated,
		managedkubernetes.SetupGated,
		serverlesskubernetes.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_ack registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_ack(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		autoscalingconfig.SetupWebhookWithManager,
		edgekubernetes.SetupWebhookWithManager,
		kubernetes.SetupWebhookWithManager,
		kubernetesaddon.SetupWebhookWithManager,
		kubernetesnodepool.SetupWebhookWithManager,
		kubernetespermissions.SetupWebhookWithManager,
		managedkubernetes.SetupWebhookWithManager,
		serverlesskubernetes.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
