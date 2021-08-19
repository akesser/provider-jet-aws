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
// +groupName=codestarconnections.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/codestarconnections/v1alpha1"
)

type CodestarconnectionsHostObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Status string `json:"status" tf:"status"`
}

type CodestarconnectionsHostParameters struct {
	Name string `json:"name" tf:"name"`

	ProviderEndpoint string `json:"providerEndpoint" tf:"provider_endpoint"`

	ProviderType string `json:"providerType" tf:"provider_type"`

	VpcConfiguration []VpcConfigurationParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration"`
}

type VpcConfigurationObservation struct {
}

type VpcConfigurationParameters struct {
	SecurityGroupIds []string `json:"securityGroupIds" tf:"security_group_ids"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	TlsCertificate *string `json:"tlsCertificate,omitempty" tf:"tls_certificate"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

// CodestarconnectionsHostSpec defines the desired state of CodestarconnectionsHost
type CodestarconnectionsHostSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodestarconnectionsHostParameters `json:"forProvider"`
}

// CodestarconnectionsHostStatus defines the observed state of CodestarconnectionsHost.
type CodestarconnectionsHostStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodestarconnectionsHostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodestarconnectionsHost is the Schema for the CodestarconnectionsHosts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CodestarconnectionsHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodestarconnectionsHostSpec   `json:"spec"`
	Status            CodestarconnectionsHostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodestarconnectionsHostList contains a list of CodestarconnectionsHosts
type CodestarconnectionsHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodestarconnectionsHost `json:"items"`
}

// Repository type metadata.
var (
	CodestarconnectionsHostKind             = "CodestarconnectionsHost"
	CodestarconnectionsHostGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CodestarconnectionsHostKind}.String()
	CodestarconnectionsHostKindAPIVersion   = CodestarconnectionsHostKind + "." + v1alpha1.GroupVersion.String()
	CodestarconnectionsHostGroupVersionKind = v1alpha1.GroupVersion.WithKind(CodestarconnectionsHostKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CodestarconnectionsHost{}, &CodestarconnectionsHostList{})
}
