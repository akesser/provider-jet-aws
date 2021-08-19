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
// +groupName=route53.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/route53/v1alpha1"
)

type Route53ResolverRuleAssociationObservation struct {
}

type Route53ResolverRuleAssociationParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`

	ResolverRuleId string `json:"resolverRuleId" tf:"resolver_rule_id"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

// Route53ResolverRuleAssociationSpec defines the desired state of Route53ResolverRuleAssociation
type Route53ResolverRuleAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Route53ResolverRuleAssociationParameters `json:"forProvider"`
}

// Route53ResolverRuleAssociationStatus defines the observed state of Route53ResolverRuleAssociation.
type Route53ResolverRuleAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Route53ResolverRuleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverRuleAssociation is the Schema for the Route53ResolverRuleAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Route53ResolverRuleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverRuleAssociationSpec   `json:"spec"`
	Status            Route53ResolverRuleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverRuleAssociationList contains a list of Route53ResolverRuleAssociations
type Route53ResolverRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverRuleAssociation `json:"items"`
}

// Repository type metadata.
var (
	Route53ResolverRuleAssociationKind             = "Route53ResolverRuleAssociation"
	Route53ResolverRuleAssociationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Route53ResolverRuleAssociationKind}.String()
	Route53ResolverRuleAssociationKindAPIVersion   = Route53ResolverRuleAssociationKind + "." + v1alpha1.GroupVersion.String()
	Route53ResolverRuleAssociationGroupVersionKind = v1alpha1.GroupVersion.WithKind(Route53ResolverRuleAssociationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Route53ResolverRuleAssociation{}, &Route53ResolverRuleAssociationList{})
}
