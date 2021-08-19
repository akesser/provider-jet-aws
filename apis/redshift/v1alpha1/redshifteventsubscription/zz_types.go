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
// +groupName=redshift.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/redshift/v1alpha1"
)

type RedshiftEventSubscriptionObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CustomerAwsId string `json:"customerAwsId" tf:"customer_aws_id"`

	Status string `json:"status" tf:"status"`
}

type RedshiftEventSubscriptionParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	EventCategories []string `json:"eventCategories,omitempty" tf:"event_categories"`

	Name string `json:"name" tf:"name"`

	Severity *string `json:"severity,omitempty" tf:"severity"`

	SnsTopicArn string `json:"snsTopicArn" tf:"sns_topic_arn"`

	SourceIds []string `json:"sourceIds,omitempty" tf:"source_ids"`

	SourceType *string `json:"sourceType,omitempty" tf:"source_type"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// RedshiftEventSubscriptionSpec defines the desired state of RedshiftEventSubscription
type RedshiftEventSubscriptionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RedshiftEventSubscriptionParameters `json:"forProvider"`
}

// RedshiftEventSubscriptionStatus defines the observed state of RedshiftEventSubscription.
type RedshiftEventSubscriptionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RedshiftEventSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftEventSubscription is the Schema for the RedshiftEventSubscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RedshiftEventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftEventSubscriptionSpec   `json:"spec"`
	Status            RedshiftEventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftEventSubscriptionList contains a list of RedshiftEventSubscriptions
type RedshiftEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedshiftEventSubscription `json:"items"`
}

// Repository type metadata.
var (
	RedshiftEventSubscriptionKind             = "RedshiftEventSubscription"
	RedshiftEventSubscriptionGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: RedshiftEventSubscriptionKind}.String()
	RedshiftEventSubscriptionKindAPIVersion   = RedshiftEventSubscriptionKind + "." + v1alpha1.GroupVersion.String()
	RedshiftEventSubscriptionGroupVersionKind = v1alpha1.GroupVersion.WithKind(RedshiftEventSubscriptionKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&RedshiftEventSubscription{}, &RedshiftEventSubscriptionList{})
}
