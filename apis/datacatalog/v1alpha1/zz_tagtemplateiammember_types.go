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

type TagTemplateIamMemberConditionObservation struct {
}

type TagTemplateIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type TagTemplateIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type TagTemplateIamMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []TagTemplateIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	TagTemplate *string `json:"tagTemplate" tf:"tag_template,omitempty"`
}

// TagTemplateIamMemberSpec defines the desired state of TagTemplateIamMember
type TagTemplateIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagTemplateIamMemberParameters `json:"forProvider"`
}

// TagTemplateIamMemberStatus defines the observed state of TagTemplateIamMember.
type TagTemplateIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagTemplateIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TagTemplateIamMember is the Schema for the TagTemplateIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type TagTemplateIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagTemplateIamMemberSpec   `json:"spec"`
	Status            TagTemplateIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagTemplateIamMemberList contains a list of TagTemplateIamMembers
type TagTemplateIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagTemplateIamMember `json:"items"`
}

// Repository type metadata.
var (
	TagTemplateIamMember_Kind             = "TagTemplateIamMember"
	TagTemplateIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TagTemplateIamMember_Kind}.String()
	TagTemplateIamMember_KindAPIVersion   = TagTemplateIamMember_Kind + "." + CRDGroupVersion.String()
	TagTemplateIamMember_GroupVersionKind = CRDGroupVersion.WithKind(TagTemplateIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&TagTemplateIamMember{}, &TagTemplateIamMemberList{})
}