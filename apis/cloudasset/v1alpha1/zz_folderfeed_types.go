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

type ConditionObservation struct {
}

type ConditionParameters struct {

	// Description of the expression. This is a longer text which describes the expression,
	// e.g. when hovered over it in a UI.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// String indicating the location of the expression for error reporting, e.g. a file
	// name and a position in the file.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Title for the expression, i.e. a short string describing its purpose.
	// This can be used e.g. in UIs which allow to enter the expression.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type FeedOutputConfigObservation struct {
}

type FeedOutputConfigParameters struct {

	// Destination on Cloud Pubsub.
	// +kubebuilder:validation:Required
	PubsubDestination []PubsubDestinationParameters `json:"pubsubDestination" tf:"pubsub_destination,omitempty"`
}

type FolderFeedObservation struct {
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FolderFeedParameters struct {

	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	// +kubebuilder:validation:Optional
	AssetNames []*string `json:"assetNames,omitempty" tf:"asset_names,omitempty"`

	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	// +kubebuilder:validation:Optional
	AssetTypes []*string `json:"assetTypes,omitempty" tf:"asset_types,omitempty"`

	// The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing.
	// +kubebuilder:validation:Required
	BillingProject *string `json:"billingProject" tf:"billing_project,omitempty"`

	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Asset content type. If not specified, no content but the asset name and type will be returned. Possible values: ["CONTENT_TYPE_UNSPECIFIED", "RESOURCE", "IAM_POLICY", "ORG_POLICY", "ACCESS_POLICY"]
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	// +kubebuilder:validation:Required
	FeedID *string `json:"feedId" tf:"feed_id,omitempty"`

	// Output configuration for asset feed destination.
	// +kubebuilder:validation:Required
	FeedOutputConfig []FeedOutputConfigParameters `json:"feedOutputConfig" tf:"feed_output_config,omitempty"`

	// The folder this feed should be created in.
	// +kubebuilder:validation:Required
	Folder *string `json:"folder" tf:"folder,omitempty"`
}

type PubsubDestinationObservation struct {
}

type PubsubDestinationParameters struct {

	// Destination on Cloud Pubsub topic.
	// +kubebuilder:validation:Required
	Topic *string `json:"topic" tf:"topic,omitempty"`
}

// FolderFeedSpec defines the desired state of FolderFeed
type FolderFeedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderFeedParameters `json:"forProvider"`
}

// FolderFeedStatus defines the observed state of FolderFeed.
type FolderFeedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderFeedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FolderFeed is the Schema for the FolderFeeds API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type FolderFeed struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderFeedSpec   `json:"spec"`
	Status            FolderFeedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderFeedList contains a list of FolderFeeds
type FolderFeedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderFeed `json:"items"`
}

// Repository type metadata.
var (
	FolderFeed_Kind             = "FolderFeed"
	FolderFeed_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderFeed_Kind}.String()
	FolderFeed_KindAPIVersion   = FolderFeed_Kind + "." + CRDGroupVersion.String()
	FolderFeed_GroupVersionKind = CRDGroupVersion.WithKind(FolderFeed_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderFeed{}, &FolderFeedList{})
}
