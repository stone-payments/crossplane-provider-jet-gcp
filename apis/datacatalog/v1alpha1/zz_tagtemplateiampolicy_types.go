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

type TagTemplateIamPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type TagTemplateIamPolicyParameters struct {

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	TagTemplate *string `json:"tagTemplate" tf:"tag_template,omitempty"`
}

// TagTemplateIamPolicySpec defines the desired state of TagTemplateIamPolicy
type TagTemplateIamPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagTemplateIamPolicyParameters `json:"forProvider"`
}

// TagTemplateIamPolicyStatus defines the observed state of TagTemplateIamPolicy.
type TagTemplateIamPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagTemplateIamPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TagTemplateIamPolicy is the Schema for the TagTemplateIamPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type TagTemplateIamPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagTemplateIamPolicySpec   `json:"spec"`
	Status            TagTemplateIamPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagTemplateIamPolicyList contains a list of TagTemplateIamPolicys
type TagTemplateIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagTemplateIamPolicy `json:"items"`
}

// Repository type metadata.
var (
	TagTemplateIamPolicy_Kind             = "TagTemplateIamPolicy"
	TagTemplateIamPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TagTemplateIamPolicy_Kind}.String()
	TagTemplateIamPolicy_KindAPIVersion   = TagTemplateIamPolicy_Kind + "." + CRDGroupVersion.String()
	TagTemplateIamPolicy_GroupVersionKind = CRDGroupVersion.WithKind(TagTemplateIamPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&TagTemplateIamPolicy{}, &TagTemplateIamPolicyList{})
}
