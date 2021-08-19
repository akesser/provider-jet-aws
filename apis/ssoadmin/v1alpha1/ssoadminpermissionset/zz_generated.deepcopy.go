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
func (in *SsoadminPermissionSet) DeepCopyInto(out *SsoadminPermissionSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminPermissionSet.
func (in *SsoadminPermissionSet) DeepCopy() *SsoadminPermissionSet {
	if in == nil {
		return nil
	}
	out := new(SsoadminPermissionSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SsoadminPermissionSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoadminPermissionSetList) DeepCopyInto(out *SsoadminPermissionSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SsoadminPermissionSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminPermissionSetList.
func (in *SsoadminPermissionSetList) DeepCopy() *SsoadminPermissionSetList {
	if in == nil {
		return nil
	}
	out := new(SsoadminPermissionSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SsoadminPermissionSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoadminPermissionSetObservation) DeepCopyInto(out *SsoadminPermissionSetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminPermissionSetObservation.
func (in *SsoadminPermissionSetObservation) DeepCopy() *SsoadminPermissionSetObservation {
	if in == nil {
		return nil
	}
	out := new(SsoadminPermissionSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoadminPermissionSetParameters) DeepCopyInto(out *SsoadminPermissionSetParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.RelayState != nil {
		in, out := &in.RelayState, &out.RelayState
		*out = new(string)
		**out = **in
	}
	if in.SessionDuration != nil {
		in, out := &in.SessionDuration, &out.SessionDuration
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminPermissionSetParameters.
func (in *SsoadminPermissionSetParameters) DeepCopy() *SsoadminPermissionSetParameters {
	if in == nil {
		return nil
	}
	out := new(SsoadminPermissionSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoadminPermissionSetSpec) DeepCopyInto(out *SsoadminPermissionSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminPermissionSetSpec.
func (in *SsoadminPermissionSetSpec) DeepCopy() *SsoadminPermissionSetSpec {
	if in == nil {
		return nil
	}
	out := new(SsoadminPermissionSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoadminPermissionSetStatus) DeepCopyInto(out *SsoadminPermissionSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminPermissionSetStatus.
func (in *SsoadminPermissionSetStatus) DeepCopy() *SsoadminPermissionSetStatus {
	if in == nil {
		return nil
	}
	out := new(SsoadminPermissionSetStatus)
	in.DeepCopyInto(out)
	return out
}
