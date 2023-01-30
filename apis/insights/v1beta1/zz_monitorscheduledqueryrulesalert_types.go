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

type MonitorScheduledQueryRulesAlertActionObservation struct {
}

type MonitorScheduledQueryRulesAlertActionParameters struct {

	// List of action group reference resource IDs.
	// +crossplane:generate:reference:type=MonitorActionGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ActionGroup []*string `json:"actionGroup,omitempty" tf:"action_group,omitempty"`

	// References to MonitorActionGroup to populate actionGroup.
	// +kubebuilder:validation:Optional
	ActionGroupRefs []v1.Reference `json:"actionGroupRefs,omitempty" tf:"-"`

	// Selector for a list of MonitorActionGroup to populate actionGroup.
	// +kubebuilder:validation:Optional
	ActionGroupSelector *v1.Selector `json:"actionGroupSelector,omitempty" tf:"-"`

	// Custom payload to be sent for all webhook payloads in alerting action.
	// +kubebuilder:validation:Optional
	CustomWebhookPayload *string `json:"customWebhookPayload,omitempty" tf:"custom_webhook_payload,omitempty"`

	// Custom subject override for all email ids in Azure action group.
	// +kubebuilder:validation:Optional
	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`
}

type MonitorScheduledQueryRulesAlertObservation struct {

	// The ID of the scheduled query rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MonitorScheduledQueryRulesAlertParameters struct {

	// An action block as defined below.
	// +kubebuilder:validation:Required
	Action []MonitorScheduledQueryRulesAlertActionParameters `json:"action" tf:"action,omitempty"`

	// List of Resource IDs referred into query.
	// +kubebuilder:validation:Optional
	AuthorizedResourceIds []*string `json:"authorizedResourceIds,omitempty" tf:"authorized_resource_ids,omitempty"`

	// Should the alerts in this Metric Alert be auto resolved? Defaults to false.
	// -> NOTE auto_mitigation_enabled and throttling are mutually exclusive and cannot both be set.
	// +kubebuilder:validation:Optional
	AutoMitigationEnabled *bool `json:"autoMitigationEnabled,omitempty" tf:"auto_mitigation_enabled,omitempty"`

	// The resource URI over which log search query is to be run.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// Reference to a ApplicationInsights in insights to populate dataSourceId.
	// +kubebuilder:validation:Optional
	DataSourceIDRef *v1.Reference `json:"dataSourceIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights in insights to populate dataSourceId.
	// +kubebuilder:validation:Optional
	DataSourceIDSelector *v1.Selector `json:"dataSourceIdSelector,omitempty" tf:"-"`

	// The description of the scheduled query rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this scheduled query rule is enabled. Default is true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Frequency (in minutes) at which rule condition should be evaluated. Values must be between 5 and 1440 (inclusive).
	// +kubebuilder:validation:Required
	Frequency *float64 `json:"frequency" tf:"frequency,omitempty"`

	// Specifies the Azure Region where the resource should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the scheduled query rule. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Log search query.
	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`

	// The type of query results. Possible values are ResultCount and Number. Default is ResultCount.
	// +kubebuilder:validation:Optional
	QueryType *string `json:"queryType,omitempty" tf:"query_type,omitempty"`

	// The name of the resource group in which to create the scheduled query rule instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Severity of the alert. Possible values include: 0, 1, 2, 3, or 4.
	// +kubebuilder:validation:Optional
	Severity *float64 `json:"severity,omitempty" tf:"severity,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Time (in minutes) for which Alerts should be throttled or suppressed. Values must be between 0 and 10000 (inclusive).
	// +kubebuilder:validation:Optional
	Throttling *float64 `json:"throttling,omitempty" tf:"throttling,omitempty"`

	// Time window for which data needs to be fetched for query (must be greater than or equal to frequency). Values must be between 5 and 2880 (inclusive).
	// +kubebuilder:validation:Required
	TimeWindow *float64 `json:"timeWindow" tf:"time_window,omitempty"`

	// A trigger block as defined below.
	// +kubebuilder:validation:Required
	Trigger []TriggerParameters `json:"trigger" tf:"trigger,omitempty"`
}

type TriggerMetricTriggerObservation struct {
}

type TriggerMetricTriggerParameters struct {

	// Evaluation of metric on a particular column.
	// +kubebuilder:validation:Required
	MetricColumn *string `json:"metricColumn" tf:"metric_column,omitempty"`

	// Metric Trigger Type - 'Consecutive' or 'Total'.
	// +kubebuilder:validation:Required
	MetricTriggerType *string `json:"metricTriggerType" tf:"metric_trigger_type,omitempty"`

	// Evaluation operation for rule - 'GreaterThan', GreaterThanOrEqual', 'LessThan', or 'LessThanOrEqual'.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Result or count threshold based on which rule should be triggered. Values must be between 0 and 10000 inclusive.
	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

type TriggerObservation struct {
}

type TriggerParameters struct {

	// A metric_trigger block as defined above. Trigger condition for metric query rule.
	// +kubebuilder:validation:Optional
	MetricTrigger []TriggerMetricTriggerParameters `json:"metricTrigger,omitempty" tf:"metric_trigger,omitempty"`

	// Evaluation operation for rule - 'GreaterThan', GreaterThanOrEqual', 'LessThan', or 'LessThanOrEqual'.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Result or count threshold based on which rule should be triggered. Values must be between 0 and 10000 inclusive.
	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

// MonitorScheduledQueryRulesAlertSpec defines the desired state of MonitorScheduledQueryRulesAlert
type MonitorScheduledQueryRulesAlertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorScheduledQueryRulesAlertParameters `json:"forProvider"`
}

// MonitorScheduledQueryRulesAlertStatus defines the observed state of MonitorScheduledQueryRulesAlert.
type MonitorScheduledQueryRulesAlertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorScheduledQueryRulesAlertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorScheduledQueryRulesAlert is the Schema for the MonitorScheduledQueryRulesAlerts API. Manages an AlertingAction Scheduled Query Rules resource within Azure Monitor
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorScheduledQueryRulesAlert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorScheduledQueryRulesAlertSpec   `json:"spec"`
	Status            MonitorScheduledQueryRulesAlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorScheduledQueryRulesAlertList contains a list of MonitorScheduledQueryRulesAlerts
type MonitorScheduledQueryRulesAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorScheduledQueryRulesAlert `json:"items"`
}

// Repository type metadata.
var (
	MonitorScheduledQueryRulesAlert_Kind             = "MonitorScheduledQueryRulesAlert"
	MonitorScheduledQueryRulesAlert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorScheduledQueryRulesAlert_Kind}.String()
	MonitorScheduledQueryRulesAlert_KindAPIVersion   = MonitorScheduledQueryRulesAlert_Kind + "." + CRDGroupVersion.String()
	MonitorScheduledQueryRulesAlert_GroupVersionKind = CRDGroupVersion.WithKind(MonitorScheduledQueryRulesAlert_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorScheduledQueryRulesAlert{}, &MonitorScheduledQueryRulesAlertList{})
}
