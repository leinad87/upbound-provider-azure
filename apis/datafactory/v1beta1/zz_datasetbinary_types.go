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

type AzureBlobStorageLocationInitParameters struct {

	// The container on the Azure Blob Storage Account hosting the file.
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// Is the container using dynamic expression, function or system variables? Defaults to false.
	DynamicContainerEnabled *bool `json:"dynamicContainerEnabled,omitempty" tf:"dynamic_container_enabled,omitempty"`

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file in the blob container.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file in the blob container.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type AzureBlobStorageLocationObservation struct {

	// The container on the Azure Blob Storage Account hosting the file.
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// Is the container using dynamic expression, function or system variables? Defaults to false.
	DynamicContainerEnabled *bool `json:"dynamicContainerEnabled,omitempty" tf:"dynamic_container_enabled,omitempty"`

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file in the blob container.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file in the blob container.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type AzureBlobStorageLocationParameters struct {

	// The container on the Azure Blob Storage Account hosting the file.
	// +kubebuilder:validation:Optional
	Container *string `json:"container" tf:"container,omitempty"`

	// Is the container using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicContainerEnabled *bool `json:"dynamicContainerEnabled,omitempty" tf:"dynamic_container_enabled,omitempty"`

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file in the blob container.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file in the blob container.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type CompressionInitParameters struct {

	// The level of compression. Possible values are Fastest and Optimal.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// The type of compression used during transport. Possible values are BZip2, Deflate, GZip, Tar, TarGZip and ZipDeflate.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CompressionObservation struct {

	// The level of compression. Possible values are Fastest and Optimal.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// The type of compression used during transport. Possible values are BZip2, Deflate, GZip, Tar, TarGZip and ZipDeflate.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CompressionParameters struct {

	// The level of compression. Possible values are Fastest and Optimal.
	// +kubebuilder:validation:Optional
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// The type of compression used during transport. Possible values are BZip2, Deflate, GZip, Tar, TarGZip and ZipDeflate.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type DataSetBinaryInitParameters struct {

	// A map of additional properties to associate with the Data Factory Binary Dataset.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Binary Dataset.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A azure_blob_storage_location block as defined below.
	AzureBlobStorageLocation []AzureBlobStorageLocationInitParameters `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

	// A compression block as defined below.
	Compression []CompressionInitParameters `json:"compression,omitempty" tf:"compression,omitempty"`

	// The description for the Data Factory Dataset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A http_server_location block as defined below.
	HTTPServerLocation []HTTPServerLocationInitParameters `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// The Data Factory Linked Service name in which to associate the Binary Dataset with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceSFTP
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceSFTP in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceSFTP in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// Specifies a list of parameters to associate with the Data Factory Binary Dataset.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A sftp_server_location block as defined below.
	SFTPServerLocation []SFTPServerLocationInitParameters `json:"sftpServerLocation,omitempty" tf:"sftp_server_location,omitempty"`
}

type DataSetBinaryObservation struct {

	// A map of additional properties to associate with the Data Factory Binary Dataset.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Binary Dataset.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A azure_blob_storage_location block as defined below.
	AzureBlobStorageLocation []AzureBlobStorageLocationObservation `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

	// A compression block as defined below.
	Compression []CompressionObservation `json:"compression,omitempty" tf:"compression,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Dataset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A http_server_location block as defined below.
	HTTPServerLocation []HTTPServerLocationObservation `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// The ID of the Data Factory Dataset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Data Factory Linked Service name in which to associate the Binary Dataset with.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Specifies a list of parameters to associate with the Data Factory Binary Dataset.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A sftp_server_location block as defined below.
	SFTPServerLocation []SFTPServerLocationObservation `json:"sftpServerLocation,omitempty" tf:"sftp_server_location,omitempty"`
}

type DataSetBinaryParameters struct {

	// A map of additional properties to associate with the Data Factory Binary Dataset.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Binary Dataset.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A azure_blob_storage_location block as defined below.
	// +kubebuilder:validation:Optional
	AzureBlobStorageLocation []AzureBlobStorageLocationParameters `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

	// A compression block as defined below.
	// +kubebuilder:validation:Optional
	Compression []CompressionParameters `json:"compression,omitempty" tf:"compression,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A http_server_location block as defined below.
	// +kubebuilder:validation:Optional
	HTTPServerLocation []HTTPServerLocationParameters `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// The Data Factory Linked Service name in which to associate the Binary Dataset with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceSFTP
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceSFTP in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceSFTP in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// Specifies a list of parameters to associate with the Data Factory Binary Dataset.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A sftp_server_location block as defined below.
	// +kubebuilder:validation:Optional
	SFTPServerLocation []SFTPServerLocationParameters `json:"sftpServerLocation,omitempty" tf:"sftp_server_location,omitempty"`
}

type HTTPServerLocationInitParameters struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The base URL to the web server hosting the file.
	RelativeURL *string `json:"relativeUrl,omitempty" tf:"relative_url,omitempty"`
}

type HTTPServerLocationObservation struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The base URL to the web server hosting the file.
	RelativeURL *string `json:"relativeUrl,omitempty" tf:"relative_url,omitempty"`
}

type HTTPServerLocationParameters struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`

	// The base URL to the web server hosting the file.
	// +kubebuilder:validation:Optional
	RelativeURL *string `json:"relativeUrl" tf:"relative_url,omitempty"`
}

type SFTPServerLocationInitParameters struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the SFTP server.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file on the SFTP server.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type SFTPServerLocationObservation struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the SFTP server.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file on the SFTP server.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type SFTPServerLocationParameters struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the SFTP server.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename" tf:"filename,omitempty"`

	// The folder path to the file on the SFTP server.
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`
}

// DataSetBinarySpec defines the desired state of DataSetBinary
type DataSetBinarySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetBinaryParameters `json:"forProvider"`
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
	InitProvider DataSetBinaryInitParameters `json:"initProvider,omitempty"`
}

// DataSetBinaryStatus defines the observed state of DataSetBinary.
type DataSetBinaryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetBinaryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetBinary is the Schema for the DataSetBinarys API. Manages a Data Factory Binary Dataset inside an Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetBinary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetBinarySpec   `json:"spec"`
	Status            DataSetBinaryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetBinaryList contains a list of DataSetBinarys
type DataSetBinaryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetBinary `json:"items"`
}

// Repository type metadata.
var (
	DataSetBinary_Kind             = "DataSetBinary"
	DataSetBinary_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetBinary_Kind}.String()
	DataSetBinary_KindAPIVersion   = DataSetBinary_Kind + "." + CRDGroupVersion.String()
	DataSetBinary_GroupVersionKind = CRDGroupVersion.WithKind(DataSetBinary_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetBinary{}, &DataSetBinaryList{})
}
