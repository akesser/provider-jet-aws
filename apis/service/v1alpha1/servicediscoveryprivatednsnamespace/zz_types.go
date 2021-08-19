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
// +groupName=service.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/service/v1alpha1"
)

type ServiceDiscoveryPrivateDnsNamespaceObservation struct {
	Arn string `json:"arn" tf:"arn"`

	HostedZone string `json:"hostedZone" tf:"hosted_zone"`
}

type ServiceDiscoveryPrivateDnsNamespaceParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Vpc string `json:"vpc" tf:"vpc"`
}

// ServiceDiscoveryPrivateDnsNamespaceSpec defines the desired state of ServiceDiscoveryPrivateDnsNamespace
type ServiceDiscoveryPrivateDnsNamespaceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServiceDiscoveryPrivateDnsNamespaceParameters `json:"forProvider"`
}

// ServiceDiscoveryPrivateDnsNamespaceStatus defines the observed state of ServiceDiscoveryPrivateDnsNamespace.
type ServiceDiscoveryPrivateDnsNamespaceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServiceDiscoveryPrivateDnsNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceDiscoveryPrivateDnsNamespace is the Schema for the ServiceDiscoveryPrivateDnsNamespaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServiceDiscoveryPrivateDnsNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceDiscoveryPrivateDnsNamespaceSpec   `json:"spec"`
	Status            ServiceDiscoveryPrivateDnsNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceDiscoveryPrivateDnsNamespaceList contains a list of ServiceDiscoveryPrivateDnsNamespaces
type ServiceDiscoveryPrivateDnsNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDiscoveryPrivateDnsNamespace `json:"items"`
}

// Repository type metadata.
var (
	ServiceDiscoveryPrivateDnsNamespaceKind             = "ServiceDiscoveryPrivateDnsNamespace"
	ServiceDiscoveryPrivateDnsNamespaceGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ServiceDiscoveryPrivateDnsNamespaceKind}.String()
	ServiceDiscoveryPrivateDnsNamespaceKindAPIVersion   = ServiceDiscoveryPrivateDnsNamespaceKind + "." + v1alpha1.GroupVersion.String()
	ServiceDiscoveryPrivateDnsNamespaceGroupVersionKind = v1alpha1.GroupVersion.WithKind(ServiceDiscoveryPrivateDnsNamespaceKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ServiceDiscoveryPrivateDnsNamespace{}, &ServiceDiscoveryPrivateDnsNamespaceList{})
}
