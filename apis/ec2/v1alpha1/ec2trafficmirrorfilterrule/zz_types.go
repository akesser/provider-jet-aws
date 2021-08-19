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
// +groupName=ec2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1"
)

type DestinationPortRangeObservation struct {
}

type DestinationPortRangeParameters struct {
	FromPort *int64 `json:"fromPort,omitempty" tf:"from_port"`

	ToPort *int64 `json:"toPort,omitempty" tf:"to_port"`
}

type Ec2TrafficMirrorFilterRuleObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type Ec2TrafficMirrorFilterRuleParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	DestinationCidrBlock string `json:"destinationCidrBlock" tf:"destination_cidr_block"`

	DestinationPortRange []DestinationPortRangeParameters `json:"destinationPortRange,omitempty" tf:"destination_port_range"`

	Protocol *int64 `json:"protocol,omitempty" tf:"protocol"`

	RuleAction string `json:"ruleAction" tf:"rule_action"`

	RuleNumber int64 `json:"ruleNumber" tf:"rule_number"`

	SourceCidrBlock string `json:"sourceCidrBlock" tf:"source_cidr_block"`

	SourcePortRange []SourcePortRangeParameters `json:"sourcePortRange,omitempty" tf:"source_port_range"`

	TrafficDirection string `json:"trafficDirection" tf:"traffic_direction"`

	TrafficMirrorFilterId string `json:"trafficMirrorFilterId" tf:"traffic_mirror_filter_id"`
}

type SourcePortRangeObservation struct {
}

type SourcePortRangeParameters struct {
	FromPort *int64 `json:"fromPort,omitempty" tf:"from_port"`

	ToPort *int64 `json:"toPort,omitempty" tf:"to_port"`
}

// Ec2TrafficMirrorFilterRuleSpec defines the desired state of Ec2TrafficMirrorFilterRule
type Ec2TrafficMirrorFilterRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2TrafficMirrorFilterRuleParameters `json:"forProvider"`
}

// Ec2TrafficMirrorFilterRuleStatus defines the observed state of Ec2TrafficMirrorFilterRule.
type Ec2TrafficMirrorFilterRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2TrafficMirrorFilterRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TrafficMirrorFilterRule is the Schema for the Ec2TrafficMirrorFilterRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Ec2TrafficMirrorFilterRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TrafficMirrorFilterRuleSpec   `json:"spec"`
	Status            Ec2TrafficMirrorFilterRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TrafficMirrorFilterRuleList contains a list of Ec2TrafficMirrorFilterRules
type Ec2TrafficMirrorFilterRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TrafficMirrorFilterRule `json:"items"`
}

// Repository type metadata.
var (
	Ec2TrafficMirrorFilterRuleKind             = "Ec2TrafficMirrorFilterRule"
	Ec2TrafficMirrorFilterRuleGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Ec2TrafficMirrorFilterRuleKind}.String()
	Ec2TrafficMirrorFilterRuleKindAPIVersion   = Ec2TrafficMirrorFilterRuleKind + "." + v1alpha1.GroupVersion.String()
	Ec2TrafficMirrorFilterRuleGroupVersionKind = v1alpha1.GroupVersion.WithKind(Ec2TrafficMirrorFilterRuleKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Ec2TrafficMirrorFilterRule{}, &Ec2TrafficMirrorFilterRuleList{})
}
