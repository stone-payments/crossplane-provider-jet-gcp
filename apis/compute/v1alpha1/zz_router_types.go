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

type AdvertisedIPRangesObservation struct {
}

type AdvertisedIPRangesParameters struct {

	// User-specified description for the IP range.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP range to advertise. The value must be a
	// CIDR-formatted string.
	// +kubebuilder:validation:Required
	Range *string `json:"range" tf:"range,omitempty"`
}

type BgpObservation struct {
}

type BgpParameters struct {

	// User-specified flag to indicate which mode to use for advertisement. Default value: "DEFAULT" Possible values: ["DEFAULT", "CUSTOM"]
	// +kubebuilder:validation:Optional
	AdvertiseMode *string `json:"advertiseMode,omitempty" tf:"advertise_mode,omitempty"`

	// User-specified list of prefix groups to advertise in custom mode.
	// This field can only be populated if advertiseMode is CUSTOM and
	// is advertised to all peers of the router. These groups will be
	// advertised in addition to any specified prefixes. Leave this field
	// blank to advertise no custom groups.
	//
	// This enum field has the one valid value: ALL_SUBNETS
	// +kubebuilder:validation:Optional
	AdvertisedGroups []*string `json:"advertisedGroups,omitempty" tf:"advertised_groups,omitempty"`

	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is CUSTOM and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// +kubebuilder:validation:Optional
	AdvertisedIPRanges []AdvertisedIPRangesParameters `json:"advertisedIpRanges,omitempty" tf:"advertised_ip_ranges,omitempty"`

	// Local BGP Autonomous System Number (ASN). Must be an RFC6996
	// private ASN, either 16-bit or 32-bit. The value will be fixed for
	// this router resource. All VPN tunnels that link to this router
	// will have the same local ASN.
	// +kubebuilder:validation:Required
	Asn *int64 `json:"asn" tf:"asn,omitempty"`
}

type RouterObservation struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type RouterParameters struct {

	// BGP information specific to this router.
	// +kubebuilder:validation:Optional
	Bgp []BgpParameters `json:"bgp,omitempty" tf:"bgp,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Field to indicate if a router is dedicated to use with encrypted
	// Interconnect Attachment (IPsec-encrypted Cloud Interconnect feature).
	//
	// Not currently available publicly.
	// +kubebuilder:validation:Optional
	EncryptedInterconnectRouter *bool `json:"encryptedInterconnectRouter,omitempty" tf:"encrypted_interconnect_router,omitempty"`

	// A reference to the network to which this router belongs.
	// +crossplane:generate:reference:type=Network
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the router resides.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// RouterSpec defines the desired state of Router
type RouterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterParameters `json:"forProvider"`
}

// RouterStatus defines the observed state of Router.
type RouterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Router is the Schema for the Routers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Router struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouterSpec   `json:"spec"`
	Status            RouterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterList contains a list of Routers
type RouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Router `json:"items"`
}

// Repository type metadata.
var (
	Router_Kind             = "Router"
	Router_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Router_Kind}.String()
	Router_KindAPIVersion   = Router_Kind + "." + CRDGroupVersion.String()
	Router_GroupVersionKind = CRDGroupVersion.WithKind(Router_Kind)
)

func init() {
	SchemeBuilder.Register(&Router{}, &RouterList{})
}
