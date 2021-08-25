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

package v1alpha1

import "github.com/crossplane-contrib/terrajet/pkg/conversion"

// GetTerraformResourceType returns Terraform resource type for this GuarddutyOrganizationAdminAccount
func (mg *GuarddutyOrganizationAdminAccount) GetTerraformResourceType() string {
	return "aws_guardduty_organization_admin_account"
}

// GetTerraformResourceIdField returns Terraform identifier field for this GuarddutyOrganizationAdminAccount
func (tr *GuarddutyOrganizationAdminAccount) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this GuarddutyOrganizationAdminAccount
func (tr *GuarddutyOrganizationAdminAccount) GetObservation() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Status.AtProvider)
}

// SetObservation for this GuarddutyOrganizationAdminAccount
func (tr *GuarddutyOrganizationAdminAccount) SetObservation(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Status.AtProvider)
}

// GetParameters of this GuarddutyOrganizationAdminAccount
func (tr *GuarddutyOrganizationAdminAccount) GetParameters() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Spec.ForProvider)
}

// SetParameters for this GuarddutyOrganizationAdminAccount
func (tr *GuarddutyOrganizationAdminAccount) SetParameters(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Spec.ForProvider)
}
