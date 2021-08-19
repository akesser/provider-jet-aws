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
// +groupName=fms.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/fms/v1alpha1"
)

type FmsAdminAccountObservation struct {
}

type FmsAdminAccountParameters struct {
	AccountId *string `json:"accountId,omitempty" tf:"account_id"`
}

// FmsAdminAccountSpec defines the desired state of FmsAdminAccount
type FmsAdminAccountSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FmsAdminAccountParameters `json:"forProvider"`
}

// FmsAdminAccountStatus defines the observed state of FmsAdminAccount.
type FmsAdminAccountStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FmsAdminAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FmsAdminAccount is the Schema for the FmsAdminAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FmsAdminAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FmsAdminAccountSpec   `json:"spec"`
	Status            FmsAdminAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FmsAdminAccountList contains a list of FmsAdminAccounts
type FmsAdminAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FmsAdminAccount `json:"items"`
}

// Repository type metadata.
var (
	FmsAdminAccountKind             = "FmsAdminAccount"
	FmsAdminAccountGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: FmsAdminAccountKind}.String()
	FmsAdminAccountKindAPIVersion   = FmsAdminAccountKind + "." + v1alpha1.GroupVersion.String()
	FmsAdminAccountGroupVersionKind = v1alpha1.GroupVersion.WithKind(FmsAdminAccountKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&FmsAdminAccount{}, &FmsAdminAccountList{})
}
