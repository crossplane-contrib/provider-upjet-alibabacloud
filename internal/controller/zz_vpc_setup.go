// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	vpc "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vpc"
	vswitch "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vswitch"
)

// Setup_vpc creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_vpc(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		vpc.Setup,
		vswitch.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
