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

type ConditionObservation struct {
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type FunctionIamBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type FunctionIamBindingParameters struct {

	// +kubebuilder:validation:Required
	CloudFunction *string `json:"cloudFunction" tf:"cloud_function,omitempty"`

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// FunctionIamBindingSpec defines the desired state of FunctionIamBinding
type FunctionIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionIamBindingParameters `json:"forProvider"`
}

// FunctionIamBindingStatus defines the observed state of FunctionIamBinding.
type FunctionIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionIamBinding is the Schema for the FunctionIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type FunctionIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionIamBindingSpec   `json:"spec"`
	Status            FunctionIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionIamBindingList contains a list of FunctionIamBindings
type FunctionIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionIamBinding `json:"items"`
}

// Repository type metadata.
var (
	FunctionIamBinding_Kind             = "FunctionIamBinding"
	FunctionIamBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionIamBinding_Kind}.String()
	FunctionIamBinding_KindAPIVersion   = FunctionIamBinding_Kind + "." + CRDGroupVersion.String()
	FunctionIamBinding_GroupVersionKind = CRDGroupVersion.WithKind(FunctionIamBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionIamBinding{}, &FunctionIamBindingList{})
}
