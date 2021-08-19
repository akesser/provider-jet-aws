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
// +groupName=dx.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/dx/v1alpha1"
)

type DxHostedPrivateVirtualInterfaceAccepterObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type DxHostedPrivateVirtualInterfaceAccepterParameters struct {
	DxGatewayId *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VirtualInterfaceId string `json:"virtualInterfaceId" tf:"virtual_interface_id"`

	VpnGatewayId *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id"`
}

// DxHostedPrivateVirtualInterfaceAccepterSpec defines the desired state of DxHostedPrivateVirtualInterfaceAccepter
type DxHostedPrivateVirtualInterfaceAccepterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DxHostedPrivateVirtualInterfaceAccepterParameters `json:"forProvider"`
}

// DxHostedPrivateVirtualInterfaceAccepterStatus defines the observed state of DxHostedPrivateVirtualInterfaceAccepter.
type DxHostedPrivateVirtualInterfaceAccepterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DxHostedPrivateVirtualInterfaceAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedPrivateVirtualInterfaceAccepter is the Schema for the DxHostedPrivateVirtualInterfaceAccepters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DxHostedPrivateVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedPrivateVirtualInterfaceAccepterSpec   `json:"spec"`
	Status            DxHostedPrivateVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedPrivateVirtualInterfaceAccepterList contains a list of DxHostedPrivateVirtualInterfaceAccepters
type DxHostedPrivateVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxHostedPrivateVirtualInterfaceAccepter `json:"items"`
}

// Repository type metadata.
var (
	DxHostedPrivateVirtualInterfaceAccepterKind             = "DxHostedPrivateVirtualInterfaceAccepter"
	DxHostedPrivateVirtualInterfaceAccepterGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DxHostedPrivateVirtualInterfaceAccepterKind}.String()
	DxHostedPrivateVirtualInterfaceAccepterKindAPIVersion   = DxHostedPrivateVirtualInterfaceAccepterKind + "." + v1alpha1.GroupVersion.String()
	DxHostedPrivateVirtualInterfaceAccepterGroupVersionKind = v1alpha1.GroupVersion.WithKind(DxHostedPrivateVirtualInterfaceAccepterKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DxHostedPrivateVirtualInterfaceAccepter{}, &DxHostedPrivateVirtualInterfaceAccepterList{})
}
