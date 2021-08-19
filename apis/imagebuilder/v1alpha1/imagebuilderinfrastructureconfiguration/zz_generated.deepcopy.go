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
func (in *ImagebuilderInfrastructureConfiguration) DeepCopyInto(out *ImagebuilderInfrastructureConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderInfrastructureConfiguration.
func (in *ImagebuilderInfrastructureConfiguration) DeepCopy() *ImagebuilderInfrastructureConfiguration {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderInfrastructureConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagebuilderInfrastructureConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagebuilderInfrastructureConfigurationList) DeepCopyInto(out *ImagebuilderInfrastructureConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImagebuilderInfrastructureConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderInfrastructureConfigurationList.
func (in *ImagebuilderInfrastructureConfigurationList) DeepCopy() *ImagebuilderInfrastructureConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderInfrastructureConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagebuilderInfrastructureConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagebuilderInfrastructureConfigurationObservation) DeepCopyInto(out *ImagebuilderInfrastructureConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderInfrastructureConfigurationObservation.
func (in *ImagebuilderInfrastructureConfigurationObservation) DeepCopy() *ImagebuilderInfrastructureConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderInfrastructureConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagebuilderInfrastructureConfigurationParameters) DeepCopyInto(out *ImagebuilderInfrastructureConfigurationParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.InstanceTypes != nil {
		in, out := &in.InstanceTypes, &out.InstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KeyPair != nil {
		in, out := &in.KeyPair, &out.KeyPair
		*out = new(string)
		**out = **in
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = make([]LoggingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceTags != nil {
		in, out := &in.ResourceTags, &out.ResourceTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnsTopicArn != nil {
		in, out := &in.SnsTopicArn, &out.SnsTopicArn
		*out = new(string)
		**out = **in
	}
	if in.SubnetId != nil {
		in, out := &in.SubnetId, &out.SubnetId
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
	if in.TerminateInstanceOnFailure != nil {
		in, out := &in.TerminateInstanceOnFailure, &out.TerminateInstanceOnFailure
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderInfrastructureConfigurationParameters.
func (in *ImagebuilderInfrastructureConfigurationParameters) DeepCopy() *ImagebuilderInfrastructureConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderInfrastructureConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagebuilderInfrastructureConfigurationSpec) DeepCopyInto(out *ImagebuilderInfrastructureConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderInfrastructureConfigurationSpec.
func (in *ImagebuilderInfrastructureConfigurationSpec) DeepCopy() *ImagebuilderInfrastructureConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderInfrastructureConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagebuilderInfrastructureConfigurationStatus) DeepCopyInto(out *ImagebuilderInfrastructureConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagebuilderInfrastructureConfigurationStatus.
func (in *ImagebuilderInfrastructureConfigurationStatus) DeepCopy() *ImagebuilderInfrastructureConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ImagebuilderInfrastructureConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingObservation) DeepCopyInto(out *LoggingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingObservation.
func (in *LoggingObservation) DeepCopy() *LoggingObservation {
	if in == nil {
		return nil
	}
	out := new(LoggingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingParameters) DeepCopyInto(out *LoggingParameters) {
	*out = *in
	if in.S3Logs != nil {
		in, out := &in.S3Logs, &out.S3Logs
		*out = make([]S3LogsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingParameters.
func (in *LoggingParameters) DeepCopy() *LoggingParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3LogsObservation) DeepCopyInto(out *S3LogsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3LogsObservation.
func (in *S3LogsObservation) DeepCopy() *S3LogsObservation {
	if in == nil {
		return nil
	}
	out := new(S3LogsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3LogsParameters) DeepCopyInto(out *S3LogsParameters) {
	*out = *in
	if in.S3KeyPrefix != nil {
		in, out := &in.S3KeyPrefix, &out.S3KeyPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3LogsParameters.
func (in *S3LogsParameters) DeepCopy() *S3LogsParameters {
	if in == nil {
		return nil
	}
	out := new(S3LogsParameters)
	in.DeepCopyInto(out)
	return out
}
