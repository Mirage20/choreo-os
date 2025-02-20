/*
 * Copyright (c) 2025, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	SchemeBuilder.Register(&Component{}, &ComponentList{})
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=comp;comps

// Component is the Schema for the components API.
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComponentSpec   `json:"spec,omitempty"`
	Status ComponentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComponentList contains a list of Component.
type ComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Component `json:"items"`
}

// ComponentSpec defines the desired state of Component.
type ComponentSpec struct {
	// Type of the component that indicates how the component deployed.
	Type ComponentType `json:"type,omitempty"`
	// Source the source information of the component where the code or image is retrieved.
	Source ComponentSource `json:"source,omitempty"`
}

// ComponentStatus defines the observed state of Component.
type ComponentStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// ComponentType defines how the component is deployed.
type ComponentType string

const (
	ComponentTypeService        ComponentType = "Service"
	ComponentTypeManualTask     ComponentType = "ManualTask"
	ComponentTypeScheduledTask  ComponentType = "ScheduledTask"
	ComponentTypeWebApplication ComponentType = "WebApplication"
	ComponentTypeWebhook        ComponentType = "Webhook"
	ComponentTypeAPIProxy       ComponentType = "APIProxy"
	ComponentTypeTestRunner     ComponentType = "TestRunner"
	ComponentTypeEventHandler   ComponentType = "EventHandler"
)

// ComponentSource defines the source information of the component where the code or image is retrieved.
type ComponentSource struct {
	// GitRepository specifies the configuration for the component source to be a Git repository indicating
	// that the component should be built from the source code.
	// This field is mutually exclusive with the other source types.
	GitRepository *GitRepository `json:"gitRepository,omitempty"`

	// ContainerRegistry specifies the configuration for the component source to be a container image indicating
	// that the component should be deployed using the provided image.
	// This field is mutually exclusive with the other source types.
	ContainerRegistry *ContainerRegistry `json:"containerRegistry,omitempty"`
}

// GitRepository defines the Git repository configuration
type GitRepository struct {
	// URL the Git repository URL
	// Examples:
	// - https://github.com/jhonb2077/customer-service
	// - https://gitlab.com/jhonb2077/customer-service
	URL string `json:"url"`

	// Authentication the authentication information to access the Git repository
	// If not provided, the Git repository should be public
	Authentication GitAuthentication `json:"authentication,omitempty"`
}

// GitAuthentication defines the authentication configuration for Git
type GitAuthentication struct {
	// SecretRef is a reference to the secret containing Git credentials
	SecretRef string `json:"secretRef"`
}

// ContainerRegistry defines the container registry configuration.
type ContainerRegistry struct {
	// Image name of the container image. Format: <registry>/<image> without the tag.
	// Example: docker.io/library/nginx
	ImageName string `json:"imageName,omitempty"`
	// Authentication information to access the container registry.
	Authentication *RegistryAuthentication `json:"authentication,omitempty"`
}

// RegistryAuthentication defines the authentication configuration for container registry
type RegistryAuthentication struct {
	// Reference to the secret that contains the container registry authentication info.
	SecretRef string `json:"secretRef,omitempty"`
}
