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

type BGPInitParameters struct {

	// The BGP speaker's ASN.
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// The BGP peering IP address.
	PeeringAddress *string `json:"peeringAddress,omitempty" tf:"peering_address,omitempty"`
}

type BGPObservation struct {

	// The BGP speaker's ASN.
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// The BGP peering IP address.
	PeeringAddress *string `json:"peeringAddress,omitempty" tf:"peering_address,omitempty"`
}

type BGPParameters struct {

	// The BGP speaker's ASN.
	// +kubebuilder:validation:Optional
	Asn *float64 `json:"asn" tf:"asn,omitempty"`

	// The BGP peering IP address.
	// +kubebuilder:validation:Optional
	PeeringAddress *string `json:"peeringAddress" tf:"peering_address,omitempty"`
}

type LinkInitParameters struct {

	// A bgp block as defined above.
	BGP []BGPInitParameters `json:"bgp,omitempty" tf:"bgp,omitempty"`

	// The FQDN of this VPN Site Link.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The IP address of this VPN Site Link.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The name which should be used for this VPN Site Link.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the physical link at the VPN Site. Example: ATT, Verizon.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// The speed of the VPN device at the branch location in unit of mbps. Defaults to 0.
	SpeedInMbps *float64 `json:"speedInMbps,omitempty" tf:"speed_in_mbps,omitempty"`
}

type LinkObservation struct {

	// A bgp block as defined above.
	BGP []BGPObservation `json:"bgp,omitempty" tf:"bgp,omitempty"`

	// The FQDN of this VPN Site Link.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The ID of the VPN Site Link.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address of this VPN Site Link.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The name which should be used for this VPN Site Link.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the physical link at the VPN Site. Example: ATT, Verizon.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// The speed of the VPN device at the branch location in unit of mbps. Defaults to 0.
	SpeedInMbps *float64 `json:"speedInMbps,omitempty" tf:"speed_in_mbps,omitempty"`
}

type LinkParameters struct {

	// A bgp block as defined above.
	// +kubebuilder:validation:Optional
	BGP []BGPParameters `json:"bgp,omitempty" tf:"bgp,omitempty"`

	// The FQDN of this VPN Site Link.
	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The IP address of this VPN Site Link.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The name which should be used for this VPN Site Link.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The name of the physical link at the VPN Site. Example: ATT, Verizon.
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// The speed of the VPN device at the branch location in unit of mbps. Defaults to 0.
	// +kubebuilder:validation:Optional
	SpeedInMbps *float64 `json:"speedInMbps,omitempty" tf:"speed_in_mbps,omitempty"`
}

type O365PolicyInitParameters struct {

	// A traffic_category block as defined above.
	TrafficCategory []TrafficCategoryInitParameters `json:"trafficCategory,omitempty" tf:"traffic_category,omitempty"`
}

type O365PolicyObservation struct {

	// A traffic_category block as defined above.
	TrafficCategory []TrafficCategoryObservation `json:"trafficCategory,omitempty" tf:"traffic_category,omitempty"`
}

type O365PolicyParameters struct {

	// A traffic_category block as defined above.
	// +kubebuilder:validation:Optional
	TrafficCategory []TrafficCategoryParameters `json:"trafficCategory,omitempty" tf:"traffic_category,omitempty"`
}

type TrafficCategoryInitParameters struct {

	// Is allow endpoint enabled? The Allow endpoint is required for connectivity to specific O365 services and features, but are not as sensitive to network performance and latency as other endpoint types. Defaults to false.
	AllowEndpointEnabled *bool `json:"allowEndpointEnabled,omitempty" tf:"allow_endpoint_enabled,omitempty"`

	// Is default endpoint enabled? The Default endpoint represents O365 services and dependencies that do not require any optimization, and can be treated by customer networks as normal Internet bound traffic. Defaults to false.
	DefaultEndpointEnabled *bool `json:"defaultEndpointEnabled,omitempty" tf:"default_endpoint_enabled,omitempty"`

	// Is optimize endpoint enabled? The Optimize endpoint is required for connectivity to every O365 service and represents the O365 scenario that is the most sensitive to network performance, latency, and availability. Defaults to false.
	OptimizeEndpointEnabled *bool `json:"optimizeEndpointEnabled,omitempty" tf:"optimize_endpoint_enabled,omitempty"`
}

type TrafficCategoryObservation struct {

	// Is allow endpoint enabled? The Allow endpoint is required for connectivity to specific O365 services and features, but are not as sensitive to network performance and latency as other endpoint types. Defaults to false.
	AllowEndpointEnabled *bool `json:"allowEndpointEnabled,omitempty" tf:"allow_endpoint_enabled,omitempty"`

	// Is default endpoint enabled? The Default endpoint represents O365 services and dependencies that do not require any optimization, and can be treated by customer networks as normal Internet bound traffic. Defaults to false.
	DefaultEndpointEnabled *bool `json:"defaultEndpointEnabled,omitempty" tf:"default_endpoint_enabled,omitempty"`

	// Is optimize endpoint enabled? The Optimize endpoint is required for connectivity to every O365 service and represents the O365 scenario that is the most sensitive to network performance, latency, and availability. Defaults to false.
	OptimizeEndpointEnabled *bool `json:"optimizeEndpointEnabled,omitempty" tf:"optimize_endpoint_enabled,omitempty"`
}

type TrafficCategoryParameters struct {

	// Is allow endpoint enabled? The Allow endpoint is required for connectivity to specific O365 services and features, but are not as sensitive to network performance and latency as other endpoint types. Defaults to false.
	// +kubebuilder:validation:Optional
	AllowEndpointEnabled *bool `json:"allowEndpointEnabled,omitempty" tf:"allow_endpoint_enabled,omitempty"`

	// Is default endpoint enabled? The Default endpoint represents O365 services and dependencies that do not require any optimization, and can be treated by customer networks as normal Internet bound traffic. Defaults to false.
	// +kubebuilder:validation:Optional
	DefaultEndpointEnabled *bool `json:"defaultEndpointEnabled,omitempty" tf:"default_endpoint_enabled,omitempty"`

	// Is optimize endpoint enabled? The Optimize endpoint is required for connectivity to every O365 service and represents the O365 scenario that is the most sensitive to network performance, latency, and availability. Defaults to false.
	// +kubebuilder:validation:Optional
	OptimizeEndpointEnabled *bool `json:"optimizeEndpointEnabled,omitempty" tf:"optimize_endpoint_enabled,omitempty"`
}

type VPNSiteInitParameters struct {

	// Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
	// +listType=set
	AddressCidrs []*string `json:"addressCidrs,omitempty" tf:"address_cidrs,omitempty"`

	// The model of the VPN device.
	DeviceModel *string `json:"deviceModel,omitempty" tf:"device_model,omitempty"`

	// The name of the VPN device vendor.
	DeviceVendor *string `json:"deviceVendor,omitempty" tf:"device_vendor,omitempty"`

	// One or more link blocks as defined below.
	Link []LinkInitParameters `json:"link,omitempty" tf:"link,omitempty"`

	// The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// An o365_policy block as defined below.
	O365Policy []O365PolicyInitParameters `json:"o365Policy,omitempty" tf:"o365_policy,omitempty"`

	// A mapping of tags which should be assigned to the VPN Site.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualWAN
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	VirtualWanID *string `json:"virtualWanId,omitempty" tf:"virtual_wan_id,omitempty"`

	// Reference to a VirtualWAN in network to populate virtualWanId.
	// +kubebuilder:validation:Optional
	VirtualWanIDRef *v1.Reference `json:"virtualWanIdRef,omitempty" tf:"-"`

	// Selector for a VirtualWAN in network to populate virtualWanId.
	// +kubebuilder:validation:Optional
	VirtualWanIDSelector *v1.Selector `json:"virtualWanIdSelector,omitempty" tf:"-"`
}

type VPNSiteObservation struct {

	// Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
	// +listType=set
	AddressCidrs []*string `json:"addressCidrs,omitempty" tf:"address_cidrs,omitempty"`

	// The model of the VPN device.
	DeviceModel *string `json:"deviceModel,omitempty" tf:"device_model,omitempty"`

	// The name of the VPN device vendor.
	DeviceVendor *string `json:"deviceVendor,omitempty" tf:"device_vendor,omitempty"`

	// The ID of the VPN Site.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more link blocks as defined below.
	Link []LinkObservation `json:"link,omitempty" tf:"link,omitempty"`

	// The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// An o365_policy block as defined below.
	O365Policy []O365PolicyObservation `json:"o365Policy,omitempty" tf:"o365_policy,omitempty"`

	// The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags which should be assigned to the VPN Site.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
	VirtualWanID *string `json:"virtualWanId,omitempty" tf:"virtual_wan_id,omitempty"`
}

type VPNSiteParameters struct {

	// Specifies a list of IP address CIDRs that are located on your on-premises site. Traffic destined for these address spaces is routed to your local site.
	// +kubebuilder:validation:Optional
	// +listType=set
	AddressCidrs []*string `json:"addressCidrs,omitempty" tf:"address_cidrs,omitempty"`

	// The model of the VPN device.
	// +kubebuilder:validation:Optional
	DeviceModel *string `json:"deviceModel,omitempty" tf:"device_model,omitempty"`

	// The name of the VPN device vendor.
	// +kubebuilder:validation:Optional
	DeviceVendor *string `json:"deviceVendor,omitempty" tf:"device_vendor,omitempty"`

	// One or more link blocks as defined below.
	// +kubebuilder:validation:Optional
	Link []LinkParameters `json:"link,omitempty" tf:"link,omitempty"`

	// The Azure Region where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// An o365_policy block as defined below.
	// +kubebuilder:validation:Optional
	O365Policy []O365PolicyParameters `json:"o365Policy,omitempty" tf:"o365_policy,omitempty"`

	// The name of the Resource Group where the VPN Site should exist. Changing this forces a new VPN Site to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the VPN Site.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the Virtual Wan where this VPN site resides in. Changing this forces a new VPN Site to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualWAN
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualWanID *string `json:"virtualWanId,omitempty" tf:"virtual_wan_id,omitempty"`

	// Reference to a VirtualWAN in network to populate virtualWanId.
	// +kubebuilder:validation:Optional
	VirtualWanIDRef *v1.Reference `json:"virtualWanIdRef,omitempty" tf:"-"`

	// Selector for a VirtualWAN in network to populate virtualWanId.
	// +kubebuilder:validation:Optional
	VirtualWanIDSelector *v1.Selector `json:"virtualWanIdSelector,omitempty" tf:"-"`
}

// VPNSiteSpec defines the desired state of VPNSite
type VPNSiteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNSiteParameters `json:"forProvider"`
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
	InitProvider VPNSiteInitParameters `json:"initProvider,omitempty"`
}

// VPNSiteStatus defines the observed state of VPNSite.
type VPNSiteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNSiteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPNSite is the Schema for the VPNSites API. Manages a VPN Site.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VPNSite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   VPNSiteSpec   `json:"spec"`
	Status VPNSiteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNSiteList contains a list of VPNSites
type VPNSiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNSite `json:"items"`
}

// Repository type metadata.
var (
	VPNSite_Kind             = "VPNSite"
	VPNSite_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNSite_Kind}.String()
	VPNSite_KindAPIVersion   = VPNSite_Kind + "." + CRDGroupVersion.String()
	VPNSite_GroupVersionKind = CRDGroupVersion.WithKind(VPNSite_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNSite{}, &VPNSiteList{})
}
