/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha11 "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/lb/v1alpha1"
	common "github.com/crossplane-contrib/provider-tf-aws/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Attachment.
func (mg *Attachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AlbTargetGroupArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.AlbTargetGroupArnRef,
		Selector:     mg.Spec.ForProvider.AlbTargetGroupArnSelector,
		To: reference.To{
			List:    &v1alpha1.LBTargetGroupList{},
			Managed: &v1alpha1.LBTargetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AlbTargetGroupArn")
	}
	mg.Spec.ForProvider.AlbTargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AlbTargetGroupArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AutoscalingGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AutoscalingGroupNameRef,
		Selector:     mg.Spec.ForProvider.AutoscalingGroupNameSelector,
		To: reference.To{
			List:    &AutoscalingGroupList{},
			Managed: &AutoscalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AutoscalingGroupName")
	}
	mg.Spec.ForProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AutoscalingGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AutoscalingGroup.
func (mg *AutoscalingGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.TargetGroupArns),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.TargetGroupArnsRefs,
		Selector:      mg.Spec.ForProvider.TargetGroupArnsSelector,
		To: reference.To{
			List:    &v1alpha1.LBTargetGroupList{},
			Managed: &v1alpha1.LBTargetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetGroupArns")
	}
	mg.Spec.ForProvider.TargetGroupArns = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.TargetGroupArnsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VpcZoneIdentifier),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VpcZoneIdentifierRefs,
		Selector:      mg.Spec.ForProvider.VpcZoneIdentifierSelector,
		To: reference.To{
			List:    &v1alpha11.SubnetList{},
			Managed: &v1alpha11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VpcZoneIdentifier")
	}
	mg.Spec.ForProvider.VpcZoneIdentifier = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VpcZoneIdentifierRefs = mrsp.ResolvedReferences

	return nil
}
