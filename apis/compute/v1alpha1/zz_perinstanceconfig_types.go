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

type PerInstanceConfigObservation struct {
}

type PerInstanceConfigParameters struct {

	// The instance group manager this instance config is part of.
	// +kubebuilder:validation:Required
	InstanceGroupManager *string `json:"instanceGroupManager" tf:"instance_group_manager,omitempty"`

	// +kubebuilder:validation:Optional
	MinimalAction *string `json:"minimalAction,omitempty" tf:"minimal_action,omitempty"`

	// +kubebuilder:validation:Optional
	MostDisruptiveAllowedAction *string `json:"mostDisruptiveAllowedAction,omitempty" tf:"most_disruptive_allowed_action,omitempty"`

	// The name for this per-instance config and its corresponding instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The preserved state for this instance.
	// +kubebuilder:validation:Optional
	PreservedState []PreservedStateParameters `json:"preservedState,omitempty" tf:"preserved_state,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	RemoveInstanceStateOnDestroy *bool `json:"removeInstanceStateOnDestroy,omitempty" tf:"remove_instance_state_on_destroy,omitempty"`

	// Zone where the containing instance group manager is located
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PreservedStateDiskObservation struct {
}

type PreservedStateDiskParameters struct {

	// A value that prescribes what should happen to the stateful disk when the VM instance is deleted.
	// The available options are 'NEVER' and 'ON_PERMANENT_INSTANCE_DELETION'.
	// 'NEVER' - detach the disk when the VM is deleted, but do not delete the disk.
	// 'ON_PERMANENT_INSTANCE_DELETION' will delete the stateful disk when the VM is permanently
	// deleted from the instance group. Default value: "NEVER" Possible values: ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"]
	// +kubebuilder:validation:Optional
	DeleteRule *string `json:"deleteRule,omitempty" tf:"delete_rule,omitempty"`

	// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.
	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// The mode of the disk. Default value: "READ_WRITE" Possible values: ["READ_ONLY", "READ_WRITE"]
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The URI of an existing persistent disk to attach under the specified device-name in the format
	// 'projects/project-id/zones/zone/disks/disk-name'.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`
}

type PreservedStateObservation struct {
}

type PreservedStateParameters struct {

	// Stateful disks for the instance.
	// +kubebuilder:validation:Optional
	Disk []PreservedStateDiskParameters `json:"disk,omitempty" tf:"disk,omitempty"`

	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

// PerInstanceConfigSpec defines the desired state of PerInstanceConfig
type PerInstanceConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PerInstanceConfigParameters `json:"forProvider"`
}

// PerInstanceConfigStatus defines the observed state of PerInstanceConfig.
type PerInstanceConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PerInstanceConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PerInstanceConfig is the Schema for the PerInstanceConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type PerInstanceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PerInstanceConfigSpec   `json:"spec"`
	Status            PerInstanceConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PerInstanceConfigList contains a list of PerInstanceConfigs
type PerInstanceConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PerInstanceConfig `json:"items"`
}

// Repository type metadata.
var (
	PerInstanceConfig_Kind             = "PerInstanceConfig"
	PerInstanceConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PerInstanceConfig_Kind}.String()
	PerInstanceConfig_KindAPIVersion   = PerInstanceConfig_Kind + "." + CRDGroupVersion.String()
	PerInstanceConfig_GroupVersionKind = CRDGroupVersion.WithKind(PerInstanceConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&PerInstanceConfig{}, &PerInstanceConfigList{})
}
