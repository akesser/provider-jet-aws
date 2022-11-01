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
	v1alpha2 "github.com/dkb-bank/provider-jet-aws/apis/kms/v1alpha2"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this HostedZoneDNSSEC.
func (mg *HostedZoneDNSSEC) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HostedZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HostedZoneIDRef,
		Selector:     mg.Spec.ForProvider.HostedZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HostedZoneID")
	}
	mg.Spec.ForProvider.HostedZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HostedZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KeySigningKey.
func (mg *KeySigningKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HostedZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HostedZoneIDRef,
		Selector:     mg.Spec.ForProvider.HostedZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HostedZoneID")
	}
	mg.Spec.ForProvider.HostedZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HostedZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyManagementServiceArn),
		Extract:      v1alpha2.KMSKeyARN(),
		Reference:    mg.Spec.ForProvider.KeyManagementServiceArnRef,
		Selector:     mg.Spec.ForProvider.KeyManagementServiceArnSelector,
		To: reference.To{
			List:    &v1alpha2.KeyList{},
			Managed: &v1alpha2.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyManagementServiceArn")
	}
	mg.Spec.ForProvider.KeyManagementServiceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyManagementServiceArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Record.
func (mg *Record) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HealthCheckID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HealthCheckIDRef,
		Selector:     mg.Spec.ForProvider.HealthCheckIDSelector,
		To: reference.To{
			List:    &HealthCheckList{},
			Managed: &HealthCheck{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HealthCheckID")
	}
	mg.Spec.ForProvider.HealthCheckID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HealthCheckIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VPCAssociationAuthorization.
func (mg *VPCAssociationAuthorization) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VpcIdRef,
		Selector:     mg.Spec.ForProvider.VpcIdSelector,
		To: reference.To{
			List:    &v1alpha21.VPCList{},
			Managed: &v1alpha21.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VpcIdRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Zone.
func (mg *Zone) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DelegationSetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DelegationSetIDRef,
		Selector:     mg.Spec.ForProvider.DelegationSetIDSelector,
		To: reference.To{
			List:    &DelegationSetList{},
			Managed: &DelegationSet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DelegationSetID")
	}
	mg.Spec.ForProvider.DelegationSetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DelegationSetIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPC); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPC[i3].VPCID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.VPC[i3].VpcIdRef,
			Selector:     mg.Spec.ForProvider.VPC[i3].VpcIdSelector,
			To: reference.To{
				List:    &v1alpha21.VPCList{},
				Managed: &v1alpha21.VPC{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPC[i3].VPCID")
		}
		mg.Spec.ForProvider.VPC[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VPC[i3].VpcIdRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ZoneAssociation.
func (mg *ZoneAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VpcIdRef,
		Selector:     mg.Spec.ForProvider.VpcIdSelector,
		To: reference.To{
			List:    &v1alpha21.VPCList{},
			Managed: &v1alpha21.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VpcIdRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}
