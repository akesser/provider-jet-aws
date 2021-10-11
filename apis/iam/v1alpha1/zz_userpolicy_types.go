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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type UserPolicyObservation struct {
}

type UserPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=User
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

// UserPolicySpec defines the desired state of UserPolicy
type UserPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPolicyParameters `json:"forProvider"`
}

// UserPolicyStatus defines the observed state of UserPolicy.
type UserPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserPolicy is the Schema for the UserPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type UserPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPolicySpec   `json:"spec"`
	Status            UserPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPolicyList contains a list of UserPolicys
type UserPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPolicy `json:"items"`
}

// Repository type metadata.
var (
	UserPolicyKind             = "UserPolicy"
	UserPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: UserPolicyKind}.String()
	UserPolicyKindAPIVersion   = UserPolicyKind + "." + GroupVersion.String()
	UserPolicyGroupVersionKind = GroupVersion.WithKind(UserPolicyKind)
)

func init() {
	SchemeBuilder.Register(&UserPolicy{}, &UserPolicyList{})
}
