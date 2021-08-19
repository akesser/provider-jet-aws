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
func (in *CognitoResourceServer) DeepCopyInto(out *CognitoResourceServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoResourceServer.
func (in *CognitoResourceServer) DeepCopy() *CognitoResourceServer {
	if in == nil {
		return nil
	}
	out := new(CognitoResourceServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CognitoResourceServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoResourceServerList) DeepCopyInto(out *CognitoResourceServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CognitoResourceServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoResourceServerList.
func (in *CognitoResourceServerList) DeepCopy() *CognitoResourceServerList {
	if in == nil {
		return nil
	}
	out := new(CognitoResourceServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CognitoResourceServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoResourceServerObservation) DeepCopyInto(out *CognitoResourceServerObservation) {
	*out = *in
	if in.ScopeIdentifiers != nil {
		in, out := &in.ScopeIdentifiers, &out.ScopeIdentifiers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoResourceServerObservation.
func (in *CognitoResourceServerObservation) DeepCopy() *CognitoResourceServerObservation {
	if in == nil {
		return nil
	}
	out := new(CognitoResourceServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoResourceServerParameters) DeepCopyInto(out *CognitoResourceServerParameters) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = make([]ScopeParameters, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoResourceServerParameters.
func (in *CognitoResourceServerParameters) DeepCopy() *CognitoResourceServerParameters {
	if in == nil {
		return nil
	}
	out := new(CognitoResourceServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoResourceServerSpec) DeepCopyInto(out *CognitoResourceServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoResourceServerSpec.
func (in *CognitoResourceServerSpec) DeepCopy() *CognitoResourceServerSpec {
	if in == nil {
		return nil
	}
	out := new(CognitoResourceServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoResourceServerStatus) DeepCopyInto(out *CognitoResourceServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoResourceServerStatus.
func (in *CognitoResourceServerStatus) DeepCopy() *CognitoResourceServerStatus {
	if in == nil {
		return nil
	}
	out := new(CognitoResourceServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopeObservation) DeepCopyInto(out *ScopeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopeObservation.
func (in *ScopeObservation) DeepCopy() *ScopeObservation {
	if in == nil {
		return nil
	}
	out := new(ScopeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopeParameters) DeepCopyInto(out *ScopeParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopeParameters.
func (in *ScopeParameters) DeepCopy() *ScopeParameters {
	if in == nil {
		return nil
	}
	out := new(ScopeParameters)
	in.DeepCopyInto(out)
	return out
}
