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

type DNSAAAARecordObservation struct {

	// The FQDN of the DNS AAAA Record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The DNS AAAA Record ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DNSAAAARecordParameters struct {

	// List of IPv6 Addresses. Conflicts with target_resource_id.
	// +kubebuilder:validation:Optional
	Records []*string `json:"records,omitempty" tf:"records,omitempty"`

	// Specifies the resource group where the DNS Zone  exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	TTL *float64 `json:"ttl" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Azure resource id of the target object. Conflicts with records
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`

	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=DNSZone
	// +kubebuilder:validation:Optional
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneNameRef *v1.Reference `json:"zoneNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ZoneNameSelector *v1.Selector `json:"zoneNameSelector,omitempty" tf:"-"`
}

// DNSAAAARecordSpec defines the desired state of DNSAAAARecord
type DNSAAAARecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DNSAAAARecordParameters `json:"forProvider"`
}

// DNSAAAARecordStatus defines the observed state of DNSAAAARecord.
type DNSAAAARecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DNSAAAARecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DNSAAAARecord is the Schema for the DNSAAAARecords API. Manages a DNS AAAA Record.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DNSAAAARecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DNSAAAARecordSpec   `json:"spec"`
	Status            DNSAAAARecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSAAAARecordList contains a list of DNSAAAARecords
type DNSAAAARecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSAAAARecord `json:"items"`
}

// Repository type metadata.
var (
	DNSAAAARecord_Kind             = "DNSAAAARecord"
	DNSAAAARecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DNSAAAARecord_Kind}.String()
	DNSAAAARecord_KindAPIVersion   = DNSAAAARecord_Kind + "." + CRDGroupVersion.String()
	DNSAAAARecord_GroupVersionKind = CRDGroupVersion.WithKind(DNSAAAARecord_Kind)
)

func init() {
	SchemeBuilder.Register(&DNSAAAARecord{}, &DNSAAAARecordList{})
}
