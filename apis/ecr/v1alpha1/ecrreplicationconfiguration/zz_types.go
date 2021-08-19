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
// +groupName=ecr.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ecr/v1alpha1"
)

type DestinationObservation struct {
}

type DestinationParameters struct {
	Region string `json:"region" tf:"region"`

	RegistryId string `json:"registryId" tf:"registry_id"`
}

type EcrReplicationConfigurationObservation struct {
	RegistryId string `json:"registryId" tf:"registry_id"`
}

type EcrReplicationConfigurationParameters struct {
	ReplicationConfiguration []ReplicationConfigurationParameters `json:"replicationConfiguration,omitempty" tf:"replication_configuration"`
}

type ReplicationConfigurationObservation struct {
}

type ReplicationConfigurationParameters struct {
	Rule []RuleParameters `json:"rule" tf:"rule"`
}

type RuleObservation struct {
}

type RuleParameters struct {
	Destination []DestinationParameters `json:"destination" tf:"destination"`
}

// EcrReplicationConfigurationSpec defines the desired state of EcrReplicationConfiguration
type EcrReplicationConfigurationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EcrReplicationConfigurationParameters `json:"forProvider"`
}

// EcrReplicationConfigurationStatus defines the observed state of EcrReplicationConfiguration.
type EcrReplicationConfigurationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EcrReplicationConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EcrReplicationConfiguration is the Schema for the EcrReplicationConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EcrReplicationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcrReplicationConfigurationSpec   `json:"spec"`
	Status            EcrReplicationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EcrReplicationConfigurationList contains a list of EcrReplicationConfigurations
type EcrReplicationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EcrReplicationConfiguration `json:"items"`
}

// Repository type metadata.
var (
	EcrReplicationConfigurationKind             = "EcrReplicationConfiguration"
	EcrReplicationConfigurationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: EcrReplicationConfigurationKind}.String()
	EcrReplicationConfigurationKindAPIVersion   = EcrReplicationConfigurationKind + "." + v1alpha1.GroupVersion.String()
	EcrReplicationConfigurationGroupVersionKind = v1alpha1.GroupVersion.WithKind(EcrReplicationConfigurationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&EcrReplicationConfiguration{}, &EcrReplicationConfigurationList{})
}
