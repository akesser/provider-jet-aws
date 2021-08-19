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
// +groupName=lightsail.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/lightsail/v1alpha1"
)

type LightsailInstancePublicPortsObservation struct {
}

type LightsailInstancePublicPortsParameters struct {
	InstanceName string `json:"instanceName" tf:"instance_name"`

	PortInfo []PortInfoParameters `json:"portInfo" tf:"port_info"`
}

type PortInfoObservation struct {
}

type PortInfoParameters struct {
	Cidrs []string `json:"cidrs,omitempty" tf:"cidrs"`

	FromPort int64 `json:"fromPort" tf:"from_port"`

	Protocol string `json:"protocol" tf:"protocol"`

	ToPort int64 `json:"toPort" tf:"to_port"`
}

// LightsailInstancePublicPortsSpec defines the desired state of LightsailInstancePublicPorts
type LightsailInstancePublicPortsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LightsailInstancePublicPortsParameters `json:"forProvider"`
}

// LightsailInstancePublicPortsStatus defines the observed state of LightsailInstancePublicPorts.
type LightsailInstancePublicPortsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LightsailInstancePublicPortsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LightsailInstancePublicPorts is the Schema for the LightsailInstancePublicPortss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LightsailInstancePublicPorts struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LightsailInstancePublicPortsSpec   `json:"spec"`
	Status            LightsailInstancePublicPortsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LightsailInstancePublicPortsList contains a list of LightsailInstancePublicPortss
type LightsailInstancePublicPortsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LightsailInstancePublicPorts `json:"items"`
}

// Repository type metadata.
var (
	LightsailInstancePublicPortsKind             = "LightsailInstancePublicPorts"
	LightsailInstancePublicPortsGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: LightsailInstancePublicPortsKind}.String()
	LightsailInstancePublicPortsKindAPIVersion   = LightsailInstancePublicPortsKind + "." + v1alpha1.GroupVersion.String()
	LightsailInstancePublicPortsGroupVersionKind = v1alpha1.GroupVersion.WithKind(LightsailInstancePublicPortsKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&LightsailInstancePublicPorts{}, &LightsailInstancePublicPortsList{})
}
