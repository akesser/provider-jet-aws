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
// +groupName=elb.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/elb/v1alpha1"
)

type AccessLogsObservation struct {
}

type AccessLogsParameters struct {
	Bucket string `json:"bucket" tf:"bucket"`

	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	Interval *int64 `json:"interval,omitempty" tf:"interval"`
}

type ElbObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DnsName string `json:"dnsName" tf:"dns_name"`

	SourceSecurityGroupId string `json:"sourceSecurityGroupId" tf:"source_security_group_id"`

	ZoneId string `json:"zoneId" tf:"zone_id"`
}

type ElbParameters struct {
	AccessLogs []AccessLogsParameters `json:"accessLogs,omitempty" tf:"access_logs"`

	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones"`

	ConnectionDraining *bool `json:"connectionDraining,omitempty" tf:"connection_draining"`

	ConnectionDrainingTimeout *int64 `json:"connectionDrainingTimeout,omitempty" tf:"connection_draining_timeout"`

	CrossZoneLoadBalancing *bool `json:"crossZoneLoadBalancing,omitempty" tf:"cross_zone_load_balancing"`

	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check"`

	IdleTimeout *int64 `json:"idleTimeout,omitempty" tf:"idle_timeout"`

	Instances []string `json:"instances,omitempty" tf:"instances"`

	Internal *bool `json:"internal,omitempty" tf:"internal"`

	Listener []ListenerParameters `json:"listener" tf:"listener"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	SourceSecurityGroup *string `json:"sourceSecurityGroup,omitempty" tf:"source_security_group"`

	Subnets []string `json:"subnets,omitempty" tf:"subnets"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type HealthCheckObservation struct {
}

type HealthCheckParameters struct {
	HealthyThreshold int64 `json:"healthyThreshold" tf:"healthy_threshold"`

	Interval int64 `json:"interval" tf:"interval"`

	Target string `json:"target" tf:"target"`

	Timeout int64 `json:"timeout" tf:"timeout"`

	UnhealthyThreshold int64 `json:"unhealthyThreshold" tf:"unhealthy_threshold"`
}

type ListenerObservation struct {
}

type ListenerParameters struct {
	InstancePort int64 `json:"instancePort" tf:"instance_port"`

	InstanceProtocol string `json:"instanceProtocol" tf:"instance_protocol"`

	LbPort int64 `json:"lbPort" tf:"lb_port"`

	LbProtocol string `json:"lbProtocol" tf:"lb_protocol"`

	SslCertificateId *string `json:"sslCertificateId,omitempty" tf:"ssl_certificate_id"`
}

// ElbSpec defines the desired state of Elb
type ElbSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElbParameters `json:"forProvider"`
}

// ElbStatus defines the observed state of Elb.
type ElbStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElbObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Elb is the Schema for the Elbs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Elb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElbSpec   `json:"spec"`
	Status            ElbStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElbList contains a list of Elbs
type ElbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Elb `json:"items"`
}

// Repository type metadata.
var (
	ElbKind             = "Elb"
	ElbGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ElbKind}.String()
	ElbKindAPIVersion   = ElbKind + "." + v1alpha1.GroupVersion.String()
	ElbGroupVersionKind = v1alpha1.GroupVersion.WithKind(ElbKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Elb{}, &ElbList{})
}
