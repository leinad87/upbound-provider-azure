// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type IdentityProviderGoogleInitParameters struct {

	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// Client Id for Google Sign-in.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

type IdentityProviderGoogleObservation struct {

	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Client Id for Google Sign-in.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The ID of the API Management Google Identity Provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type IdentityProviderGoogleParameters struct {

	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// Client Id for Google Sign-in.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Client secret for Google Sign-in.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// IdentityProviderGoogleSpec defines the desired state of IdentityProviderGoogle
type IdentityProviderGoogleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityProviderGoogleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider IdentityProviderGoogleInitParameters `json:"initProvider,omitempty"`
}

// IdentityProviderGoogleStatus defines the observed state of IdentityProviderGoogle.
type IdentityProviderGoogleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityProviderGoogleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderGoogle is the Schema for the IdentityProviderGoogles API. Manages an API Management Google Identity Provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IdentityProviderGoogle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId) || (has(self.initProvider) && has(self.initProvider.clientId))",message="spec.forProvider.clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientSecretSecretRef)",message="spec.forProvider.clientSecretSecretRef is a required parameter"
	Spec   IdentityProviderGoogleSpec   `json:"spec"`
	Status IdentityProviderGoogleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderGoogleList contains a list of IdentityProviderGoogles
type IdentityProviderGoogleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderGoogle `json:"items"`
}

// Repository type metadata.
var (
	IdentityProviderGoogle_Kind             = "IdentityProviderGoogle"
	IdentityProviderGoogle_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProviderGoogle_Kind}.String()
	IdentityProviderGoogle_KindAPIVersion   = IdentityProviderGoogle_Kind + "." + CRDGroupVersion.String()
	IdentityProviderGoogle_GroupVersionKind = CRDGroupVersion.WithKind(IdentityProviderGoogle_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityProviderGoogle{}, &IdentityProviderGoogleList{})
}
