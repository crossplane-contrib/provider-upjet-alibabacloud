// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	common "github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this EndpointAcl.
func (mg *EndpointAcl) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EndpointType),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.ForProvider.EndpointTypeRef,
		Selector:     mg.Spec.ForProvider.EndpointTypeSelector,
		To: reference.To{
			List:    &EndpointList{},
			Managed: &Endpoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EndpointType")
	}
	mg.Spec.ForProvider.EndpointType = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EndpointTypeRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EndpointType),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.InitProvider.EndpointTypeRef,
		Selector:     mg.Spec.InitProvider.EndpointTypeSelector,
		To: reference.To{
			List:    &EndpointList{},
			Managed: &Endpoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EndpointType")
	}
	mg.Spec.InitProvider.EndpointType = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EndpointTypeRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Subscription.
func (mg *Subscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TopicName),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.ForProvider.TopicNameRef,
		Selector:     mg.Spec.ForProvider.TopicNameSelector,
		To: reference.To{
			List:    &TopicList{},
			Managed: &Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TopicName")
	}
	mg.Spec.ForProvider.TopicName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TopicName),
		Extract:      common.IdExtractor(),
		Reference:    mg.Spec.InitProvider.TopicNameRef,
		Selector:     mg.Spec.InitProvider.TopicNameSelector,
		To: reference.To{
			List:    &TopicList{},
			Managed: &Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TopicName")
	}
	mg.Spec.InitProvider.TopicName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TopicNameRef = rsp.ResolvedReference

	return nil
}
