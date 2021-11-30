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

type WebBackendServiceIamMemberConditionObservation struct {
}

type WebBackendServiceIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type WebBackendServiceIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type WebBackendServiceIamMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []WebBackendServiceIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	WebBackendService *string `json:"webBackendService" tf:"web_backend_service,omitempty"`
}

// WebBackendServiceIamMemberSpec defines the desired state of WebBackendServiceIamMember
type WebBackendServiceIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebBackendServiceIamMemberParameters `json:"forProvider"`
}

// WebBackendServiceIamMemberStatus defines the observed state of WebBackendServiceIamMember.
type WebBackendServiceIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebBackendServiceIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebBackendServiceIamMember is the Schema for the WebBackendServiceIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type WebBackendServiceIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebBackendServiceIamMemberSpec   `json:"spec"`
	Status            WebBackendServiceIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebBackendServiceIamMemberList contains a list of WebBackendServiceIamMembers
type WebBackendServiceIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebBackendServiceIamMember `json:"items"`
}

// Repository type metadata.
var (
	WebBackendServiceIamMember_Kind             = "WebBackendServiceIamMember"
	WebBackendServiceIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebBackendServiceIamMember_Kind}.String()
	WebBackendServiceIamMember_KindAPIVersion   = WebBackendServiceIamMember_Kind + "." + CRDGroupVersion.String()
	WebBackendServiceIamMember_GroupVersionKind = CRDGroupVersion.WithKind(WebBackendServiceIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&WebBackendServiceIamMember{}, &WebBackendServiceIamMemberList{})
}