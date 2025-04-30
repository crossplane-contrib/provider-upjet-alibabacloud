// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/account"
	auditlogconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/auditlogconfig"
	connection "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/connection"
	instance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/instance"
	tairinstance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/tair/tairinstance"
)

// Setup_tair creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_tair(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		auditlogconfig.Setup,
		connection.Setup,
		instance.Setup,
		tairinstance.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
