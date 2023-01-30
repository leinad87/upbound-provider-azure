/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OutputServiceBusTopicObservation struct {

	// The ID of the Stream Analytics Output ServiceBus Topic.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OutputServiceBusTopicParameters struct {

	// The authentication mode for the Stream Output. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	// +kubebuilder:validation:Optional
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The name of the Stream Output. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A list of property columns to add to the Service Bus Topic output.
	// +kubebuilder:validation:Optional
	PropertyColumns []*string `json:"propertyColumns,omitempty" tf:"property_columns,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A serialization block as defined below.
	// +kubebuilder:validation:Required
	Serialization []OutputServiceBusTopicSerializationParameters `json:"serialization" tf:"serialization,omitempty"`

	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta1.ServiceBusNamespace
	// +kubebuilder:validation:Optional
	ServiceBusNamespace *string `json:"servicebusNamespace,omitempty" tf:"servicebus_namespace,omitempty"`

	// Reference to a ServiceBusNamespace in servicebus to populate servicebusNamespace.
	// +kubebuilder:validation:Optional
	ServiceBusNamespaceRef *v1.Reference `json:"servicebusNamespaceRef,omitempty" tf:"-"`

	// Selector for a ServiceBusNamespace in servicebus to populate servicebusNamespace.
	// +kubebuilder:validation:Optional
	ServiceBusNamespaceSelector *v1.Selector `json:"servicebusNamespaceSelector,omitempty" tf:"-"`

	// The shared access policy key for the specified shared access policy. Required if authentication_mode is ConnectionString.
	// +kubebuilder:validation:Optional
	SharedAccessPolicyKeySecretRef *v1.SecretKeySelector `json:"sharedAccessPolicyKeySecretRef,omitempty" tf:"-"`

	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc. Required if authentication_mode is ConnectionString.
	// +kubebuilder:validation:Optional
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty" tf:"shared_access_policy_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Job
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Reference to a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// Selector for a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`

	// A key-value pair of system property columns that will be attached to the outgoing messages for the Service Bus Topic Output.
	// +kubebuilder:validation:Optional
	SystemPropertyColumns map[string]*string `json:"systemPropertyColumns,omitempty" tf:"system_property_columns,omitempty"`

	// The name of the Service Bus Topic.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta1.Topic
	// +kubebuilder:validation:Optional
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`

	// Reference to a Topic in servicebus to populate topicName.
	// +kubebuilder:validation:Optional
	TopicNameRef *v1.Reference `json:"topicNameRef,omitempty" tf:"-"`

	// Selector for a Topic in servicebus to populate topicName.
	// +kubebuilder:validation:Optional
	TopicNameSelector *v1.Selector `json:"topicNameSelector,omitempty" tf:"-"`
}

type OutputServiceBusTopicSerializationObservation struct {
}

type OutputServiceBusTopicSerializationParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// Specifies the format of the JSON the output will be written in. Possible values are Array and LineSeparated.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The serialization format used for outgoing data streams. Possible values are Avro, Csv, Json and Parquet.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// OutputServiceBusTopicSpec defines the desired state of OutputServiceBusTopic
type OutputServiceBusTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputServiceBusTopicParameters `json:"forProvider"`
}

// OutputServiceBusTopicStatus defines the observed state of OutputServiceBusTopic.
type OutputServiceBusTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputServiceBusTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutputServiceBusTopic is the Schema for the OutputServiceBusTopics API. Manages a Stream Analytics Output to a ServiceBus Topic.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OutputServiceBusTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OutputServiceBusTopicSpec   `json:"spec"`
	Status            OutputServiceBusTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputServiceBusTopicList contains a list of OutputServiceBusTopics
type OutputServiceBusTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputServiceBusTopic `json:"items"`
}

// Repository type metadata.
var (
	OutputServiceBusTopic_Kind             = "OutputServiceBusTopic"
	OutputServiceBusTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputServiceBusTopic_Kind}.String()
	OutputServiceBusTopic_KindAPIVersion   = OutputServiceBusTopic_Kind + "." + CRDGroupVersion.String()
	OutputServiceBusTopic_GroupVersionKind = CRDGroupVersion.WithKind(OutputServiceBusTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputServiceBusTopic{}, &OutputServiceBusTopicList{})
}
