// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	logtailattachment "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/sls/logtailattachment"
	logtailconfig "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/sls/logtailconfig"
	machinegroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/sls/machinegroup"
	project "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/sls/project"
	store "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/sls/store"
	storeindex "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/sls/storeindex"
)

// Setup_sls creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sls(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		logtailattachment.Setup,
		logtailconfig.Setup,
		machinegroup.Setup,
		project.Setup,
		store.Setup,
		storeindex.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
