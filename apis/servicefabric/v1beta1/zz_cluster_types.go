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

type ApplicationPortsObservation struct {
}

type ApplicationPortsParameters struct {

	// The end of the Ephemeral Port Range on this Node Type.
	// +kubebuilder:validation:Required
	EndPort *float64 `json:"endPort" tf:"end_port,omitempty"`

	// The start of the Ephemeral Port Range on this Node Type.
	// +kubebuilder:validation:Required
	StartPort *float64 `json:"startPort" tf:"start_port,omitempty"`
}

type AzureActiveDirectoryObservation struct {
}

type AzureActiveDirectoryParameters struct {

	// The Azure Active Directory Client ID which should be used for the Client Application.
	// +kubebuilder:validation:Required
	ClientApplicationID *string `json:"clientApplicationId" tf:"client_application_id,omitempty"`

	// The Azure Active Directory Cluster Application ID.
	// +kubebuilder:validation:Required
	ClusterApplicationID *string `json:"clusterApplicationId" tf:"cluster_application_id,omitempty"`

	// The Azure Active Directory Tenant ID.
	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

type CertificateCommonNamesObservation struct {
}

type CertificateCommonNamesParameters struct {

	// A common_names block as defined below.
	// +kubebuilder:validation:Required
	CommonNames []CommonNamesParameters `json:"commonNames" tf:"common_names,omitempty"`

	// The X509 Store where the Certificate Exists, such as My.
	// +kubebuilder:validation:Required
	X509StoreName *string `json:"x509StoreName" tf:"x509_store_name,omitempty"`
}

type CertificateObservation struct {
}

type CertificateParameters struct {

	// The Thumbprint of the Certificate.
	// +kubebuilder:validation:Required
	Thumbprint *string `json:"thumbprint" tf:"thumbprint,omitempty"`

	// The Secondary Thumbprint of the Certificate.
	// +kubebuilder:validation:Optional
	ThumbprintSecondary *string `json:"thumbprintSecondary,omitempty" tf:"thumbprint_secondary,omitempty"`

	// The X509 Store where the Certificate Exists, such as My.
	// +kubebuilder:validation:Required
	X509StoreName *string `json:"x509StoreName" tf:"x509_store_name,omitempty"`
}

type ClientCertificateCommonNameObservation struct {
}

type ClientCertificateCommonNameParameters struct {

	// The common or subject name of the certificate.
	// +kubebuilder:validation:Required
	CommonName *string `json:"commonName" tf:"common_name,omitempty"`

	// Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
	// +kubebuilder:validation:Required
	IsAdmin *bool `json:"isAdmin" tf:"is_admin,omitempty"`

	// The Issuer Thumbprint of the Certificate.
	// +kubebuilder:validation:Optional
	IssuerThumbprint *string `json:"issuerThumbprint,omitempty" tf:"issuer_thumbprint,omitempty"`
}

type ClientCertificateThumbprintObservation struct {
}

type ClientCertificateThumbprintParameters struct {

	// Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
	// +kubebuilder:validation:Required
	IsAdmin *bool `json:"isAdmin" tf:"is_admin,omitempty"`

	// The Thumbprint associated with the Client Certificate.
	// +kubebuilder:validation:Required
	Thumbprint *string `json:"thumbprint" tf:"thumbprint,omitempty"`
}

type ClusterObservation struct {

	// The Cluster Endpoint for this Service Fabric Cluster.
	ClusterEndpoint *string `json:"clusterEndpoint,omitempty" tf:"cluster_endpoint,omitempty"`

	// The ID of the Service Fabric Cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ClusterParameters struct {

	// A List of one or more features which should be enabled, such as DnsService.
	// +kubebuilder:validation:Optional
	AddOnFeatures []*string `json:"addOnFeatures,omitempty" tf:"add_on_features,omitempty"`

	// An azure_active_directory block as defined below.
	// +kubebuilder:validation:Optional
	AzureActiveDirectory []AzureActiveDirectoryParameters `json:"azureActiveDirectory,omitempty" tf:"azure_active_directory,omitempty"`

	// A certificate block as defined below. Conflicts with certificate_common_names.
	// +kubebuilder:validation:Optional
	Certificate []CertificateParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// A certificate_common_names block as defined below. Conflicts with certificate.
	// +kubebuilder:validation:Optional
	CertificateCommonNames []CertificateCommonNamesParameters `json:"certificateCommonNames,omitempty" tf:"certificate_common_names,omitempty"`

	// A client_certificate_common_name block as defined below.
	// +kubebuilder:validation:Optional
	ClientCertificateCommonName []ClientCertificateCommonNameParameters `json:"clientCertificateCommonName,omitempty" tf:"client_certificate_common_name,omitempty"`

	// One or more client_certificate_thumbprint blocks as defined below.
	// +kubebuilder:validation:Optional
	ClientCertificateThumbprint []ClientCertificateThumbprintParameters `json:"clientCertificateThumbprint,omitempty" tf:"client_certificate_thumbprint,omitempty"`

	// Required if Upgrade Mode set to Manual, Specifies the Version of the Cluster Code of the cluster.
	// +kubebuilder:validation:Optional
	ClusterCodeVersion *string `json:"clusterCodeVersion,omitempty" tf:"cluster_code_version,omitempty"`

	// A diagnostics_config block as defined below.
	// +kubebuilder:validation:Optional
	DiagnosticsConfig []DiagnosticsConfigParameters `json:"diagnosticsConfig,omitempty" tf:"diagnostics_config,omitempty"`

	// One or more fabric_settings blocks as defined below.
	// +kubebuilder:validation:Optional
	FabricSettings []FabricSettingsParameters `json:"fabricSettings,omitempty" tf:"fabric_settings,omitempty"`

	// Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the Management Endpoint of the cluster such as http://example.com. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	ManagementEndpoint *string `json:"managementEndpoint" tf:"management_endpoint,omitempty"`

	// One or more node_type blocks as defined below.
	// +kubebuilder:validation:Required
	NodeType []NodeTypeParameters `json:"nodeType" tf:"node_type,omitempty"`

	// Specifies the Reliability Level of the Cluster. Possible values include None, Bronze, Silver, Gold and Platinum.
	// +kubebuilder:validation:Required
	ReliabilityLevel *string `json:"reliabilityLevel" tf:"reliability_level,omitempty"`

	// The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A reverse_proxy_certificate block as defined below. Conflicts with reverse_proxy_certificate_common_names.
	// +kubebuilder:validation:Optional
	ReverseProxyCertificate []ReverseProxyCertificateParameters `json:"reverseProxyCertificate,omitempty" tf:"reverse_proxy_certificate,omitempty"`

	// A reverse_proxy_certificate_common_names block as defined below. Conflicts with reverse_proxy_certificate.
	// +kubebuilder:validation:Optional
	ReverseProxyCertificateCommonNames []ReverseProxyCertificateCommonNamesParameters `json:"reverseProxyCertificateCommonNames,omitempty" tf:"reverse_proxy_certificate_common_names,omitempty"`

	// Specifies the logical grouping of VMs in upgrade domains. Possible values are Hierarchical or Parallel.
	// +kubebuilder:validation:Optional
	ServiceFabricZonalUpgradeMode *string `json:"serviceFabricZonalUpgradeMode,omitempty" tf:"service_fabric_zonal_upgrade_mode,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Upgrade Mode of the cluster. Possible values are Automatic or Manual.
	// +kubebuilder:validation:Required
	UpgradeMode *string `json:"upgradeMode" tf:"upgrade_mode,omitempty"`

	// A upgrade_policy block as defined below.
	// +kubebuilder:validation:Optional
	UpgradePolicy []UpgradePolicyParameters `json:"upgradePolicy,omitempty" tf:"upgrade_policy,omitempty"`

	// Specifies the Image expected for the Service Fabric Cluster, such as Windows. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	VMImage *string `json:"vmImage" tf:"vm_image,omitempty"`

	// Specifies the upgrade mode for the virtual machine scale set updates that happen in all availability zones at once. Possible values are Hierarchical or Parallel.
	// +kubebuilder:validation:Optional
	VmssZonalUpgradeMode *string `json:"vmssZonalUpgradeMode,omitempty" tf:"vmss_zonal_upgrade_mode,omitempty"`
}

type CommonNamesObservation struct {
}

type CommonNamesParameters struct {

	// The common or subject name of the certificate.
	// +kubebuilder:validation:Required
	CertificateCommonName *string `json:"certificateCommonName" tf:"certificate_common_name,omitempty"`

	// The Issuer Thumbprint of the Certificate.
	// +kubebuilder:validation:Optional
	CertificateIssuerThumbprint *string `json:"certificateIssuerThumbprint,omitempty" tf:"certificate_issuer_thumbprint,omitempty"`
}

type DeltaHealthPolicyObservation struct {
}

type DeltaHealthPolicyParameters struct {

	// Specifies the maximum tolerated percentage of delta unhealthy applications that can have aggregated health states of error. If the current unhealthy applications do not respect the percentage relative to the state at the beginning of the upgrade, the cluster is unhealthy. Defaults to 0.
	// +kubebuilder:validation:Optional
	MaxDeltaUnhealthyApplicationsPercent *float64 `json:"maxDeltaUnhealthyApplicationsPercent,omitempty" tf:"max_delta_unhealthy_applications_percent,omitempty"`

	// Specifies the maximum tolerated percentage of delta unhealthy nodes that can have aggregated health states of error. If the current unhealthy nodes do not respect the percentage relative to the state at the beginning of the upgrade, the cluster is unhealthy. Defaults to 0.
	// +kubebuilder:validation:Optional
	MaxDeltaUnhealthyNodesPercent *float64 `json:"maxDeltaUnhealthyNodesPercent,omitempty" tf:"max_delta_unhealthy_nodes_percent,omitempty"`

	// Specifies the maximum tolerated percentage of upgrade domain delta unhealthy nodes that can have aggregated health state of error. If there is any upgrade domain where the current unhealthy nodes do not respect the percentage relative to the state at the beginning of the upgrade, the cluster is unhealthy. Defaults to 0.
	// +kubebuilder:validation:Optional
	MaxUpgradeDomainDeltaUnhealthyNodesPercent *float64 `json:"maxUpgradeDomainDeltaUnhealthyNodesPercent,omitempty" tf:"max_upgrade_domain_delta_unhealthy_nodes_percent,omitempty"`
}

type DiagnosticsConfigObservation struct {
}

type DiagnosticsConfigParameters struct {

	// The Blob Endpoint of the Storage Account.
	// +kubebuilder:validation:Required
	BlobEndpoint *string `json:"blobEndpoint" tf:"blob_endpoint,omitempty"`

	// The protected diagnostics storage key name, such as StorageAccountKey1.
	// +kubebuilder:validation:Required
	ProtectedAccountKeyName *string `json:"protectedAccountKeyName" tf:"protected_account_key_name,omitempty"`

	// The Queue Endpoint of the Storage Account.
	// +kubebuilder:validation:Required
	QueueEndpoint *string `json:"queueEndpoint" tf:"queue_endpoint,omitempty"`

	// The name of the Storage Account where the Diagnostics should be sent to.
	// +kubebuilder:validation:Required
	StorageAccountName *string `json:"storageAccountName" tf:"storage_account_name,omitempty"`

	// The Table Endpoint of the Storage Account.
	// +kubebuilder:validation:Required
	TableEndpoint *string `json:"tableEndpoint" tf:"table_endpoint,omitempty"`
}

type EphemeralPortsObservation struct {
}

type EphemeralPortsParameters struct {

	// The end of the Ephemeral Port Range on this Node Type.
	// +kubebuilder:validation:Required
	EndPort *float64 `json:"endPort" tf:"end_port,omitempty"`

	// The start of the Ephemeral Port Range on this Node Type.
	// +kubebuilder:validation:Required
	StartPort *float64 `json:"startPort" tf:"start_port,omitempty"`
}

type FabricSettingsObservation struct {
}

type FabricSettingsParameters struct {

	// The name of the Fabric Setting, such as Security or Federation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map containing settings for the specified Fabric Setting.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type HealthPolicyObservation struct {
}

type HealthPolicyParameters struct {

	// Specifies the maximum tolerated percentage of applications that can have aggregated health state of error. If the upgrade exceeds this percentage, the cluster is unhealthy. Defaults to 0.
	// +kubebuilder:validation:Optional
	MaxUnhealthyApplicationsPercent *float64 `json:"maxUnhealthyApplicationsPercent,omitempty" tf:"max_unhealthy_applications_percent,omitempty"`

	// Specifies the maximum tolerated percentage of nodes that can have aggregated health states of error. If an upgrade exceeds this percentage, the cluster is unhealthy. Defaults to 0.
	// +kubebuilder:validation:Optional
	MaxUnhealthyNodesPercent *float64 `json:"maxUnhealthyNodesPercent,omitempty" tf:"max_unhealthy_nodes_percent,omitempty"`
}

type NodeTypeObservation struct {
}

type NodeTypeParameters struct {

	// A application_ports block as defined below.
	// +kubebuilder:validation:Optional
	ApplicationPorts []ApplicationPortsParameters `json:"applicationPorts,omitempty" tf:"application_ports,omitempty"`

	// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
	// +kubebuilder:validation:Optional
	Capacities map[string]*string `json:"capacities,omitempty" tf:"capacities,omitempty"`

	// The Port used for the Client Endpoint for this Node Type.
	// +kubebuilder:validation:Required
	ClientEndpointPort *float64 `json:"clientEndpointPort" tf:"client_endpoint_port,omitempty"`

	// The Durability Level for this Node Type. Possible values include Bronze, Gold and Silver. Defaults to Bronze.
	// +kubebuilder:validation:Optional
	DurabilityLevel *string `json:"durabilityLevel,omitempty" tf:"durability_level,omitempty"`

	// A ephemeral_ports block as defined below.
	// +kubebuilder:validation:Optional
	EphemeralPorts []EphemeralPortsParameters `json:"ephemeralPorts,omitempty" tf:"ephemeral_ports,omitempty"`

	// The Port used for the HTTP Endpoint for this Node Type.
	// +kubebuilder:validation:Required
	HTTPEndpointPort *float64 `json:"httpEndpointPort" tf:"http_endpoint_port,omitempty"`

	// The number of nodes for this Node Type.
	// +kubebuilder:validation:Required
	InstanceCount *float64 `json:"instanceCount" tf:"instance_count,omitempty"`

	// Is this the Primary Node Type?
	// +kubebuilder:validation:Required
	IsPrimary *bool `json:"isPrimary" tf:"is_primary,omitempty"`

	// Should this node type run only stateless services?
	// +kubebuilder:validation:Optional
	IsStateless *bool `json:"isStateless,omitempty" tf:"is_stateless,omitempty"`

	// Does this node type span availability zones?
	// +kubebuilder:validation:Optional
	MultipleAvailabilityZones *bool `json:"multipleAvailabilityZones,omitempty" tf:"multiple_availability_zones,omitempty"`

	// The name of the Node Type.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
	// +kubebuilder:validation:Optional
	PlacementProperties map[string]*string `json:"placementProperties,omitempty" tf:"placement_properties,omitempty"`

	// The Port used for the Reverse Proxy Endpoint for this Node Type. Changing this will upgrade the cluster.
	// +kubebuilder:validation:Optional
	ReverseProxyEndpointPort *float64 `json:"reverseProxyEndpointPort,omitempty" tf:"reverse_proxy_endpoint_port,omitempty"`
}

type ReverseProxyCertificateCommonNamesCommonNamesObservation struct {
}

type ReverseProxyCertificateCommonNamesCommonNamesParameters struct {

	// The common or subject name of the certificate.
	// +kubebuilder:validation:Required
	CertificateCommonName *string `json:"certificateCommonName" tf:"certificate_common_name,omitempty"`

	// The Issuer Thumbprint of the Certificate.
	// +kubebuilder:validation:Optional
	CertificateIssuerThumbprint *string `json:"certificateIssuerThumbprint,omitempty" tf:"certificate_issuer_thumbprint,omitempty"`
}

type ReverseProxyCertificateCommonNamesObservation struct {
}

type ReverseProxyCertificateCommonNamesParameters struct {

	// A common_names block as defined below.
	// +kubebuilder:validation:Required
	CommonNames []ReverseProxyCertificateCommonNamesCommonNamesParameters `json:"commonNames" tf:"common_names,omitempty"`

	// The X509 Store where the Certificate Exists, such as My.
	// +kubebuilder:validation:Required
	X509StoreName *string `json:"x509StoreName" tf:"x509_store_name,omitempty"`
}

type ReverseProxyCertificateObservation struct {
}

type ReverseProxyCertificateParameters struct {

	// The Thumbprint of the Certificate.
	// +kubebuilder:validation:Required
	Thumbprint *string `json:"thumbprint" tf:"thumbprint,omitempty"`

	// The Secondary Thumbprint of the Certificate.
	// +kubebuilder:validation:Optional
	ThumbprintSecondary *string `json:"thumbprintSecondary,omitempty" tf:"thumbprint_secondary,omitempty"`

	// The X509 Store where the Certificate Exists, such as My.
	// +kubebuilder:validation:Required
	X509StoreName *string `json:"x509StoreName" tf:"x509_store_name,omitempty"`
}

type UpgradePolicyObservation struct {
}

type UpgradePolicyParameters struct {

	// A delta_health_policy block as defined below
	// +kubebuilder:validation:Optional
	DeltaHealthPolicy []DeltaHealthPolicyParameters `json:"deltaHealthPolicy,omitempty" tf:"delta_health_policy,omitempty"`

	// Indicates whether to restart the Service Fabric node even if only dynamic configurations have changed.
	// +kubebuilder:validation:Optional
	ForceRestartEnabled *bool `json:"forceRestartEnabled,omitempty" tf:"force_restart_enabled,omitempty"`

	// Specifies the duration, in "hh:mm:ss" string format, after which Service Fabric retries the health check if the previous health check fails. Defaults to 00:45:00.
	// +kubebuilder:validation:Optional
	HealthCheckRetryTimeout *string `json:"healthCheckRetryTimeout,omitempty" tf:"health_check_retry_timeout,omitempty"`

	// Specifies the duration, in "hh:mm:ss" string format, that Service Fabric waits in order to verify that the cluster is stable before it continues to the next upgrade domain or completes the upgrade. This wait duration prevents undetected changes of health right after the health check is performed. Defaults to 00:01:00.
	// +kubebuilder:validation:Optional
	HealthCheckStableDuration *string `json:"healthCheckStableDuration,omitempty" tf:"health_check_stable_duration,omitempty"`

	// Specifies the duration, in "hh:mm:ss" string format, that Service Fabric waits before it performs the initial health check after it finishes the upgrade on the upgrade domain. Defaults to 00:00:30.
	// +kubebuilder:validation:Optional
	HealthCheckWaitDuration *string `json:"healthCheckWaitDuration,omitempty" tf:"health_check_wait_duration,omitempty"`

	// A health_policy block as defined below
	// +kubebuilder:validation:Optional
	HealthPolicy []HealthPolicyParameters `json:"healthPolicy,omitempty" tf:"health_policy,omitempty"`

	// Specifies the duration, in "hh:mm:ss" string format, that Service Fabric takes to upgrade a single upgrade domain. After this period, the upgrade fails. Defaults to 02:00:00.
	// +kubebuilder:validation:Optional
	UpgradeDomainTimeout *string `json:"upgradeDomainTimeout,omitempty" tf:"upgrade_domain_timeout,omitempty"`

	// Specifies the duration, in "hh:mm:ss" string format, that Service Fabric waits for a replica set to reconfigure into a safe state, if it is not already in a safe state, before Service Fabric proceeds with the upgrade. Defaults to 10675199.02:48:05.4775807.
	// +kubebuilder:validation:Optional
	UpgradeReplicaSetCheckTimeout *string `json:"upgradeReplicaSetCheckTimeout,omitempty" tf:"upgrade_replica_set_check_timeout,omitempty"`

	// Specifies the duration, in "hh:mm:ss" string format, that Service Fabric takes for the entire upgrade. After this period, the upgrade fails. Defaults to 12:00:00.
	// +kubebuilder:validation:Optional
	UpgradeTimeout *string `json:"upgradeTimeout,omitempty" tf:"upgrade_timeout,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. Manages a Service Fabric Cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
