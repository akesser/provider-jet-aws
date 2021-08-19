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
// +groupName=sns.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/sns/v1alpha1"
)

type SnsPlatformApplicationObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SnsPlatformApplicationParameters struct {
	EventDeliveryFailureTopicArn *string `json:"eventDeliveryFailureTopicArn,omitempty" tf:"event_delivery_failure_topic_arn"`

	EventEndpointCreatedTopicArn *string `json:"eventEndpointCreatedTopicArn,omitempty" tf:"event_endpoint_created_topic_arn"`

	EventEndpointDeletedTopicArn *string `json:"eventEndpointDeletedTopicArn,omitempty" tf:"event_endpoint_deleted_topic_arn"`

	EventEndpointUpdatedTopicArn *string `json:"eventEndpointUpdatedTopicArn,omitempty" tf:"event_endpoint_updated_topic_arn"`

	FailureFeedbackRoleArn *string `json:"failureFeedbackRoleArn,omitempty" tf:"failure_feedback_role_arn"`

	Name string `json:"name" tf:"name"`

	Platform string `json:"platform" tf:"platform"`

	PlatformCredential string `json:"platformCredential" tf:"platform_credential"`

	PlatformPrincipal *string `json:"platformPrincipal,omitempty" tf:"platform_principal"`

	SuccessFeedbackRoleArn *string `json:"successFeedbackRoleArn,omitempty" tf:"success_feedback_role_arn"`

	SuccessFeedbackSampleRate *string `json:"successFeedbackSampleRate,omitempty" tf:"success_feedback_sample_rate"`
}

// SnsPlatformApplicationSpec defines the desired state of SnsPlatformApplication
type SnsPlatformApplicationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SnsPlatformApplicationParameters `json:"forProvider"`
}

// SnsPlatformApplicationStatus defines the observed state of SnsPlatformApplication.
type SnsPlatformApplicationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SnsPlatformApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnsPlatformApplication is the Schema for the SnsPlatformApplications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SnsPlatformApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnsPlatformApplicationSpec   `json:"spec"`
	Status            SnsPlatformApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnsPlatformApplicationList contains a list of SnsPlatformApplications
type SnsPlatformApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnsPlatformApplication `json:"items"`
}

// Repository type metadata.
var (
	SnsPlatformApplicationKind             = "SnsPlatformApplication"
	SnsPlatformApplicationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SnsPlatformApplicationKind}.String()
	SnsPlatformApplicationKindAPIVersion   = SnsPlatformApplicationKind + "." + v1alpha1.GroupVersion.String()
	SnsPlatformApplicationGroupVersionKind = v1alpha1.GroupVersion.WithKind(SnsPlatformApplicationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SnsPlatformApplication{}, &SnsPlatformApplicationList{})
}
