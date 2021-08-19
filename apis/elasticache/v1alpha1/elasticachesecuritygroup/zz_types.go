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
// +groupName=elasticache.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/elasticache/v1alpha1"
)

type ElasticacheSecurityGroupObservation struct {
}

type ElasticacheSecurityGroupParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	SecurityGroupNames []string `json:"securityGroupNames" tf:"security_group_names"`
}

// ElasticacheSecurityGroupSpec defines the desired state of ElasticacheSecurityGroup
type ElasticacheSecurityGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElasticacheSecurityGroupParameters `json:"forProvider"`
}

// ElasticacheSecurityGroupStatus defines the observed state of ElasticacheSecurityGroup.
type ElasticacheSecurityGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElasticacheSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheSecurityGroup is the Schema for the ElasticacheSecurityGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ElasticacheSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheSecurityGroupSpec   `json:"spec"`
	Status            ElasticacheSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheSecurityGroupList contains a list of ElasticacheSecurityGroups
type ElasticacheSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticacheSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	ElasticacheSecurityGroupKind             = "ElasticacheSecurityGroup"
	ElasticacheSecurityGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ElasticacheSecurityGroupKind}.String()
	ElasticacheSecurityGroupKindAPIVersion   = ElasticacheSecurityGroupKind + "." + v1alpha1.GroupVersion.String()
	ElasticacheSecurityGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(ElasticacheSecurityGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ElasticacheSecurityGroup{}, &ElasticacheSecurityGroupList{})
}
