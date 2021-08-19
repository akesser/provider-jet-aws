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

type S3BucketObjectObservation struct {
	VersionId string `json:"versionId" tf:"version_id"`
}

type S3BucketObjectParameters struct {
	Acl *string `json:"acl,omitempty" tf:"acl"`

	Bucket string `json:"bucket" tf:"bucket"`

	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled"`

	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control"`

	Content *string `json:"content,omitempty" tf:"content"`

	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64"`

	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition"`

	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding"`

	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language"`

	ContentType *string `json:"contentType,omitempty" tf:"content_type"`

	Etag *string `json:"etag,omitempty" tf:"etag"`

	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`

	Key string `json:"key" tf:"key"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata"`

	ObjectLockLegalHoldStatus *string `json:"objectLockLegalHoldStatus,omitempty" tf:"object_lock_legal_hold_status"`

	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode"`

	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date"`

	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption"`

	Source *string `json:"source,omitempty" tf:"source"`

	SourceHash *string `json:"sourceHash,omitempty" tf:"source_hash"`

	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect"`
}

// S3BucketObjectSpec defines the desired state of S3BucketObject
type S3BucketObjectSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       S3BucketObjectParameters `json:"forProvider"`
}

// S3BucketObjectStatus defines the observed state of S3BucketObject.
type S3BucketObjectStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          S3BucketObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketObject is the Schema for the S3BucketObjects API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type S3BucketObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketObjectSpec   `json:"spec"`
	Status            S3BucketObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketObjectList contains a list of S3BucketObjects
type S3BucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketObject `json:"items"`
}

// Repository type metadata.
var (
	S3BucketObjectKind             = "S3BucketObject"
	S3BucketObjectGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: S3BucketObjectKind}.String()
	S3BucketObjectKindAPIVersion   = S3BucketObjectKind + "." + v1alpha1.GroupVersion.String()
	S3BucketObjectGroupVersionKind = v1alpha1.GroupVersion.WithKind(S3BucketObjectKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&S3BucketObject{}, &S3BucketObjectList{})
}
