// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	alias "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/alias"
	instance "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/instance"
	key "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/key"
	secret "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/kms/secret"
)

// Setup_kms creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.Setup,
		instance.Setup,
		key.Setup,
		secret.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
