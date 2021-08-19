// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsConfigurationObservation) DeepCopyInto(out *AnalyticsConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsConfigurationObservation.
func (in *AnalyticsConfigurationObservation) DeepCopy() *AnalyticsConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(AnalyticsConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsConfigurationParameters) DeepCopyInto(out *AnalyticsConfigurationParameters) {
	*out = *in
	if in.ApplicationArn != nil {
		in, out := &in.ApplicationArn, &out.ApplicationArn
		*out = new(string)
		**out = **in
	}
	if in.ApplicationId != nil {
		in, out := &in.ApplicationId, &out.ApplicationId
		*out = new(string)
		**out = **in
	}
	if in.ExternalId != nil {
		in, out := &in.ExternalId, &out.ExternalId
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.UserDataShared != nil {
		in, out := &in.UserDataShared, &out.UserDataShared
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsConfigurationParameters.
func (in *AnalyticsConfigurationParameters) DeepCopy() *AnalyticsConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(AnalyticsConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClient) DeepCopyInto(out *CognitoUserPoolClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClient.
func (in *CognitoUserPoolClient) DeepCopy() *CognitoUserPoolClient {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CognitoUserPoolClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClientList) DeepCopyInto(out *CognitoUserPoolClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CognitoUserPoolClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClientList.
func (in *CognitoUserPoolClientList) DeepCopy() *CognitoUserPoolClientList {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CognitoUserPoolClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClientObservation) DeepCopyInto(out *CognitoUserPoolClientObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClientObservation.
func (in *CognitoUserPoolClientObservation) DeepCopy() *CognitoUserPoolClientObservation {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClientObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClientParameters) DeepCopyInto(out *CognitoUserPoolClientParameters) {
	*out = *in
	if in.AccessTokenValidity != nil {
		in, out := &in.AccessTokenValidity, &out.AccessTokenValidity
		*out = new(int64)
		**out = **in
	}
	if in.AllowedOauthFlows != nil {
		in, out := &in.AllowedOauthFlows, &out.AllowedOauthFlows
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedOauthFlowsUserPoolClient != nil {
		in, out := &in.AllowedOauthFlowsUserPoolClient, &out.AllowedOauthFlowsUserPoolClient
		*out = new(bool)
		**out = **in
	}
	if in.AllowedOauthScopes != nil {
		in, out := &in.AllowedOauthScopes, &out.AllowedOauthScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AnalyticsConfiguration != nil {
		in, out := &in.AnalyticsConfiguration, &out.AnalyticsConfiguration
		*out = make([]AnalyticsConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CallbackUrls != nil {
		in, out := &in.CallbackUrls, &out.CallbackUrls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultRedirectUri != nil {
		in, out := &in.DefaultRedirectUri, &out.DefaultRedirectUri
		*out = new(string)
		**out = **in
	}
	if in.EnableTokenRevocation != nil {
		in, out := &in.EnableTokenRevocation, &out.EnableTokenRevocation
		*out = new(bool)
		**out = **in
	}
	if in.ExplicitAuthFlows != nil {
		in, out := &in.ExplicitAuthFlows, &out.ExplicitAuthFlows
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GenerateSecret != nil {
		in, out := &in.GenerateSecret, &out.GenerateSecret
		*out = new(bool)
		**out = **in
	}
	if in.IdTokenValidity != nil {
		in, out := &in.IdTokenValidity, &out.IdTokenValidity
		*out = new(int64)
		**out = **in
	}
	if in.LogoutUrls != nil {
		in, out := &in.LogoutUrls, &out.LogoutUrls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PreventUserExistenceErrors != nil {
		in, out := &in.PreventUserExistenceErrors, &out.PreventUserExistenceErrors
		*out = new(string)
		**out = **in
	}
	if in.ReadAttributes != nil {
		in, out := &in.ReadAttributes, &out.ReadAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RefreshTokenValidity != nil {
		in, out := &in.RefreshTokenValidity, &out.RefreshTokenValidity
		*out = new(int64)
		**out = **in
	}
	if in.SupportedIdentityProviders != nil {
		in, out := &in.SupportedIdentityProviders, &out.SupportedIdentityProviders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TokenValidityUnits != nil {
		in, out := &in.TokenValidityUnits, &out.TokenValidityUnits
		*out = make([]TokenValidityUnitsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WriteAttributes != nil {
		in, out := &in.WriteAttributes, &out.WriteAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClientParameters.
func (in *CognitoUserPoolClientParameters) DeepCopy() *CognitoUserPoolClientParameters {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClientParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClientSpec) DeepCopyInto(out *CognitoUserPoolClientSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClientSpec.
func (in *CognitoUserPoolClientSpec) DeepCopy() *CognitoUserPoolClientSpec {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolClientStatus) DeepCopyInto(out *CognitoUserPoolClientStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolClientStatus.
func (in *CognitoUserPoolClientStatus) DeepCopy() *CognitoUserPoolClientStatus {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolClientStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenValidityUnitsObservation) DeepCopyInto(out *TokenValidityUnitsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenValidityUnitsObservation.
func (in *TokenValidityUnitsObservation) DeepCopy() *TokenValidityUnitsObservation {
	if in == nil {
		return nil
	}
	out := new(TokenValidityUnitsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenValidityUnitsParameters) DeepCopyInto(out *TokenValidityUnitsParameters) {
	*out = *in
	if in.AccessToken != nil {
		in, out := &in.AccessToken, &out.AccessToken
		*out = new(string)
		**out = **in
	}
	if in.IdToken != nil {
		in, out := &in.IdToken, &out.IdToken
		*out = new(string)
		**out = **in
	}
	if in.RefreshToken != nil {
		in, out := &in.RefreshToken, &out.RefreshToken
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenValidityUnitsParameters.
func (in *TokenValidityUnitsParameters) DeepCopy() *TokenValidityUnitsParameters {
	if in == nil {
		return nil
	}
	out := new(TokenValidityUnitsParameters)
	in.DeepCopyInto(out)
	return out
}
