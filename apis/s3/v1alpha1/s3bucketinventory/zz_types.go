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
// +groupName=s3.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/s3/v1alpha1"
)

type BucketObservation struct {
}

type BucketParameters struct {
	AccountId *string `json:"accountId,omitempty" tf:"account_id"`

	BucketArn string `json:"bucketArn" tf:"bucket_arn"`

	Encryption []EncryptionParameters `json:"encryption,omitempty" tf:"encryption"`

	Format string `json:"format" tf:"format"`

	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

type DestinationObservation struct {
}

type DestinationParameters struct {
	Bucket []BucketParameters `json:"bucket" tf:"bucket"`
}

type EncryptionObservation struct {
}

type EncryptionParameters struct {
	SseKms []SseKmsParameters `json:"sseKms,omitempty" tf:"sse_kms"`

	SseS3 []SseS3Parameters `json:"sseS3,omitempty" tf:"sse_s3"`
}

type FilterObservation struct {
}

type FilterParameters struct {
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

type S3BucketInventoryObservation struct {
}

type S3BucketInventoryParameters struct {
	Bucket string `json:"bucket" tf:"bucket"`

	Destination []DestinationParameters `json:"destination" tf:"destination"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	Filter []FilterParameters `json:"filter,omitempty" tf:"filter"`

	IncludedObjectVersions string `json:"includedObjectVersions" tf:"included_object_versions"`

	Name string `json:"name" tf:"name"`

	OptionalFields []string `json:"optionalFields,omitempty" tf:"optional_fields"`

	Schedule []ScheduleParameters `json:"schedule" tf:"schedule"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {
	Frequency string `json:"frequency" tf:"frequency"`
}

type SseKmsObservation struct {
}

type SseKmsParameters struct {
	KeyId string `json:"keyId" tf:"key_id"`
}

type SseS3Observation struct {
}

type SseS3Parameters struct {
}

// S3BucketInventorySpec defines the desired state of S3BucketInventory
type S3BucketInventorySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       S3BucketInventoryParameters `json:"forProvider"`
}

// S3BucketInventoryStatus defines the observed state of S3BucketInventory.
type S3BucketInventoryStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          S3BucketInventoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketInventory is the Schema for the S3BucketInventorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type S3BucketInventory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketInventorySpec   `json:"spec"`
	Status            S3BucketInventoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketInventoryList contains a list of S3BucketInventorys
type S3BucketInventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketInventory `json:"items"`
}

// Repository type metadata.
var (
	S3BucketInventoryKind             = "S3BucketInventory"
	S3BucketInventoryGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: S3BucketInventoryKind}.String()
	S3BucketInventoryKindAPIVersion   = S3BucketInventoryKind + "." + v1alpha1.GroupVersion.String()
	S3BucketInventoryGroupVersionKind = v1alpha1.GroupVersion.WithKind(S3BucketInventoryKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&S3BucketInventory{}, &S3BucketInventoryList{})
}
