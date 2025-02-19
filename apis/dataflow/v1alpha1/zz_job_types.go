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

type JobObservation struct {
	JobID *string `json:"jobId,omitempty" tf:"job_id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type JobParameters struct {

	// List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].
	// +kubebuilder:validation:Optional
	AdditionalExperiments []*string `json:"additionalExperiments,omitempty" tf:"additional_experiments,omitempty"`

	// Indicates if the job should use the streaming engine feature.
	// +kubebuilder:validation:Optional
	EnableStreamingEngine *bool `json:"enableStreamingEngine,omitempty" tf:"enable_streaming_engine,omitempty"`

	// The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".
	// +kubebuilder:validation:Optional
	IPConfiguration *string `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// The name for the Cloud KMS key for the job. Key format is: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY
	// +kubebuilder:validation:Optional
	KmsKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// User labels to be specified for the job. Keys and values should follow the restrictions specified in the labeling restrictions page. NOTE: Google-provided Dataflow templates often provide default labels that begin with goog-dataflow-provided. Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The machine type to use for the job.
	// +kubebuilder:validation:Optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// The number of workers permitted to work on the job. More workers may improve processing speed at additional cost.
	// +kubebuilder:validation:Optional
	MaxWorkers *int64 `json:"maxWorkers,omitempty" tf:"max_workers,omitempty"`

	// A unique name for the resource, required by Dataflow.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// One of "drain" or "cancel". Specifies behavior of deletion during terraform destroy.
	// +kubebuilder:validation:Optional
	OnDelete *string `json:"onDelete,omitempty" tf:"on_delete,omitempty"`

	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	// +kubebuilder:validation:Optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The project in which the resource belongs.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region in which the created job should run.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The Service Account email used to create the job.
	// +kubebuilder:validation:Optional
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// A writeable location on Google Cloud Storage for the Dataflow job to dump its temporary data.
	// +kubebuilder:validation:Required
	TempGcsLocation *string `json:"tempGcsLocation" tf:"temp_gcs_location,omitempty"`

	// The Google Cloud Storage path to the Dataflow job template.
	// +kubebuilder:validation:Required
	TemplateGcsPath *string `json:"templateGcsPath" tf:"template_gcs_path,omitempty"`

	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job.
	// +kubebuilder:validation:Optional
	TransformNameMapping map[string]string `json:"transformNameMapping,omitempty" tf:"transform_name_mapping,omitempty"`

	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// JobSpec defines the desired state of Job
type JobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobParameters `json:"forProvider"`
}

// JobStatus defines the observed state of Job.
type JobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Job is the Schema for the Jobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Job struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobSpec   `json:"spec"`
	Status            JobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobList contains a list of Jobs
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Job `json:"items"`
}

// Repository type metadata.
var (
	Job_Kind             = "Job"
	Job_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Job_Kind}.String()
	Job_KindAPIVersion   = Job_Kind + "." + CRDGroupVersion.String()
	Job_GroupVersionKind = CRDGroupVersion.WithKind(Job_Kind)
)

func init() {
	SchemeBuilder.Register(&Job{}, &JobList{})
}
