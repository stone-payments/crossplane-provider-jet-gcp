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

type ApisObservation struct {
	Methods []MethodsObservation `json:"methods,omitempty" tf:"methods,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Syntax *string `json:"syntax,omitempty" tf:"syntax,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ApisParameters struct {
}

type EndpointsObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type EndpointsParameters struct {
}

type MethodsObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	RequestType *string `json:"requestType,omitempty" tf:"request_type,omitempty"`

	ResponseType *string `json:"responseType,omitempty" tf:"response_type,omitempty"`

	Syntax *string `json:"syntax,omitempty" tf:"syntax,omitempty"`
}

type MethodsParameters struct {
}

type ServiceObservation struct {
	Apis []ApisObservation `json:"apis,omitempty" tf:"apis,omitempty"`

	ConfigID *string `json:"configId,omitempty" tf:"config_id,omitempty"`

	DNSAddress *string `json:"dnsAddress,omitempty" tf:"dns_address,omitempty"`

	Endpoints []EndpointsObservation `json:"endpoints,omitempty" tf:"endpoints,omitempty"`
}

type ServiceParameters struct {

	// The full text of the Service Config YAML file (Example located here). If provided, must also provide protoc_output_base64. open_api config must not be provided.
	// +kubebuilder:validation:Optional
	GrpcConfig *string `json:"grpcConfig,omitempty" tf:"grpc_config,omitempty"`

	// The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and protoc_output_base64 must be specified.
	// +kubebuilder:validation:Optional
	OpenapiConfig *string `json:"openapiConfig,omitempty" tf:"openapi_config,omitempty"`

	// The project ID that the service belongs to. If not provided, provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file, base64-encoded.
	// +kubebuilder:validation:Optional
	ProtocOutputBase64 *string `json:"protocOutputBase64,omitempty" tf:"protoc_output_base64,omitempty"`

	// The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Service is the Schema for the Services API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec"`
	Status            ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	Service_Kind             = "Service"
	Service_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Service_Kind}.String()
	Service_KindAPIVersion   = Service_Kind + "." + CRDGroupVersion.String()
	Service_GroupVersionKind = CRDGroupVersion.WithKind(Service_Kind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
