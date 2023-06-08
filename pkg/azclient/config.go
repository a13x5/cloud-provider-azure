/*
Copyright 2023 The Kubernetes Authors.

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

package azclient

import (
	"errors"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/policy/ratelimit"
)

type ClientFactoryConfig struct {
	ratelimit.CloudProviderRateLimitConfig
	ARMClientConfig

	// Enable exponential backoff to manage resource request retries
	CloudProviderBackoff bool `json:"cloudProviderBackoff,omitempty" yaml:"cloudProviderBackoff,omitempty"`
}

type ARMClientConfig struct {
	// The cloud environment identifier. Takes values from https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azcore@v1.6.0/cloud
	Cloud string `json:"cloud,omitempty" yaml:"cloud,omitempty"`
	// The user agent for Azure customer usage attribution
	UserAgent string `json:"userAgent,omitempty" yaml:"userAgent,omitempty"`
	// ResourceManagerEndpoint is the cloud's resource manager endpoint. If set, cloud provider queries this endpoint
	// in order to generate an autorest.Environment instance instead of using one of the pre-defined Environments.
	ResourceManagerEndpoint string `json:"resourceManagerEndpoint,omitempty" yaml:"resourceManagerEndpoint,omitempty"`
}

var EnvironmentMapping = map[string]cloud.Configuration{
	"AZURECHINACLOUD":        cloud.AzureChina,
	"AZURECLOUD":             cloud.AzurePublic,
	"AZUREPUBLICCLOUD":       cloud.AzurePublic,
	"AZUREUSGOVERNMENT":      cloud.AzureGovernment,
	"AZUREUSGOVERNMENTCLOUD": cloud.AzureGovernment, //TODO: deprecate
}

var ConfigNotFoundError = errors.New("resource manager service config is not found")

func AzureCloudConfigFromName(cloudName string, endpoint string) (cloud.Configuration, error) {
	cloudName = strings.ToUpper(cloudName)
	var config *cloud.Configuration
	if cloudName == "" {
		cloudName = "AZUREPUBLICCLOUD"
	}
	if cloudConfig, ok := EnvironmentMapping[cloudName]; ok {
		config = &cloudConfig
	}
	if endpoint != "" {
		if config == nil {
			//if cloudName is customized profile, ActiveDirectoryAuthorityHost is not set.
			//todo: load ActiveDirectoryAuthorityHost from azure.json
			return cloud.Configuration{}, ConfigNotFoundError
		}
		var serviceConfig cloud.ServiceConfiguration
		var ok bool
		if serviceConfig, ok = config.Services[cloud.ResourceManager]; !ok {
			//todo: load cloud.ResourceManager.Audience from azure.json
			return cloud.Configuration{}, ConfigNotFoundError
		}
		serviceConfig.Endpoint = endpoint
		config.Services[cloud.ResourceManager] = serviceConfig
	}
	//todo: load endpoint config from file.
	if config == nil {
		return cloud.Configuration{}, ConfigNotFoundError
	}
	return *config, nil
}
