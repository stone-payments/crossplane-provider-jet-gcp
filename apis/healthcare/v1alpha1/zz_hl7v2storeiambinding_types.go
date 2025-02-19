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

type Hl7V2StoreIamBindingConditionObservation struct {
}

type Hl7V2StoreIamBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type Hl7V2StoreIamBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type Hl7V2StoreIamBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []Hl7V2StoreIamBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Hl7V2StoreID *string `json:"hl7V2StoreId" tf:"hl7_v2_store_id,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// Hl7V2StoreIamBindingSpec defines the desired state of Hl7V2StoreIamBinding
type Hl7V2StoreIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Hl7V2StoreIamBindingParameters `json:"forProvider"`
}

// Hl7V2StoreIamBindingStatus defines the observed state of Hl7V2StoreIamBinding.
type Hl7V2StoreIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Hl7V2StoreIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Hl7V2StoreIamBinding is the Schema for the Hl7V2StoreIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Hl7V2StoreIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Hl7V2StoreIamBindingSpec   `json:"spec"`
	Status            Hl7V2StoreIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Hl7V2StoreIamBindingList contains a list of Hl7V2StoreIamBindings
type Hl7V2StoreIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hl7V2StoreIamBinding `json:"items"`
}

// Repository type metadata.
var (
	Hl7V2StoreIamBinding_Kind             = "Hl7V2StoreIamBinding"
	Hl7V2StoreIamBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Hl7V2StoreIamBinding_Kind}.String()
	Hl7V2StoreIamBinding_KindAPIVersion   = Hl7V2StoreIamBinding_Kind + "." + CRDGroupVersion.String()
	Hl7V2StoreIamBinding_GroupVersionKind = CRDGroupVersion.WithKind(Hl7V2StoreIamBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&Hl7V2StoreIamBinding{}, &Hl7V2StoreIamBindingList{})
}
