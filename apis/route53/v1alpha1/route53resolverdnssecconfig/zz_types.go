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

type Route53ResolverDnssecConfigObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Id string `json:"id" tf:"id"`

	OwnerId string `json:"ownerId" tf:"owner_id"`

	ValidationStatus string `json:"validationStatus" tf:"validation_status"`
}

type Route53ResolverDnssecConfigParameters struct {
	ResourceId string `json:"resourceId" tf:"resource_id"`
}

// Route53ResolverDnssecConfigSpec defines the desired state of Route53ResolverDnssecConfig
type Route53ResolverDnssecConfigSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Route53ResolverDnssecConfigParameters `json:"forProvider"`
}

// Route53ResolverDnssecConfigStatus defines the observed state of Route53ResolverDnssecConfig.
type Route53ResolverDnssecConfigStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Route53ResolverDnssecConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverDnssecConfig is the Schema for the Route53ResolverDnssecConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Route53ResolverDnssecConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverDnssecConfigSpec   `json:"spec"`
	Status            Route53ResolverDnssecConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverDnssecConfigList contains a list of Route53ResolverDnssecConfigs
type Route53ResolverDnssecConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverDnssecConfig `json:"items"`
}

// Repository type metadata.
var (
	Route53ResolverDnssecConfigKind             = "Route53ResolverDnssecConfig"
	Route53ResolverDnssecConfigGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Route53ResolverDnssecConfigKind}.String()
	Route53ResolverDnssecConfigKindAPIVersion   = Route53ResolverDnssecConfigKind + "." + v1alpha1.GroupVersion.String()
	Route53ResolverDnssecConfigGroupVersionKind = v1alpha1.GroupVersion.WithKind(Route53ResolverDnssecConfigKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Route53ResolverDnssecConfig{}, &Route53ResolverDnssecConfigList{})
}
