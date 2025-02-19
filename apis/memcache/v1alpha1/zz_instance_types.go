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

type InstanceObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	DiscoveryEndpoint *string `json:"discoveryEndpoint,omitempty" tf:"discovery_endpoint,omitempty"`

	MemcacheFullVersion *string `json:"memcacheFullVersion,omitempty" tf:"memcache_full_version,omitempty"`

	MemcacheNodes []MemcacheNodesObservation `json:"memcacheNodes,omitempty" tf:"memcache_nodes,omitempty"`
}

type InstanceParameters struct {

	// The full name of the GCE network to connect the instance to.  If not provided,
	// 'default' will be used.
	// +kubebuilder:validation:Optional
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty"`

	// A user-visible name for the instance.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Resource labels to represent user-provided metadata.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// User-specified parameters for this memcache instance.
	// +kubebuilder:validation:Optional
	MemcacheParameters []MemcacheParametersParameters `json:"memcacheParameters,omitempty" tf:"memcache_parameters,omitempty"`

	// The major version of Memcached software. If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version. Default value: "MEMCACHE_1_5" Possible values: ["MEMCACHE_1_5"]
	// +kubebuilder:validation:Optional
	MemcacheVersion *string `json:"memcacheVersion,omitempty" tf:"memcache_version,omitempty"`

	// The resource name of the instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Configuration for memcache nodes.
	// +kubebuilder:validation:Required
	NodeConfig []NodeConfigParameters `json:"nodeConfig" tf:"node_config,omitempty"`

	// Number of nodes in the memcache instance.
	// +kubebuilder:validation:Required
	NodeCount *int64 `json:"nodeCount" tf:"node_count,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the Memcache instance. If it is not provided, the provider region is used.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Zones where memcache nodes should be provisioned.  If not
	// provided, all zones will be used.
	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type MemcacheNodesObservation struct {
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	NodeID *string `json:"nodeId,omitempty" tf:"node_id,omitempty"`

	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type MemcacheNodesParameters struct {
}

type MemcacheParametersObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MemcacheParametersParameters struct {

	// User-defined set of parameters to use in the memcache process.
	// +kubebuilder:validation:Optional
	Params map[string]*string `json:"params,omitempty" tf:"params,omitempty"`
}

type NodeConfigObservation struct {
}

type NodeConfigParameters struct {

	// Number of CPUs per node.
	// +kubebuilder:validation:Required
	CPUCount *int64 `json:"cpuCount" tf:"cpu_count,omitempty"`

	// Memory size in Mebibytes for each memcache node.
	// +kubebuilder:validation:Required
	MemorySizeMb *int64 `json:"memorySizeMb" tf:"memory_size_mb,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
