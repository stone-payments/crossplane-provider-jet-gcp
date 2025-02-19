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

type DiskIamPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type DiskIamPolicyParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// DiskIamPolicySpec defines the desired state of DiskIamPolicy
type DiskIamPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskIamPolicyParameters `json:"forProvider"`
}

// DiskIamPolicyStatus defines the observed state of DiskIamPolicy.
type DiskIamPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskIamPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiskIamPolicy is the Schema for the DiskIamPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type DiskIamPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskIamPolicySpec   `json:"spec"`
	Status            DiskIamPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskIamPolicyList contains a list of DiskIamPolicys
type DiskIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskIamPolicy `json:"items"`
}

// Repository type metadata.
var (
	DiskIamPolicy_Kind             = "DiskIamPolicy"
	DiskIamPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskIamPolicy_Kind}.String()
	DiskIamPolicy_KindAPIVersion   = DiskIamPolicy_Kind + "." + CRDGroupVersion.String()
	DiskIamPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DiskIamPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskIamPolicy{}, &DiskIamPolicyList{})
}
