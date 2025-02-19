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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CxEnvironmentObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type CxEnvironmentParameters struct {

	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// The Agent to create an Environment for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
	// +kubebuilder:validation:Required
	VersionConfigs []VersionConfigsParameters `json:"versionConfigs" tf:"version_configs,omitempty"`
}

type VersionConfigsObservation struct {
}

type VersionConfigsParameters struct {

	// Format: projects/{{project}}/locations/{{location}}/agents/{{agent}}/flows/{{flow}}/versions/{{version}}.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

// CxEnvironmentSpec defines the desired state of CxEnvironment
type CxEnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CxEnvironmentParameters `json:"forProvider"`
}

// CxEnvironmentStatus defines the observed state of CxEnvironment.
type CxEnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CxEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CxEnvironment is the Schema for the CxEnvironments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type CxEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CxEnvironmentSpec   `json:"spec"`
	Status            CxEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CxEnvironmentList contains a list of CxEnvironments
type CxEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CxEnvironment `json:"items"`
}

// Repository type metadata.
var (
	CxEnvironment_Kind             = "CxEnvironment"
	CxEnvironment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CxEnvironment_Kind}.String()
	CxEnvironment_KindAPIVersion   = CxEnvironment_Kind + "." + CRDGroupVersion.String()
	CxEnvironment_GroupVersionKind = CRDGroupVersion.WithKind(CxEnvironment_Kind)
)

func init() {
	SchemeBuilder.Register(&CxEnvironment{}, &CxEnvironmentList{})
}
