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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GroupPolicyAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GroupPolicyAttachmentParameters struct {

	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// +kubebuilder:validation:Optional
	GroupRef *v1.Reference `json:"groupRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	GroupSelector *v1.Selector `json:"groupSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Policy
	// +crossplane:generate:reference:extractor=github.com/dkb-bank/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyArnRef *v1.Reference `json:"policyArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PolicyArnSelector *v1.Selector `json:"policyArnSelector,omitempty" tf:"-"`
}

// GroupPolicyAttachmentSpec defines the desired state of GroupPolicyAttachment
type GroupPolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupPolicyAttachmentParameters `json:"forProvider"`
}

// GroupPolicyAttachmentStatus defines the observed state of GroupPolicyAttachment.
type GroupPolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPolicyAttachment is the Schema for the GroupPolicyAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type GroupPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupPolicyAttachmentSpec   `json:"spec"`
	Status            GroupPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPolicyAttachmentList contains a list of GroupPolicyAttachments
type GroupPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	GroupPolicyAttachment_Kind             = "GroupPolicyAttachment"
	GroupPolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupPolicyAttachment_Kind}.String()
	GroupPolicyAttachment_KindAPIVersion   = GroupPolicyAttachment_Kind + "." + CRDGroupVersion.String()
	GroupPolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(GroupPolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupPolicyAttachment{}, &GroupPolicyAttachmentList{})
}
