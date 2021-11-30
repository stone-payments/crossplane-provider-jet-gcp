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

type SslCertObservation struct {
	Cert *string `json:"cert,omitempty" tf:"cert,omitempty"`

	CertSerialNumber *string `json:"certSerialNumber,omitempty" tf:"cert_serial_number,omitempty"`

	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	ServerCaCert *string `json:"serverCaCert,omitempty" tf:"server_ca_cert,omitempty"`

	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty" tf:"sha1_fingerprint,omitempty"`
}

type SslCertParameters struct {

	// The common name to be used in the certificate to identify the client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	CommonName *string `json:"commonName" tf:"common_name,omitempty"`

	// The name of the Cloud SQL instance. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Instance *string `json:"instance" tf:"instance,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// SslCertSpec defines the desired state of SslCert
type SslCertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SslCertParameters `json:"forProvider"`
}

// SslCertStatus defines the observed state of SslCert.
type SslCertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SslCertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SslCert is the Schema for the SslCerts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type SslCert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SslCertSpec   `json:"spec"`
	Status            SslCertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SslCertList contains a list of SslCerts
type SslCertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SslCert `json:"items"`
}

// Repository type metadata.
var (
	SslCert_Kind             = "SslCert"
	SslCert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SslCert_Kind}.String()
	SslCert_KindAPIVersion   = SslCert_Kind + "." + CRDGroupVersion.String()
	SslCert_GroupVersionKind = CRDGroupVersion.WithKind(SslCert_Kind)
)

func init() {
	SchemeBuilder.Register(&SslCert{}, &SslCertList{})
}