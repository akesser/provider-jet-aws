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

package lakeformationpermissions

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane-contrib/terrajet/pkg/terraform"
	"github.com/crossplane/crossplane-runtime/pkg/logging"

	lakeformationpermissions "github.com/crossplane-contrib/provider-tf-aws/apis/lakeformation/v1alpha1/lakeformationpermissions"
	clients "github.com/crossplane-contrib/provider-tf-aws/internal/clients"
)

// Setup adds a controller that reconciles LakeformationPermissions managed resources.
func Setup(mgr ctrl.Manager, l logging.Logger) error {
	return terraform.SetupController(mgr, l, &lakeformationpermissions.LakeformationPermissions{}, lakeformationpermissions.LakeformationPermissionsGroupVersionKind, clients.ProviderConfigBuilder)
}
