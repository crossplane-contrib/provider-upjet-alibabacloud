// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha11 "github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/ecs/v1alpha1"
	v1alpha1 "github.com/crossplane-contrib/provider-upjet-alibabacloud/apis/vpc/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AclEntryAttachment.
func (mg *AclEntryAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ACLID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ACLIDRef,
		Selector:     mg.Spec.ForProvider.ACLIDSelector,
		To: reference.To{
			List:    &AclList{},
			Managed: &Acl{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ACLID")
	}
	mg.Spec.ForProvider.ACLID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ACLIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ACLID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ACLIDRef,
		Selector:     mg.Spec.InitProvider.ACLIDSelector,
		To: reference.To{
			List:    &AclList{},
			Managed: &Acl{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ACLID")
	}
	mg.Spec.InitProvider.ACLID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ACLIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Ascript.
func (mg *Ascript) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ListenerIDRef,
		Selector:     mg.Spec.ForProvider.ListenerIDSelector,
		To: reference.To{
			List:    &ListenerList{},
			Managed: &Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenerID")
	}
	mg.Spec.ForProvider.ListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ListenerIDRef,
		Selector:     mg.Spec.InitProvider.ListenerIDSelector,
		To: reference.To{
			List:    &ListenerList{},
			Managed: &Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ListenerID")
	}
	mg.Spec.InitProvider.ListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ListenerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Listener.
func (mg *Listener) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DefaultActions); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.DefaultActions[i3].ForwardGroupConfig); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDRef,
					Selector:     mg.Spec.ForProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDSelector,
					To: reference.To{
						List:    &ServerGroupList{},
						Managed: &ServerGroup{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID")
				}
				mg.Spec.ForProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDRef = rsp.ResolvedReference

			}
		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LoadBalancerIDRef,
		Selector:     mg.Spec.ForProvider.LoadBalancerIDSelector,
		To: reference.To{
			List:    &LoadBalancerList{},
			Managed: &LoadBalancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancerID")
	}
	mg.Spec.ForProvider.LoadBalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityPolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecurityPolicyIDRef,
		Selector:     mg.Spec.ForProvider.SecurityPolicyIDSelector,
		To: reference.To{
			List:    &SecurityPolicyList{},
			Managed: &SecurityPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityPolicyID")
	}
	mg.Spec.ForProvider.SecurityPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityPolicyIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.DefaultActions); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.DefaultActions[i3].ForwardGroupConfig); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDRef,
					Selector:     mg.Spec.InitProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDSelector,
					To: reference.To{
						List:    &ServerGroupList{},
						Managed: &ServerGroup{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID")
				}
				mg.Spec.InitProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.DefaultActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDRef = rsp.ResolvedReference

			}
		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LoadBalancerIDRef,
		Selector:     mg.Spec.InitProvider.LoadBalancerIDSelector,
		To: reference.To{
			List:    &LoadBalancerList{},
			Managed: &LoadBalancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancerID")
	}
	mg.Spec.InitProvider.LoadBalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SecurityPolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SecurityPolicyIDRef,
		Selector:     mg.Spec.InitProvider.SecurityPolicyIDSelector,
		To: reference.To{
			List:    &SecurityPolicyList{},
			Managed: &SecurityPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityPolicyID")
	}
	mg.Spec.InitProvider.SecurityPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SecurityPolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ListenerAclAttachment.
func (mg *ListenerAclAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ACLID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ACLIDRef,
		Selector:     mg.Spec.ForProvider.ACLIDSelector,
		To: reference.To{
			List:    &AclList{},
			Managed: &Acl{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ACLID")
	}
	mg.Spec.ForProvider.ACLID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ACLIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ListenerIDRef,
		Selector:     mg.Spec.ForProvider.ListenerIDSelector,
		To: reference.To{
			List:    &ListenerList{},
			Managed: &Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenerID")
	}
	mg.Spec.ForProvider.ListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ACLID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ACLIDRef,
		Selector:     mg.Spec.InitProvider.ACLIDSelector,
		To: reference.To{
			List:    &AclList{},
			Managed: &Acl{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ACLID")
	}
	mg.Spec.InitProvider.ACLID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ACLIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ListenerIDRef,
		Selector:     mg.Spec.InitProvider.ListenerIDSelector,
		To: reference.To{
			List:    &ListenerList{},
			Managed: &Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ListenerID")
	}
	mg.Spec.InitProvider.ListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ListenerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LoadBalancer.
func (mg *LoadBalancer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ZoneMappings); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneMappings[i3].VswitchID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ZoneMappings[i3].VswitchIDRef,
			Selector:     mg.Spec.ForProvider.ZoneMappings[i3].VswitchIDSelector,
			To: reference.To{
				List:    &v1alpha1.VswitchList{},
				Managed: &v1alpha1.Vswitch{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ZoneMappings[i3].VswitchID")
		}
		mg.Spec.ForProvider.ZoneMappings[i3].VswitchID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ZoneMappings[i3].VswitchIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VPCIDRef,
		Selector:     mg.Spec.InitProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCID")
	}
	mg.Spec.InitProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.ZoneMappings); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneMappings[i3].VswitchID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ZoneMappings[i3].VswitchIDRef,
			Selector:     mg.Spec.InitProvider.ZoneMappings[i3].VswitchIDSelector,
			To: reference.To{
				List:    &v1alpha1.VswitchList{},
				Managed: &v1alpha1.Vswitch{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ZoneMappings[i3].VswitchID")
		}
		mg.Spec.InitProvider.ZoneMappings[i3].VswitchID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ZoneMappings[i3].VswitchIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this LoadBalancerSecurityGroupAttachment.
func (mg *LoadBalancerSecurityGroupAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LoadBalancerIDRef,
		Selector:     mg.Spec.ForProvider.LoadBalancerIDSelector,
		To: reference.To{
			List:    &LoadBalancerList{},
			Managed: &LoadBalancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancerID")
	}
	mg.Spec.ForProvider.LoadBalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LoadBalancerIDRef,
		Selector:     mg.Spec.InitProvider.LoadBalancerIDSelector,
		To: reference.To{
			List:    &LoadBalancerList{},
			Managed: &LoadBalancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancerID")
	}
	mg.Spec.InitProvider.LoadBalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.InitProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroupID")
	}
	mg.Spec.InitProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SecurityGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LoadBalancerZoneShiftedAttachment.
func (mg *LoadBalancerZoneShiftedAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LoadBalancerIDRef,
		Selector:     mg.Spec.ForProvider.LoadBalancerIDSelector,
		To: reference.To{
			List:    &LoadBalancerList{},
			Managed: &LoadBalancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancerID")
	}
	mg.Spec.ForProvider.LoadBalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VswitchID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VswitchIDRef,
		Selector:     mg.Spec.ForProvider.VswitchIDSelector,
		To: reference.To{
			List:    &v1alpha1.VswitchList{},
			Managed: &v1alpha1.Vswitch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VswitchID")
	}
	mg.Spec.ForProvider.VswitchID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VswitchIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      resource.ExtractParamPath("zone_id", false),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha1.VswitchList{},
			Managed: &v1alpha1.Vswitch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LoadBalancerIDRef,
		Selector:     mg.Spec.InitProvider.LoadBalancerIDSelector,
		To: reference.To{
			List:    &LoadBalancerList{},
			Managed: &LoadBalancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancerID")
	}
	mg.Spec.InitProvider.LoadBalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VswitchID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VswitchIDRef,
		Selector:     mg.Spec.InitProvider.VswitchIDSelector,
		To: reference.To{
			List:    &v1alpha1.VswitchList{},
			Managed: &v1alpha1.Vswitch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VswitchID")
	}
	mg.Spec.InitProvider.VswitchID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VswitchIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
		Extract:      resource.ExtractParamPath("zone_id", false),
		Reference:    mg.Spec.InitProvider.ZoneIDRef,
		Selector:     mg.Spec.InitProvider.ZoneIDSelector,
		To: reference.To{
			List:    &v1alpha1.VswitchList{},
			Managed: &v1alpha1.Vswitch{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Rule.
func (mg *Rule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ListenerIDRef,
		Selector:     mg.Spec.ForProvider.ListenerIDSelector,
		To: reference.To{
			List:    &ListenerList{},
			Managed: &Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenerID")
	}
	mg.Spec.ForProvider.ListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenerIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.RuleActions); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.RuleActions[i3].ForwardGroupConfig); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDRef,
					Selector:     mg.Spec.ForProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDSelector,
					To: reference.To{
						List:    &ServerGroupList{},
						Managed: &ServerGroup{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID")
				}
				mg.Spec.ForProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDRef = rsp.ResolvedReference

			}
		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ListenerIDRef,
		Selector:     mg.Spec.InitProvider.ListenerIDSelector,
		To: reference.To{
			List:    &ListenerList{},
			Managed: &Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ListenerID")
	}
	mg.Spec.InitProvider.ListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ListenerIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.RuleActions); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.RuleActions[i3].ForwardGroupConfig); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDRef,
					Selector:     mg.Spec.InitProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDSelector,
					To: reference.To{
						List:    &ServerGroupList{},
						Managed: &ServerGroup{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID")
				}
				mg.Spec.InitProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.RuleActions[i3].ForwardGroupConfig[i4].ServerGroupTuples[i5].ServerGroupIDRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}

// ResolveReferences of this ServerGroup.
func (mg *ServerGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Servers); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Servers[i3].ServerID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.Servers[i3].ServerIDRef,
			Selector:     mg.Spec.ForProvider.Servers[i3].ServerIDSelector,
			To: reference.To{
				List:    &v1alpha11.InstanceList{},
				Managed: &v1alpha11.Instance{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Servers[i3].ServerID")
		}
		mg.Spec.ForProvider.Servers[i3].ServerID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Servers[i3].ServerIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Servers); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Servers[i3].ServerIP),
			Extract:      resource.ExtractParamPath("private_ip", false),
			Reference:    mg.Spec.ForProvider.Servers[i3].ServerIPRef,
			Selector:     mg.Spec.ForProvider.Servers[i3].ServerIPSelector,
			To: reference.To{
				List:    &v1alpha11.InstanceList{},
				Managed: &v1alpha11.Instance{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Servers[i3].ServerIP")
		}
		mg.Spec.ForProvider.Servers[i3].ServerIP = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Servers[i3].ServerIPRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Servers); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Servers[i3].ServerID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.Servers[i3].ServerIDRef,
			Selector:     mg.Spec.InitProvider.Servers[i3].ServerIDSelector,
			To: reference.To{
				List:    &v1alpha11.InstanceList{},
				Managed: &v1alpha11.Instance{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Servers[i3].ServerID")
		}
		mg.Spec.InitProvider.Servers[i3].ServerID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Servers[i3].ServerIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Servers); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Servers[i3].ServerIP),
			Extract:      resource.ExtractParamPath("private_ip", false),
			Reference:    mg.Spec.InitProvider.Servers[i3].ServerIPRef,
			Selector:     mg.Spec.InitProvider.Servers[i3].ServerIPSelector,
			To: reference.To{
				List:    &v1alpha11.InstanceList{},
				Managed: &v1alpha11.Instance{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Servers[i3].ServerIP")
		}
		mg.Spec.InitProvider.Servers[i3].ServerIP = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Servers[i3].ServerIPRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VPCIDRef,
		Selector:     mg.Spec.InitProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1alpha1.VPCList{},
			Managed: &v1alpha1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCID")
	}
	mg.Spec.InitProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}
