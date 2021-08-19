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
func (in *CreateRuleObservation) DeepCopyInto(out *CreateRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateRuleObservation.
func (in *CreateRuleObservation) DeepCopy() *CreateRuleObservation {
	if in == nil {
		return nil
	}
	out := new(CreateRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateRuleParameters) DeepCopyInto(out *CreateRuleParameters) {
	*out = *in
	if in.IntervalUnit != nil {
		in, out := &in.IntervalUnit, &out.IntervalUnit
		*out = new(string)
		**out = **in
	}
	if in.Times != nil {
		in, out := &in.Times, &out.Times
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateRuleParameters.
func (in *CreateRuleParameters) DeepCopy() *CreateRuleParameters {
	if in == nil {
		return nil
	}
	out := new(CreateRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DlmLifecyclePolicy) DeepCopyInto(out *DlmLifecyclePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DlmLifecyclePolicy.
func (in *DlmLifecyclePolicy) DeepCopy() *DlmLifecyclePolicy {
	if in == nil {
		return nil
	}
	out := new(DlmLifecyclePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DlmLifecyclePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DlmLifecyclePolicyList) DeepCopyInto(out *DlmLifecyclePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DlmLifecyclePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DlmLifecyclePolicyList.
func (in *DlmLifecyclePolicyList) DeepCopy() *DlmLifecyclePolicyList {
	if in == nil {
		return nil
	}
	out := new(DlmLifecyclePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DlmLifecyclePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DlmLifecyclePolicyObservation) DeepCopyInto(out *DlmLifecyclePolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DlmLifecyclePolicyObservation.
func (in *DlmLifecyclePolicyObservation) DeepCopy() *DlmLifecyclePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(DlmLifecyclePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DlmLifecyclePolicyParameters) DeepCopyInto(out *DlmLifecyclePolicyParameters) {
	*out = *in
	if in.PolicyDetails != nil {
		in, out := &in.PolicyDetails, &out.PolicyDetails
		*out = make([]PolicyDetailsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.State != nil {
		in, out := &in.State, &out.State
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DlmLifecyclePolicyParameters.
func (in *DlmLifecyclePolicyParameters) DeepCopy() *DlmLifecyclePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(DlmLifecyclePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DlmLifecyclePolicySpec) DeepCopyInto(out *DlmLifecyclePolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DlmLifecyclePolicySpec.
func (in *DlmLifecyclePolicySpec) DeepCopy() *DlmLifecyclePolicySpec {
	if in == nil {
		return nil
	}
	out := new(DlmLifecyclePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DlmLifecyclePolicyStatus) DeepCopyInto(out *DlmLifecyclePolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DlmLifecyclePolicyStatus.
func (in *DlmLifecyclePolicyStatus) DeepCopy() *DlmLifecyclePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(DlmLifecyclePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDetailsObservation) DeepCopyInto(out *PolicyDetailsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDetailsObservation.
func (in *PolicyDetailsObservation) DeepCopy() *PolicyDetailsObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyDetailsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDetailsParameters) DeepCopyInto(out *PolicyDetailsParameters) {
	*out = *in
	if in.ResourceTypes != nil {
		in, out := &in.ResourceTypes, &out.ResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]ScheduleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TargetTags != nil {
		in, out := &in.TargetTags, &out.TargetTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDetailsParameters.
func (in *PolicyDetailsParameters) DeepCopy() *PolicyDetailsParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyDetailsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetainRuleObservation) DeepCopyInto(out *RetainRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetainRuleObservation.
func (in *RetainRuleObservation) DeepCopy() *RetainRuleObservation {
	if in == nil {
		return nil
	}
	out := new(RetainRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetainRuleParameters) DeepCopyInto(out *RetainRuleParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetainRuleParameters.
func (in *RetainRuleParameters) DeepCopy() *RetainRuleParameters {
	if in == nil {
		return nil
	}
	out := new(RetainRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleObservation) DeepCopyInto(out *ScheduleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleObservation.
func (in *ScheduleObservation) DeepCopy() *ScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(ScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleParameters) DeepCopyInto(out *ScheduleParameters) {
	*out = *in
	if in.CopyTags != nil {
		in, out := &in.CopyTags, &out.CopyTags
		*out = new(bool)
		**out = **in
	}
	if in.CreateRule != nil {
		in, out := &in.CreateRule, &out.CreateRule
		*out = make([]CreateRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RetainRule != nil {
		in, out := &in.RetainRule, &out.RetainRule
		*out = make([]RetainRuleParameters, len(*in))
		copy(*out, *in)
	}
	if in.TagsToAdd != nil {
		in, out := &in.TagsToAdd, &out.TagsToAdd
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleParameters.
func (in *ScheduleParameters) DeepCopy() *ScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(ScheduleParameters)
	in.DeepCopyInto(out)
	return out
}
