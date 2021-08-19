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
func (in *Macie2InvitationAccepter) DeepCopyInto(out *Macie2InvitationAccepter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Macie2InvitationAccepter.
func (in *Macie2InvitationAccepter) DeepCopy() *Macie2InvitationAccepter {
	if in == nil {
		return nil
	}
	out := new(Macie2InvitationAccepter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Macie2InvitationAccepter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Macie2InvitationAccepterList) DeepCopyInto(out *Macie2InvitationAccepterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Macie2InvitationAccepter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Macie2InvitationAccepterList.
func (in *Macie2InvitationAccepterList) DeepCopy() *Macie2InvitationAccepterList {
	if in == nil {
		return nil
	}
	out := new(Macie2InvitationAccepterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Macie2InvitationAccepterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Macie2InvitationAccepterObservation) DeepCopyInto(out *Macie2InvitationAccepterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Macie2InvitationAccepterObservation.
func (in *Macie2InvitationAccepterObservation) DeepCopy() *Macie2InvitationAccepterObservation {
	if in == nil {
		return nil
	}
	out := new(Macie2InvitationAccepterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Macie2InvitationAccepterParameters) DeepCopyInto(out *Macie2InvitationAccepterParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Macie2InvitationAccepterParameters.
func (in *Macie2InvitationAccepterParameters) DeepCopy() *Macie2InvitationAccepterParameters {
	if in == nil {
		return nil
	}
	out := new(Macie2InvitationAccepterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Macie2InvitationAccepterSpec) DeepCopyInto(out *Macie2InvitationAccepterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Macie2InvitationAccepterSpec.
func (in *Macie2InvitationAccepterSpec) DeepCopy() *Macie2InvitationAccepterSpec {
	if in == nil {
		return nil
	}
	out := new(Macie2InvitationAccepterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Macie2InvitationAccepterStatus) DeepCopyInto(out *Macie2InvitationAccepterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Macie2InvitationAccepterStatus.
func (in *Macie2InvitationAccepterStatus) DeepCopy() *Macie2InvitationAccepterStatus {
	if in == nil {
		return nil
	}
	out := new(Macie2InvitationAccepterStatus)
	in.DeepCopyInto(out)
	return out
}
