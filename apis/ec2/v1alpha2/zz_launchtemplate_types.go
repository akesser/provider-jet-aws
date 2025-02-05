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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AcceleratorCountObservation struct {
}

type AcceleratorCountParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type AcceleratorTotalMemoryMibObservation struct {
}

type AcceleratorTotalMemoryMibParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type BaselineEBSBandwidthMbpsObservation struct {
}

type BaselineEBSBandwidthMbpsParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type BlockDeviceMappingsObservation struct {
}

type BlockDeviceMappingsParameters struct {

	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	EBS []EBSParameters `json:"ebs,omitempty" tf:"ebs,omitempty"`

	// +kubebuilder:validation:Optional
	NoDevice *string `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type CPUOptionsObservation struct {
}

type CPUOptionsParameters struct {

	// +kubebuilder:validation:Optional
	CoreCount *float64 `json:"coreCount,omitempty" tf:"core_count,omitempty"`

	// +kubebuilder:validation:Optional
	ThreadsPerCore *float64 `json:"threadsPerCore,omitempty" tf:"threads_per_core,omitempty"`
}

type CapacityReservationSpecificationCapacityReservationTargetObservation struct {
}

type CapacityReservationSpecificationCapacityReservationTargetParameters struct {

	// +kubebuilder:validation:Optional
	CapacityReservationID *string `json:"capacityReservationId,omitempty" tf:"capacity_reservation_id,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityReservationResourceGroupArn *string `json:"capacityReservationResourceGroupArn,omitempty" tf:"capacity_reservation_resource_group_arn,omitempty"`
}

type EBSObservation struct {
}

type EBSParameters struct {

	// +kubebuilder:validation:Optional
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +crossplane:generate:reference:type=github.com/dkb-bank/provider-jet-aws/apis/kms/v1alpha2.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type ElasticGpuSpecificationsObservation struct {
}

type ElasticGpuSpecificationsParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ElasticInferenceAcceleratorObservation struct {
}

type ElasticInferenceAcceleratorParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type HibernationOptionsObservation struct {
}

type HibernationOptionsParameters struct {

	// +kubebuilder:validation:Required
	Configured *bool `json:"configured" tf:"configured,omitempty"`
}

type IAMInstanceProfileObservation struct {
}

type IAMInstanceProfileParameters struct {

	// +crossplane:generate:reference:type=github.com/dkb-bank/provider-jet-aws/apis/iam/v1alpha2.InstanceProfile
	// +crossplane:generate:reference:extractor=github.com/dkb-bank/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/dkb-bank/provider-jet-aws/apis/iam/v1alpha2.InstanceProfile
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`
}

type InstanceMarketOptionsObservation struct {
}

type InstanceMarketOptionsParameters struct {

	// +kubebuilder:validation:Optional
	MarketType *string `json:"marketType,omitempty" tf:"market_type,omitempty"`

	// +kubebuilder:validation:Optional
	SpotOptions []SpotOptionsParameters `json:"spotOptions,omitempty" tf:"spot_options,omitempty"`
}

type InstanceRequirementsObservation struct {
}

type InstanceRequirementsParameters struct {

	// +kubebuilder:validation:Optional
	AcceleratorCount []AcceleratorCountParameters `json:"acceleratorCount,omitempty" tf:"accelerator_count,omitempty"`

	// +kubebuilder:validation:Optional
	AcceleratorManufacturers []*string `json:"acceleratorManufacturers,omitempty" tf:"accelerator_manufacturers,omitempty"`

	// +kubebuilder:validation:Optional
	AcceleratorNames []*string `json:"acceleratorNames,omitempty" tf:"accelerator_names,omitempty"`

	// +kubebuilder:validation:Optional
	AcceleratorTotalMemoryMib []AcceleratorTotalMemoryMibParameters `json:"acceleratorTotalMemoryMib,omitempty" tf:"accelerator_total_memory_mib,omitempty"`

	// +kubebuilder:validation:Optional
	AcceleratorTypes []*string `json:"acceleratorTypes,omitempty" tf:"accelerator_types,omitempty"`

	// +kubebuilder:validation:Optional
	BareMetal *string `json:"bareMetal,omitempty" tf:"bare_metal,omitempty"`

	// +kubebuilder:validation:Optional
	BaselineEBSBandwidthMbps []BaselineEBSBandwidthMbpsParameters `json:"baselineEbsBandwidthMbps,omitempty" tf:"baseline_ebs_bandwidth_mbps,omitempty"`

	// +kubebuilder:validation:Optional
	BurstablePerformance *string `json:"burstablePerformance,omitempty" tf:"burstable_performance,omitempty"`

	// +kubebuilder:validation:Optional
	CPUManufacturers []*string `json:"cpuManufacturers,omitempty" tf:"cpu_manufacturers,omitempty"`

	// +kubebuilder:validation:Optional
	ExcludedInstanceTypes []*string `json:"excludedInstanceTypes,omitempty" tf:"excluded_instance_types,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceGenerations []*string `json:"instanceGenerations,omitempty" tf:"instance_generations,omitempty"`

	// +kubebuilder:validation:Optional
	LocalStorage *string `json:"localStorage,omitempty" tf:"local_storage,omitempty"`

	// +kubebuilder:validation:Optional
	LocalStorageTypes []*string `json:"localStorageTypes,omitempty" tf:"local_storage_types,omitempty"`

	// +kubebuilder:validation:Optional
	MemoryGibPerVcpu []MemoryGibPerVcpuParameters `json:"memoryGibPerVcpu,omitempty" tf:"memory_gib_per_vcpu,omitempty"`

	// +kubebuilder:validation:Required
	MemoryMib []MemoryMibParameters `json:"memoryMib" tf:"memory_mib,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceCount []NetworkInterfaceCountParameters `json:"networkInterfaceCount,omitempty" tf:"network_interface_count,omitempty"`

	// +kubebuilder:validation:Optional
	OnDemandMaxPricePercentageOverLowestPrice *float64 `json:"onDemandMaxPricePercentageOverLowestPrice,omitempty" tf:"on_demand_max_price_percentage_over_lowest_price,omitempty"`

	// +kubebuilder:validation:Optional
	RequireHibernateSupport *bool `json:"requireHibernateSupport,omitempty" tf:"require_hibernate_support,omitempty"`

	// +kubebuilder:validation:Optional
	SpotMaxPricePercentageOverLowestPrice *float64 `json:"spotMaxPricePercentageOverLowestPrice,omitempty" tf:"spot_max_price_percentage_over_lowest_price,omitempty"`

	// +kubebuilder:validation:Optional
	TotalLocalStorageGb []TotalLocalStorageGbParameters `json:"totalLocalStorageGb,omitempty" tf:"total_local_storage_gb,omitempty"`

	// +kubebuilder:validation:Required
	VcpuCount []VcpuCountParameters `json:"vcpuCount" tf:"vcpu_count,omitempty"`
}

type LaunchTemplateCapacityReservationSpecificationObservation struct {
}

type LaunchTemplateCapacityReservationSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	CapacityReservationPreference *string `json:"capacityReservationPreference,omitempty" tf:"capacity_reservation_preference,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityReservationTarget []CapacityReservationSpecificationCapacityReservationTargetParameters `json:"capacityReservationTarget,omitempty" tf:"capacity_reservation_target,omitempty"`
}

type LaunchTemplateCreditSpecificationObservation struct {
}

type LaunchTemplateCreditSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	CPUCredits *string `json:"cpuCredits,omitempty" tf:"cpu_credits,omitempty"`
}

type LaunchTemplateEnclaveOptionsObservation struct {
}

type LaunchTemplateEnclaveOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type LaunchTemplateMaintenanceOptionsObservation struct {
}

type LaunchTemplateMaintenanceOptionsParameters struct {

	// +kubebuilder:validation:Optional
	AutoRecovery *string `json:"autoRecovery,omitempty" tf:"auto_recovery,omitempty"`
}

type LaunchTemplateMetadataOptionsObservation struct {
}

type LaunchTemplateMetadataOptionsParameters struct {

	// +kubebuilder:validation:Optional
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPProtocolIPv6 *string `json:"httpProtocolIpv6,omitempty" tf:"http_protocol_ipv6,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPPutResponseHopLimit *float64 `json:"httpPutResponseHopLimit,omitempty" tf:"http_put_response_hop_limit,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPTokens *string `json:"httpTokens,omitempty" tf:"http_tokens,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceMetadataTags *string `json:"instanceMetadataTags,omitempty" tf:"instance_metadata_tags,omitempty"`
}

type LaunchTemplateObservation_2 struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LatestVersion *float64 `json:"latestVersion,omitempty" tf:"latest_version,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LaunchTemplateParameters_2 struct {

	// +kubebuilder:validation:Optional
	BlockDeviceMappings []BlockDeviceMappingsParameters `json:"blockDeviceMappings,omitempty" tf:"block_device_mappings,omitempty"`

	// +kubebuilder:validation:Optional
	CPUOptions []CPUOptionsParameters `json:"cpuOptions,omitempty" tf:"cpu_options,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityReservationSpecification []LaunchTemplateCapacityReservationSpecificationParameters `json:"capacityReservationSpecification,omitempty" tf:"capacity_reservation_specification,omitempty"`

	// +kubebuilder:validation:Optional
	CreditSpecification []LaunchTemplateCreditSpecificationParameters `json:"creditSpecification,omitempty" tf:"credit_specification,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultVersion *float64 `json:"defaultVersion,omitempty" tf:"default_version,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DisableAPIStop *bool `json:"disableApiStop,omitempty" tf:"disable_api_stop,omitempty"`

	// +kubebuilder:validation:Optional
	DisableAPITermination *bool `json:"disableApiTermination,omitempty" tf:"disable_api_termination,omitempty"`

	// +kubebuilder:validation:Optional
	EBSOptimized *string `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticGpuSpecifications []ElasticGpuSpecificationsParameters `json:"elasticGpuSpecifications,omitempty" tf:"elastic_gpu_specifications,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticInferenceAccelerator []ElasticInferenceAcceleratorParameters `json:"elasticInferenceAccelerator,omitempty" tf:"elastic_inference_accelerator,omitempty"`

	// +kubebuilder:validation:Optional
	EnclaveOptions []LaunchTemplateEnclaveOptionsParameters `json:"enclaveOptions,omitempty" tf:"enclave_options,omitempty"`

	// +kubebuilder:validation:Optional
	HibernationOptions []HibernationOptionsParameters `json:"hibernationOptions,omitempty" tf:"hibernation_options,omitempty"`

	// +kubebuilder:validation:Optional
	IAMInstanceProfile []IAMInstanceProfileParameters `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`

	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceInitiatedShutdownBehavior *string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceMarketOptions []InstanceMarketOptionsParameters `json:"instanceMarketOptions,omitempty" tf:"instance_market_options,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceRequirements []InstanceRequirementsParameters `json:"instanceRequirements,omitempty" tf:"instance_requirements,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	KernelID *string `json:"kernelId,omitempty" tf:"kernel_id,omitempty"`

	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// +kubebuilder:validation:Optional
	LicenseSpecification []LicenseSpecificationParameters `json:"licenseSpecification,omitempty" tf:"license_specification,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceOptions []LaunchTemplateMaintenanceOptionsParameters `json:"maintenanceOptions,omitempty" tf:"maintenance_options,omitempty"`

	// +kubebuilder:validation:Optional
	MetadataOptions []LaunchTemplateMetadataOptionsParameters `json:"metadataOptions,omitempty" tf:"metadata_options,omitempty"`

	// +kubebuilder:validation:Optional
	Monitoring []MonitoringParameters `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterfaces []NetworkInterfacesParameters `json:"networkInterfaces,omitempty" tf:"network_interfaces,omitempty"`

	// +kubebuilder:validation:Optional
	Placement []PlacementParameters `json:"placement,omitempty" tf:"placement,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateDNSNameOptions []LaunchTemplatePrivateDNSNameOptionsParameters `json:"privateDnsNameOptions,omitempty" tf:"private_dns_name_options,omitempty"`

	// +kubebuilder:validation:Optional
	RAMDiskID *string `json:"ramDiskId,omitempty" tf:"ram_disk_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupNameRefs []v1.Reference `json:"securityGroupNameRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupNameSelector *v1.Selector `json:"securityGroupNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupNameRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupNameSelector
	// +kubebuilder:validation:Optional
	SecurityGroupNames []*string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`

	// +kubebuilder:validation:Optional
	TagSpecifications []TagSpecificationsParameters `json:"tagSpecifications,omitempty" tf:"tag_specifications,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	UpdateDefaultVersion *bool `json:"updateDefaultVersion,omitempty" tf:"update_default_version,omitempty"`

	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:refFieldName=VpcSecurityGroupIdRefs
	// +crossplane:generate:reference:selectorFieldName=VpcSecurityGroupIdSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	VpcSecurityGroupIdRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcSecurityGroupIdSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`
}

type LaunchTemplatePrivateDNSNameOptionsObservation struct {
}

type LaunchTemplatePrivateDNSNameOptionsParameters struct {

	// +kubebuilder:validation:Optional
	EnableResourceNameDNSARecord *bool `json:"enableResourceNameDnsARecord,omitempty" tf:"enable_resource_name_dns_a_record,omitempty"`

	// +kubebuilder:validation:Optional
	EnableResourceNameDNSAaaaRecord *bool `json:"enableResourceNameDnsAaaaRecord,omitempty" tf:"enable_resource_name_dns_aaaa_record,omitempty"`

	// +kubebuilder:validation:Optional
	HostnameType *string `json:"hostnameType,omitempty" tf:"hostname_type,omitempty"`
}

type LicenseSpecificationObservation struct {
}

type LicenseSpecificationParameters struct {

	// +kubebuilder:validation:Required
	LicenseConfigurationArn *string `json:"licenseConfigurationArn" tf:"license_configuration_arn,omitempty"`
}

type MemoryGibPerVcpuObservation struct {
}

type MemoryGibPerVcpuParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type MemoryMibObservation struct {
}

type MemoryMibParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type MonitoringObservation struct {
}

type MonitoringParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type NetworkInterfaceCountObservation struct {
}

type NetworkInterfaceCountParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type NetworkInterfacesObservation struct {
}

type NetworkInterfacesParameters struct {

	// +kubebuilder:validation:Optional
	AssociateCarrierIPAddress *string `json:"associateCarrierIpAddress,omitempty" tf:"associate_carrier_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	AssociatePublicIPAddress *string `json:"associatePublicIpAddress,omitempty" tf:"associate_public_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceIndex *float64 `json:"deviceIndex,omitempty" tf:"device_index,omitempty"`

	// +kubebuilder:validation:Optional
	IPv4AddressCount *float64 `json:"ipv4AddressCount,omitempty" tf:"ipv4_address_count,omitempty"`

	// +kubebuilder:validation:Optional
	IPv4Addresses []*string `json:"ipv4Addresses,omitempty" tf:"ipv4_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	IPv4PrefixCount *float64 `json:"ipv4PrefixCount,omitempty" tf:"ipv4_prefix_count,omitempty"`

	// +kubebuilder:validation:Optional
	IPv4Prefixes []*string `json:"ipv4Prefixes,omitempty" tf:"ipv4_prefixes,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6AddressCount *float64 `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Addresses []*string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6PrefixCount *float64 `json:"ipv6PrefixCount,omitempty" tf:"ipv6_prefix_count,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Prefixes []*string `json:"ipv6Prefixes,omitempty" tf:"ipv6_prefixes,omitempty"`

	// +kubebuilder:validation:Optional
	InterfaceType *string `json:"interfaceType,omitempty" tf:"interface_type,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkCardIndex *float64 `json:"networkCardIndex,omitempty" tf:"network_card_index,omitempty"`

	// +crossplane:generate:reference:type=NetworkInterface
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type PlacementObservation struct {
}

type PlacementParameters struct {

	// +kubebuilder:validation:Optional
	Affinity *string `json:"affinity,omitempty" tf:"affinity,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	// +kubebuilder:validation:Optional
	HostID *string `json:"hostId,omitempty" tf:"host_id,omitempty"`

	// +kubebuilder:validation:Optional
	HostResourceGroupArn *string `json:"hostResourceGroupArn,omitempty" tf:"host_resource_group_arn,omitempty"`

	// +kubebuilder:validation:Optional
	PartitionNumber *float64 `json:"partitionNumber,omitempty" tf:"partition_number,omitempty"`

	// +kubebuilder:validation:Optional
	SpreadDomain *string `json:"spreadDomain,omitempty" tf:"spread_domain,omitempty"`

	// +kubebuilder:validation:Optional
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`
}

type SpotOptionsObservation struct {
}

type SpotOptionsParameters struct {

	// +kubebuilder:validation:Optional
	BlockDurationMinutes *float64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceInterruptionBehavior *string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	MaxPrice *string `json:"maxPrice,omitempty" tf:"max_price,omitempty"`

	// +kubebuilder:validation:Optional
	SpotInstanceType *string `json:"spotInstanceType,omitempty" tf:"spot_instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
}

type TagSpecificationsObservation struct {
}

type TagSpecificationsParameters struct {

	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TotalLocalStorageGbObservation struct {
}

type TotalLocalStorageGbParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type VcpuCountObservation struct {
}

type VcpuCountParameters struct {

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

// LaunchTemplateSpec defines the desired state of LaunchTemplate
type LaunchTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LaunchTemplateParameters_2 `json:"forProvider"`
}

// LaunchTemplateStatus defines the observed state of LaunchTemplate.
type LaunchTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LaunchTemplateObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LaunchTemplate is the Schema for the LaunchTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LaunchTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LaunchTemplateSpec   `json:"spec"`
	Status            LaunchTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LaunchTemplateList contains a list of LaunchTemplates
type LaunchTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LaunchTemplate `json:"items"`
}

// Repository type metadata.
var (
	LaunchTemplate_Kind             = "LaunchTemplate"
	LaunchTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LaunchTemplate_Kind}.String()
	LaunchTemplate_KindAPIVersion   = LaunchTemplate_Kind + "." + CRDGroupVersion.String()
	LaunchTemplate_GroupVersionKind = CRDGroupVersion.WithKind(LaunchTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&LaunchTemplate{}, &LaunchTemplateList{})
}
