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

type CustomBGPAddressesObservation struct {
}

type CustomBGPAddressesParameters struct {

	// single IP address that is part of the azurerm_virtual_network_gateway ip_configuration
	// +kubebuilder:validation:Required
	Primary *string `json:"primary" tf:"primary,omitempty"`

	// single IP address that is part of the azurerm_virtual_network_gateway ip_configuration
	// +kubebuilder:validation:Required
	Secondary *string `json:"secondary" tf:"secondary,omitempty"`
}

type IpsecPolicyObservation struct {
}

type IpsecPolicyParameters struct {

	// The DH group used in IKE phase 1 for initial SA. Valid
	// options are DHGroup1, DHGroup14, DHGroup2, DHGroup2048, DHGroup24,
	// ECP256, ECP384, or None.
	// +kubebuilder:validation:Required
	DhGroup *string `json:"dhGroup" tf:"dh_group,omitempty"`

	// The IKE encryption algorithm. Valid
	// options are AES128, AES192, AES256, DES, DES3, GCMAES128, or GCMAES256.
	// +kubebuilder:validation:Required
	IkeEncryption *string `json:"ikeEncryption" tf:"ike_encryption,omitempty"`

	// The IKE integrity algorithm. Valid
	// options are GCMAES128, GCMAES256, MD5, SHA1, SHA256, or SHA384.
	// +kubebuilder:validation:Required
	IkeIntegrity *string `json:"ikeIntegrity" tf:"ike_integrity,omitempty"`

	// The IPSec encryption algorithm. Valid
	// options are AES128, AES192, AES256, DES, DES3, GCMAES128, GCMAES192, GCMAES256, or None.
	// +kubebuilder:validation:Required
	IpsecEncryption *string `json:"ipsecEncryption" tf:"ipsec_encryption,omitempty"`

	// The IPSec integrity algorithm. Valid
	// options are GCMAES128, GCMAES192, GCMAES256, MD5, SHA1, or SHA256.
	// +kubebuilder:validation:Required
	IpsecIntegrity *string `json:"ipsecIntegrity" tf:"ipsec_integrity,omitempty"`

	// The DH group used in IKE phase 2 for new child SA.
	// Valid options are ECP256, ECP384, PFS1, PFS14, PFS2, PFS2048, PFS24, PFSMM,
	// or None.
	// +kubebuilder:validation:Required
	PfsGroup *string `json:"pfsGroup" tf:"pfs_group,omitempty"`

	// The IPSec SA payload size in KB. Must be at least
	// 1024 KB. Defaults to 102400000 KB.
	// +kubebuilder:validation:Optional
	SaDatasize *float64 `json:"saDatasize,omitempty" tf:"sa_datasize,omitempty"`

	// The IPSec SA lifetime in seconds. Must be at least
	// 300 seconds. Defaults to 27000 seconds.
	// +kubebuilder:validation:Optional
	SaLifetime *float64 `json:"saLifetime,omitempty" tf:"sa_lifetime,omitempty"`
}

type TrafficSelectorPolicyObservation struct {
}

type TrafficSelectorPolicyParameters struct {

	// List of local CIDRs.
	// +kubebuilder:validation:Required
	LocalAddressCidrs []*string `json:"localAddressCidrs" tf:"local_address_cidrs,omitempty"`

	// List of remote CIDRs.
	// +kubebuilder:validation:Required
	RemoteAddressCidrs []*string `json:"remoteAddressCidrs" tf:"remote_address_cidrs,omitempty"`
}

type VirtualNetworkGatewayConnectionObservation struct {

	// The ID of the Virtual Network Gateway Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VirtualNetworkGatewayConnectionParameters struct {

	// The authorization key associated with the
	// Express Route Circuit. This field is required only if the type is an
	// ExpressRoute connection.
	// +kubebuilder:validation:Optional
	AuthorizationKeySecretRef *v1.SecretKeySelector `json:"authorizationKeySecretRef,omitempty" tf:"-"`

	// Connection mode to use. Possible
	// values are Default, InitiatorOnly and ResponderOnly. Defaults to Default.
	// Changing this value will force a resource to be created.
	// +kubebuilder:validation:Optional
	ConnectionMode *string `json:"connectionMode,omitempty" tf:"connection_mode,omitempty"`

	// The IKE protocol version to use. Possible
	// values are IKEv1 and IKEv2. Defaults to IKEv2.
	// Changing this value will force a resource to be created.
	// -> Note: Only valid for IPSec connections on virtual network gateways with SKU VpnGw1, VpnGw2, VpnGw3, VpnGw1AZ, VpnGw2AZ or VpnGw3AZ.
	// +kubebuilder:validation:Optional
	ConnectionProtocol *string `json:"connectionProtocol,omitempty" tf:"connection_protocol,omitempty"`

	// A custom_bgp_addresses  block which is documented below.
	// The block can only be used on IPSec / activeactive connections,
	// For details about see the relevant section in the Azure documentation.
	// +kubebuilder:validation:Optional
	CustomBGPAddresses []CustomBGPAddressesParameters `json:"customBgpAddresses,omitempty" tf:"custom_bgp_addresses,omitempty"`

	// The dead peer detection timeout of this connection in seconds. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DpdTimeoutSeconds *float64 `json:"dpdTimeoutSeconds,omitempty" tf:"dpd_timeout_seconds,omitempty"`

	// A list of the egress NAT Rule Ids.
	// +kubebuilder:validation:Optional
	EgressNATRuleIds []*string `json:"egressNatRuleIds,omitempty" tf:"egress_nat_rule_ids,omitempty"`

	// If true, BGP  is enabled
	// for this connection. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableBGP *bool `json:"enableBgp,omitempty" tf:"enable_bgp,omitempty"`

	// The ID of the Express Route Circuit
	// when creating an ExpressRoute connection .
	// The Express Route Circuit can be in the same or in a different subscription.
	// +kubebuilder:validation:Optional
	ExpressRouteCircuitID *string `json:"expressRouteCircuitId,omitempty" tf:"express_route_circuit_id,omitempty"`

	// If true, data packets will bypass ExpressRoute Gateway for data forwarding This is only valid for ExpressRoute connections.
	// +kubebuilder:validation:Optional
	ExpressRouteGatewayBypass *bool `json:"expressRouteGatewayBypass,omitempty" tf:"express_route_gateway_bypass,omitempty"`

	// A list of the ingress NAT Rule Ids.
	// +kubebuilder:validation:Optional
	IngressNATRuleIds []*string `json:"ingressNatRuleIds,omitempty" tf:"ingress_nat_rule_ids,omitempty"`

	// A ipsec_policy block which is documented below.
	// Only a single policy can be defined for a connection. For details on
	// custom policies refer to the relevant section in the Azure documentation.
	// +kubebuilder:validation:Optional
	IpsecPolicy []IpsecPolicyParameters `json:"ipsecPolicy,omitempty" tf:"ipsec_policy,omitempty"`

	// Use private local Azure IP for the connection. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	LocalAzureIPAddressEnabled *bool `json:"localAzureIpAddressEnabled,omitempty" tf:"local_azure_ip_address_enabled,omitempty"`

	// The ID of the local network gateway
	// when creating Site-to-Site connection .
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/network/v1beta1.LocalNetworkGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LocalNetworkGatewayID *string `json:"localNetworkGatewayId,omitempty" tf:"local_network_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	LocalNetworkGatewayIDRef *v1.Reference `json:"localNetworkGatewayIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LocalNetworkGatewayIDSelector *v1.Selector `json:"localNetworkGatewayIdSelector,omitempty" tf:"-"`

	// The location/region where the connection is
	// located. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The ID of the peer virtual
	// network gateway when creating a VNet-to-VNet connection . The peer Virtual Network Gateway can be in the same or
	// in a different subscription.
	// +crossplane:generate:reference:type=VirtualNetworkGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PeerVirtualNetworkGatewayID *string `json:"peerVirtualNetworkGatewayId,omitempty" tf:"peer_virtual_network_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	PeerVirtualNetworkGatewayIDRef *v1.Reference `json:"peerVirtualNetworkGatewayIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PeerVirtualNetworkGatewayIDSelector *v1.Selector `json:"peerVirtualNetworkGatewayIdSelector,omitempty" tf:"-"`

	// The name of the resource group in which to
	// create the connection Changing the name forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The routing weight. Defaults to 10.
	// +kubebuilder:validation:Optional
	RoutingWeight *float64 `json:"routingWeight,omitempty" tf:"routing_weight,omitempty"`

	// The shared IPSec key. A key could be provided if a
	// Site-to-Site, VNet-to-VNet or ExpressRoute connection is created.
	// +kubebuilder:validation:Optional
	SharedKeySecretRef *v1.SecretKeySelector `json:"sharedKeySecretRef,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more traffic_selector_policy blocks which are documented below.
	// A traffic_selector_policy allows to specify a traffic selector policy proposal to be used in a virtual network gateway connection.
	// For details about traffic selectors refer to the relevant section in the Azure documentation.
	// +kubebuilder:validation:Optional
	TrafficSelectorPolicy []TrafficSelectorPolicyParameters `json:"trafficSelectorPolicy,omitempty" tf:"traffic_selector_policy,omitempty"`

	// The type of connection. Valid options are IPsec
	// , ExpressRoute , and Vnet2Vnet .
	// Each connection type requires different mandatory arguments . Changing the connection type will force a new connection
	// to be created.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// If true, policy-based traffic
	// selectors are enabled for this connection. Enabling policy-based traffic
	// selectors requires an ipsec_policy block. Defaults to false.
	// +kubebuilder:validation:Optional
	UsePolicyBasedTrafficSelectors *bool `json:"usePolicyBasedTrafficSelectors,omitempty" tf:"use_policy_based_traffic_selectors,omitempty"`

	// The ID of the Virtual Network Gateway
	// in which the connection will be created. Changing the gateway forces a new
	// resource to be created.
	// +crossplane:generate:reference:type=VirtualNetworkGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualNetworkGatewayID *string `json:"virtualNetworkGatewayId,omitempty" tf:"virtual_network_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkGatewayIDRef *v1.Reference `json:"virtualNetworkGatewayIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VirtualNetworkGatewayIDSelector *v1.Selector `json:"virtualNetworkGatewayIdSelector,omitempty" tf:"-"`
}

// VirtualNetworkGatewayConnectionSpec defines the desired state of VirtualNetworkGatewayConnection
type VirtualNetworkGatewayConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkGatewayConnectionParameters `json:"forProvider"`
}

// VirtualNetworkGatewayConnectionStatus defines the observed state of VirtualNetworkGatewayConnection.
type VirtualNetworkGatewayConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkGatewayConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGatewayConnection is the Schema for the VirtualNetworkGatewayConnections API. Manages a connection in an existing Virtual Network Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetworkGatewayConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGatewayConnectionSpec   `json:"spec"`
	Status            VirtualNetworkGatewayConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGatewayConnectionList contains a list of VirtualNetworkGatewayConnections
type VirtualNetworkGatewayConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkGatewayConnection `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkGatewayConnection_Kind             = "VirtualNetworkGatewayConnection"
	VirtualNetworkGatewayConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualNetworkGatewayConnection_Kind}.String()
	VirtualNetworkGatewayConnection_KindAPIVersion   = VirtualNetworkGatewayConnection_Kind + "." + CRDGroupVersion.String()
	VirtualNetworkGatewayConnection_GroupVersionKind = CRDGroupVersion.WithKind(VirtualNetworkGatewayConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetworkGatewayConnection{}, &VirtualNetworkGatewayConnectionList{})
}
