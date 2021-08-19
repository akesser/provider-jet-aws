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

// Code generated by terrajet. DO NOT EDIT.

// +kubebuilder:object:generate=true
// +groupName=iot.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/iot/v1alpha1"
)

type IotPolicyObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DefaultVersionId string `json:"defaultVersionId" tf:"default_version_id"`
}

type IotPolicyParameters struct {
	Name string `json:"name" tf:"name"`

	Policy string `json:"policy" tf:"policy"`
}

// IotPolicySpec defines the desired state of IotPolicy
type IotPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IotPolicyParameters `json:"forProvider"`
}

// IotPolicyStatus defines the observed state of IotPolicy.
type IotPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IotPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IotPolicy is the Schema for the IotPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IotPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotPolicySpec   `json:"spec"`
	Status            IotPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotPolicyList contains a list of IotPolicys
type IotPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotPolicy `json:"items"`
}

// Repository type metadata.
var (
	IotPolicyKind             = "IotPolicy"
	IotPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: IotPolicyKind}.String()
	IotPolicyKindAPIVersion   = IotPolicyKind + "." + v1alpha1.GroupVersion.String()
	IotPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(IotPolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&IotPolicy{}, &IotPolicyList{})
}
