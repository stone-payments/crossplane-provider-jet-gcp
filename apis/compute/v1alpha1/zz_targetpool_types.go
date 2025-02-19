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

type TargetPoolObservation struct {
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type TargetPoolParameters struct {

	// URL to the backup target pool. Must also set failover_ratio.
	// +kubebuilder:validation:Optional
	BackupPool *string `json:"backupPool,omitempty" tf:"backup_pool,omitempty"`

	// Textual description field.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Ratio (0 to 1) of failed nodes before using the backup pool (which must also be set).
	// +kubebuilder:validation:Optional
	FailoverRatio *float64 `json:"failoverRatio,omitempty" tf:"failover_ratio,omitempty"`

	// List of zero or one health check name or self_link. Only legacy google_compute_http_health_check is supported.
	// +kubebuilder:validation:Optional
	HealthChecks []*string `json:"healthChecks,omitempty" tf:"health_checks,omitempty"`

	// List of instances in the pool. They can be given as URLs, or in the form of "zone/name". Note that the instances need not exist at the time of target pool creation, so there is no need to use the Terraform interpolators to create a dependency on the instances from the target pool.
	// +kubebuilder:validation:Optional
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// A unique name for the resource, required by GCE. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Where the target pool resides. Defaults to project region.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// How to distribute load. Options are "NONE" (no affinity). "CLIENT_IP" (hash of the source/dest addresses / ports), and "CLIENT_IP_PROTO" also includes the protocol (default "NONE").
	// +kubebuilder:validation:Optional
	SessionAffinity *string `json:"sessionAffinity,omitempty" tf:"session_affinity,omitempty"`
}

// TargetPoolSpec defines the desired state of TargetPool
type TargetPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetPoolParameters `json:"forProvider"`
}

// TargetPoolStatus defines the observed state of TargetPool.
type TargetPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TargetPool is the Schema for the TargetPools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type TargetPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetPoolSpec   `json:"spec"`
	Status            TargetPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetPoolList contains a list of TargetPools
type TargetPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetPool `json:"items"`
}

// Repository type metadata.
var (
	TargetPool_Kind             = "TargetPool"
	TargetPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetPool_Kind}.String()
	TargetPool_KindAPIVersion   = TargetPool_Kind + "." + CRDGroupVersion.String()
	TargetPool_GroupVersionKind = CRDGroupVersion.WithKind(TargetPool_Kind)
)

func init() {
	SchemeBuilder.Register(&TargetPool{}, &TargetPoolList{})
}
