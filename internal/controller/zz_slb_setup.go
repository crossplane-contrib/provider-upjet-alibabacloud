// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	acl "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/slb/acl"
	loadbalancer "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/slb/loadbalancer"
)

// Setup_slb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_slb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.Setup,
		loadbalancer.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
