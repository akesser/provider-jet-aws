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

package v1alpha2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha21 "github.com/dkb-bank/provider-jet-aws/apis/ec2/v1alpha2"
	v1alpha22 "github.com/dkb-bank/provider-jet-aws/apis/kms/v1alpha2"
	v1alpha2 "github.com/dkb-bank/provider-jet-aws/apis/s3/v1alpha2"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBSubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBSubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.DBSubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBSubnetGroupName")
	}
	mg.Spec.ForProvider.DBSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBSubnetGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.RestoreToPointInTime); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifierRef,
			Selector:     mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifierSelector,
			To: reference.To{
				List:    &ClusterList{},
				Managed: &Cluster{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifier")
		}
		mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceClusterIdentifierRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.S3Import); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3Import[i3].BucketName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.S3Import[i3].BucketNameRef,
			Selector:     mg.Spec.ForProvider.S3Import[i3].BucketNameSelector,
			To: reference.To{
				List:    &v1alpha2.BucketList{},
				Managed: &v1alpha2.Bucket{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.S3Import[i3].BucketName")
		}
		mg.Spec.ForProvider.S3Import[i3].BucketName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.S3Import[i3].BucketNameRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VpcSecurityGroupIdRefs,
		Selector:      mg.Spec.ForProvider.VpcSecurityGroupIdSelector,
		To: reference.To{
			List:    &v1alpha21.SecurityGroupList{},
			Managed: &v1alpha21.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VpcSecurityGroupIdRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBSubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBSubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.DBSubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBSubnetGroupName")
	}
	mg.Spec.ForProvider.DBSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBSubnetGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha22.KeyList{},
			Managed: &v1alpha22.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ParameterGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ParameterGroupNameRef,
		Selector:     mg.Spec.ForProvider.ParameterGroupNameSelector,
		To: reference.To{
			List:    &ParameterGroupList{},
			Managed: &ParameterGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ParameterGroupName")
	}
	mg.Spec.ForProvider.ParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParameterGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PerformanceInsightsKMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PerformanceInsightsKMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.PerformanceInsightsKMSKeyIDSelector,
		To: reference.To{
			List:    &v1alpha22.KeyList{},
			Managed: &v1alpha22.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PerformanceInsightsKMSKeyID")
	}
	mg.Spec.ForProvider.PerformanceInsightsKMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PerformanceInsightsKMSKeyIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.RestoreToPointInTime); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceDBInstanceIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceDBInstanceIdentifierRef,
			Selector:     mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceDBInstanceIdentifierSelector,
			To: reference.To{
				List:    &InstanceList{},
				Managed: &Instance{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceDBInstanceIdentifier")
		}
		mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceDBInstanceIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.RestoreToPointInTime[i3].SourceDBInstanceIdentifierRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.S3Import); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3Import[i3].BucketName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.S3Import[i3].BucketNameRef,
			Selector:     mg.Spec.ForProvider.S3Import[i3].BucketNameSelector,
			To: reference.To{
				List:    &v1alpha2.BucketList{},
				Managed: &v1alpha2.Bucket{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.S3Import[i3].BucketName")
		}
		mg.Spec.ForProvider.S3Import[i3].BucketName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.S3Import[i3].BucketNameRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupNames),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupNameRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupNameSelector,
		To: reference.To{
			List:    &v1alpha21.SecurityGroupList{},
			Managed: &v1alpha21.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupNames")
	}
	mg.Spec.ForProvider.SecurityGroupNames = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupNameRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VpcSecurityGroupIdRefs,
		Selector:      mg.Spec.ForProvider.VpcSecurityGroupIdSelector,
		To: reference.To{
			List:    &v1alpha21.SecurityGroupList{},
			Managed: &v1alpha21.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VpcSecurityGroupIdRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this SubnetGroup.
func (mg *SubnetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIdRefs,
		Selector:      mg.Spec.ForProvider.SubnetIdSelector,
		To: reference.To{
			List:    &v1alpha21.SubnetList{},
			Managed: &v1alpha21.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIdRefs = mrsp.ResolvedReferences

	return nil
}
