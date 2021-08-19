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
// +groupName=load.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/load/v1alpha1"
)

type LoadBalancerPolicyObservation struct {
}

type LoadBalancerPolicyParameters struct {
	LoadBalancerName string `json:"loadBalancerName" tf:"load_balancer_name"`

	PolicyAttribute []PolicyAttributeParameters `json:"policyAttribute,omitempty" tf:"policy_attribute"`

	PolicyName string `json:"policyName" tf:"policy_name"`

	PolicyTypeName string `json:"policyTypeName" tf:"policy_type_name"`
}

type PolicyAttributeObservation struct {
}

type PolicyAttributeParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`

	Value *string `json:"value,omitempty" tf:"value"`
}

// LoadBalancerPolicySpec defines the desired state of LoadBalancerPolicy
type LoadBalancerPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LoadBalancerPolicyParameters `json:"forProvider"`
}

// LoadBalancerPolicyStatus defines the observed state of LoadBalancerPolicy.
type LoadBalancerPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LoadBalancerPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerPolicy is the Schema for the LoadBalancerPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LoadBalancerPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerPolicySpec   `json:"spec"`
	Status            LoadBalancerPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerPolicyList contains a list of LoadBalancerPolicys
type LoadBalancerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerPolicy `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerPolicyKind             = "LoadBalancerPolicy"
	LoadBalancerPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: LoadBalancerPolicyKind}.String()
	LoadBalancerPolicyKindAPIVersion   = LoadBalancerPolicyKind + "." + v1alpha1.GroupVersion.String()
	LoadBalancerPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(LoadBalancerPolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&LoadBalancerPolicy{}, &LoadBalancerPolicyList{})
}
