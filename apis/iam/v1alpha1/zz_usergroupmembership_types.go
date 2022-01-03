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

type UserGroupMembershipObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type UserGroupMembershipParameters struct {

	// +kubebuilder:validation:Optional
	GroupRefs []v1.Reference `json:"groupRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	GroupSelector *v1.Selector `json:"groupSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Group
	// +crossplane:generate:reference:refFieldName=GroupRefs
	// +crossplane:generate:reference:selectorFieldName=GroupSelector
	// +kubebuilder:validation:Optional
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// +crossplane:generate:reference:type=User
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

// UserGroupMembershipSpec defines the desired state of UserGroupMembership
type UserGroupMembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserGroupMembershipParameters `json:"forProvider"`
}

// UserGroupMembershipStatus defines the observed state of UserGroupMembership.
type UserGroupMembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserGroupMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserGroupMembership is the Schema for the UserGroupMemberships API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type UserGroupMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserGroupMembershipSpec   `json:"spec"`
	Status            UserGroupMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserGroupMembershipList contains a list of UserGroupMemberships
type UserGroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserGroupMembership `json:"items"`
}

// Repository type metadata.
var (
	UserGroupMembership_Kind             = "UserGroupMembership"
	UserGroupMembership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserGroupMembership_Kind}.String()
	UserGroupMembership_KindAPIVersion   = UserGroupMembership_Kind + "." + CRDGroupVersion.String()
	UserGroupMembership_GroupVersionKind = CRDGroupVersion.WithKind(UserGroupMembership_Kind)
)

func init() {
	SchemeBuilder.Register(&UserGroupMembership{}, &UserGroupMembershipList{})
}
