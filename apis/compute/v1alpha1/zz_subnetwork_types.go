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

type SecondaryIPRangeObservation struct {
}

type SecondaryIPRangeParameters struct {

	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	// +kubebuilder:validation:Required
	IPCidrRange *string `json:"ipCidrRange" tf:"ip_cidr_range,omitempty"`

	// The name associated with this subnetwork secondary range, used
	// when adding an alias IP range to a VM instance. The name must
	// be 1-63 characters long, and comply with RFC1035. The name
	// must be unique within the subnetwork.
	// +kubebuilder:validation:Required
	RangeName *string `json:"rangeName" tf:"range_name,omitempty"`
}

type SubnetworkLogConfigObservation struct {
}

type SubnetworkLogConfigParameters struct {

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// Toggles the aggregation interval for collecting flow logs. Increasing the
	// interval time will reduce the amount of generated flow logs for long
	// lasting connections. Default is an interval of 5 seconds per connection. Default value: "INTERVAL_5_SEC" Possible values: ["INTERVAL_5_SEC", "INTERVAL_30_SEC", "INTERVAL_1_MIN", "INTERVAL_5_MIN", "INTERVAL_10_MIN", "INTERVAL_15_MIN"]
	// +kubebuilder:validation:Optional
	AggregationInterval *string `json:"aggregationInterval,omitempty" tf:"aggregation_interval,omitempty"`

	// Export filter used to define which VPC flow logs should be logged, as as CEL expression. See
	// https://cloud.google.com/vpc/docs/flow-logs#filtering for details on how to format this field.
	// The default value is 'true', which evaluates to include everything.
	// +kubebuilder:validation:Optional
	FilterExpr *string `json:"filterExpr,omitempty" tf:"filter_expr,omitempty"`

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// The value of the field must be in [0, 1]. Set the sampling rate of VPC
	// flow logs within the subnetwork where 1.0 means all collected logs are
	// reported and 0.0 means no logs are reported. Default is 0.5 which means
	// half of all collected logs are reported.
	// +kubebuilder:validation:Optional
	FlowSampling *float64 `json:"flowSampling,omitempty" tf:"flow_sampling,omitempty"`

	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// Configures whether metadata fields should be added to the reported VPC
	// flow logs. Default value: "INCLUDE_ALL_METADATA" Possible values: ["EXCLUDE_ALL_METADATA", "INCLUDE_ALL_METADATA", "CUSTOM_METADATA"]
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// List of metadata fields that should be added to reported logs.
	// Can only be specified if VPC flow logs for this subnetwork is enabled and "metadata" is set to CUSTOM_METADATA.
	// +kubebuilder:validation:Optional
	MetadataFields []*string `json:"metadataFields,omitempty" tf:"metadata_fields,omitempty"`
}

type SubnetworkObservation_2 struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	ExternalIPv6Prefix *string `json:"externalIpv6Prefix,omitempty" tf:"external_ipv6_prefix,omitempty"`

	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	GatewayAddress *string `json:"gatewayAddress,omitempty" tf:"gateway_address,omitempty"`

	IPv6CidrRange *string `json:"ipv6CidrRange,omitempty" tf:"ipv6_cidr_range,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type SubnetworkParameters_2 struct {

	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The range of internal addresses that are owned by this subnetwork.
	// Provide this property when you create the subnetwork. For example,
	// 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
	// non-overlapping within a network. Only IPv4 is supported.
	// +kubebuilder:validation:Required
	IPCidrRange *string `json:"ipCidrRange" tf:"ip_cidr_range,omitempty"`

	// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation
	// or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet
	// cannot enable direct path. Possible values: ["EXTERNAL"]
	// +kubebuilder:validation:Optional
	IPv6AccessType *string `json:"ipv6AccessType,omitempty" tf:"ipv6_access_type,omitempty"`

	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the 'purpose' of this
	// subnetwork is 'INTERNAL_HTTPS_LOAD_BALANCER'
	// +kubebuilder:validation:Optional
	LogConfig []SubnetworkLogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	// +kubebuilder:validation:Optional
	PrivateIPGoogleAccess *bool `json:"privateIpGoogleAccess,omitempty" tf:"private_ip_google_access,omitempty"`

	// The private IPv6 google access type for the VMs in this subnet.
	// +kubebuilder:validation:Optional
	PrivateIPv6GoogleAccess *string `json:"privateIpv6GoogleAccess,omitempty" tf:"private_ipv6_google_access,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of the resource. This field can be either PRIVATE
	// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
	// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
	// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
	// purpose defaults to PRIVATE.
	//
	// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set 'role'.
	// +kubebuilder:validation:Optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// The GCP region for this subnetwork.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The role of subnetwork. Currently, this field is only used when
	// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
	// or BACKUP. An ACTIVE subnetwork is one that is currently being used
	// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
	// is ready to be promoted to ACTIVE or is currently draining. Possible values: ["ACTIVE", "BACKUP"]
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges.
	//
	// **Note**: This field uses [attr-as-block mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html) to avoid
	// breaking users during the 0.12 upgrade. To explicitly send a list
	// of zero objects you must use the following syntax:
	// 'example=[]'
	// For more details about this behavior, see [this section](https://www.terraform.io/docs/configuration/attr-as-blocks.html#defining-a-fixed-object-collection-value).
	// +kubebuilder:validation:Optional
	SecondaryIPRange []SecondaryIPRangeParameters `json:"secondaryIpRange,omitempty" tf:"secondary_ip_range,omitempty"`

	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not.
	// If not specified IPV4_ONLY will be used. Possible values: ["IPV4_ONLY", "IPV4_IPV6"]
	// +kubebuilder:validation:Optional
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`
}

// SubnetworkSpec defines the desired state of Subnetwork
type SubnetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetworkParameters_2 `json:"forProvider"`
}

// SubnetworkStatus defines the observed state of Subnetwork.
type SubnetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetworkObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subnetwork is the Schema for the Subnetworks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Subnetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetworkSpec   `json:"spec"`
	Status            SubnetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetworkList contains a list of Subnetworks
type SubnetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnetwork `json:"items"`
}

// Repository type metadata.
var (
	Subnetwork_Kind             = "Subnetwork"
	Subnetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subnetwork_Kind}.String()
	Subnetwork_KindAPIVersion   = Subnetwork_Kind + "." + CRDGroupVersion.String()
	Subnetwork_GroupVersionKind = CRDGroupVersion.WithKind(Subnetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&Subnetwork{}, &SubnetworkList{})
}
