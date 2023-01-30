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

type PolicyDefinitionObservation struct {

	// The ID of the Policy Definition.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of role definition id extracted from policy_rule required for remediation.
	RoleDefinitionIds []*string `json:"roleDefinitionIds,omitempty" tf:"role_definition_ids,omitempty"`
}

type PolicyDefinitionParameters struct {

	// The description of the policy definition.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the policy definition.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// The id of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ManagementGroupID *string `json:"managementGroupId,omitempty" tf:"management_group_id,omitempty"`

	// The metadata for the policy definition. This is a JSON string representing additional metadata that should be stored with the policy definition.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The policy resource manager mode that allows you to specify which resource types will be evaluated. Possible values are All, Indexed, Microsoft.ContainerService.Data, Microsoft.CustomerLockbox.Data, Microsoft.DataCatalog.Data, Microsoft.KeyVault.Data, Microsoft.Kubernetes.Data, Microsoft.MachineLearningServices.Data, Microsoft.Network.Data and Microsoft.Synapse.Data.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// Parameters for the policy definition. This field is a JSON string that allows you to parameterize your policy definition.
	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The policy rule for the policy definition. This is a JSON string representing the rule that contains an if and a then block.
	// +kubebuilder:validation:Optional
	PolicyRule *string `json:"policyRule,omitempty" tf:"policy_rule,omitempty"`

	// The policy type. Possible values are BuiltIn, Custom, NotSpecified and Static. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	PolicyType *string `json:"policyType" tf:"policy_type,omitempty"`
}

// PolicyDefinitionSpec defines the desired state of PolicyDefinition
type PolicyDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyDefinitionParameters `json:"forProvider"`
}

// PolicyDefinitionStatus defines the observed state of PolicyDefinition.
type PolicyDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyDefinition is the Schema for the PolicyDefinitions API. Manages a policy rule definition. Policy definitions do not take effect until they are assigned to a scope using a Policy Assignment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PolicyDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyDefinitionSpec   `json:"spec"`
	Status            PolicyDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyDefinitionList contains a list of PolicyDefinitions
type PolicyDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyDefinition `json:"items"`
}

// Repository type metadata.
var (
	PolicyDefinition_Kind             = "PolicyDefinition"
	PolicyDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyDefinition_Kind}.String()
	PolicyDefinition_KindAPIVersion   = PolicyDefinition_Kind + "." + CRDGroupVersion.String()
	PolicyDefinition_GroupVersionKind = CRDGroupVersion.WithKind(PolicyDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyDefinition{}, &PolicyDefinitionList{})
}
