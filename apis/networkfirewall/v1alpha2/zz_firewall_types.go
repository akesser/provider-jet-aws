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

type AttachmentObservation struct {
	EndpointID *string `json:"endpointId,omitempty" tf:"endpoint_id,omitempty"`

	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type AttachmentParameters struct {
}

type FirewallObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	FirewallStatus []FirewallStatusObservation `json:"firewallStatus,omitempty" tf:"firewall_status,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	UpdateToken *string `json:"updateToken,omitempty" tf:"update_token,omitempty"`
}

type FirewallParameters struct {

	// +kubebuilder:validation:Optional
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=FirewallPolicy
	// +kubebuilder:validation:Optional
	FirewallPolicyArn *string `json:"firewallPolicyArn,omitempty" tf:"firewall_policy_arn,omitempty"`

	// +kubebuilder:validation:Optional
	FirewallPolicyArnRef *v1.Reference `json:"firewallPolicyArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FirewallPolicyArnSelector *v1.Selector `json:"firewallPolicyArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FirewallPolicyChangeProtection *bool `json:"firewallPolicyChangeProtection,omitempty" tf:"firewall_policy_change_protection,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetChangeProtection *bool `json:"subnetChangeProtection,omitempty" tf:"subnet_change_protection,omitempty"`

	// +kubebuilder:validation:Required
	SubnetMapping []SubnetMappingParameters `json:"subnetMapping" tf:"subnet_mapping,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=github.com/dkb-bank/provider-jet-aws/apis/ec2/v1alpha2.VPC
	// +crossplane:generate:reference:refFieldName=VpcIdRef
	// +crossplane:generate:reference:selectorFieldName=VpcIdSelector
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcIdRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcIdSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type FirewallStatusObservation struct {
	SyncStates []SyncStatesObservation `json:"syncStates,omitempty" tf:"sync_states,omitempty"`
}

type FirewallStatusParameters struct {
}

type SubnetMappingObservation struct {
}

type SubnetMappingParameters struct {

	// +crossplane:generate:reference:type=github.com/dkb-bank/provider-jet-aws/apis/ec2/v1alpha2.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIdRef
	// +crossplane:generate:reference:selectorFieldName=SubnetIdSelector
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIdRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIdSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type SyncStatesObservation struct {
	Attachment []AttachmentObservation `json:"attachment,omitempty" tf:"attachment,omitempty"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
}

type SyncStatesParameters struct {
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallParameters `json:"forProvider"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Firewall is the Schema for the Firewalls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec"`
	Status            FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	Firewall_Kind             = "Firewall"
	Firewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Firewall_Kind}.String()
	Firewall_KindAPIVersion   = Firewall_Kind + "." + CRDGroupVersion.String()
	Firewall_GroupVersionKind = CRDGroupVersion.WithKind(Firewall_Kind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
