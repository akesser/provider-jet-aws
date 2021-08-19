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
// +groupName=apprunner.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/apprunner/v1alpha1"
)

type ApprunnerConnectionObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Status string `json:"status" tf:"status"`
}

type ApprunnerConnectionParameters struct {
	ConnectionName string `json:"connectionName" tf:"connection_name"`

	ProviderType string `json:"providerType" tf:"provider_type"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ApprunnerConnectionSpec defines the desired state of ApprunnerConnection
type ApprunnerConnectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApprunnerConnectionParameters `json:"forProvider"`
}

// ApprunnerConnectionStatus defines the observed state of ApprunnerConnection.
type ApprunnerConnectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApprunnerConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApprunnerConnection is the Schema for the ApprunnerConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ApprunnerConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApprunnerConnectionSpec   `json:"spec"`
	Status            ApprunnerConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApprunnerConnectionList contains a list of ApprunnerConnections
type ApprunnerConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApprunnerConnection `json:"items"`
}

// Repository type metadata.
var (
	ApprunnerConnectionKind             = "ApprunnerConnection"
	ApprunnerConnectionGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ApprunnerConnectionKind}.String()
	ApprunnerConnectionKindAPIVersion   = ApprunnerConnectionKind + "." + v1alpha1.GroupVersion.String()
	ApprunnerConnectionGroupVersionKind = v1alpha1.GroupVersion.WithKind(ApprunnerConnectionKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ApprunnerConnection{}, &ApprunnerConnectionList{})
}
