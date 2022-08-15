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

type AudioAnalyzerPresetObservation struct {
}

type AudioAnalyzerPresetParameters struct {

	// +kubebuilder:validation:Optional
	AudioAnalysisMode *string `json:"audioAnalysisMode,omitempty" tf:"audio_analysis_mode,omitempty"`

	// +kubebuilder:validation:Optional
	AudioLanguage *string `json:"audioLanguage,omitempty" tf:"audio_language,omitempty"`
}

type BuiltinPresetObservation struct {
}

type BuiltinPresetParameters struct {

	// The built-in preset to be used for encoding videos. The allowed values are AACGoodQualityAudio, AdaptiveStreaming,ContentAwareEncoding, ContentAwareEncodingExperimental,CopyAllBitrateNonInterleaved, H264MultipleBitrate1080p,H264MultipleBitrate720p, H264MultipleBitrateSD,H264SingleBitrate1080p, H264SingleBitrate720p and H264SingleBitrateSD.
	// +kubebuilder:validation:Optional
	PresetName *string `json:"presetName,omitempty" tf:"preset_name,omitempty"`
}

type FaceDetectorPresetObservation struct {
}

type FaceDetectorPresetParameters struct {

	// Possibles value are SourceResolution or StandardDefinition. Specifies the maximum resolution at which your video is analyzed. The default behavior is SourceResolution which will keep the input video at its original resolution when analyzed. Using StandardDefinition will resize input videos to standard definition while preserving the appropriate aspect ratio. It will only resize if the video is of higher resolution. For example, a 1920x1080 input would be scaled to 640x360 before processing. Switching to StandardDefinition will reduce the time it takes to process high resolution video. It may also reduce the cost of using this component . However, faces that end up being too small in the resized video may not be detected.
	// +kubebuilder:validation:Optional
	AnalysisResolution *string `json:"analysisResolution,omitempty" tf:"analysis_resolution,omitempty"`
}

type OutputObservation struct {
}

type OutputParameters struct {

	// A audio_analyzer_preset block as defined below.
	// +kubebuilder:validation:Optional
	AudioAnalyzerPreset []AudioAnalyzerPresetParameters `json:"audioAnalyzerPreset,omitempty" tf:"audio_analyzer_preset,omitempty"`

	// A builtin_preset block as defined below.
	// +kubebuilder:validation:Optional
	BuiltinPreset []BuiltinPresetParameters `json:"builtinPreset,omitempty" tf:"builtin_preset,omitempty"`

	// A face_detector_preset block as defined below.
	// +kubebuilder:validation:Optional
	FaceDetectorPreset []FaceDetectorPresetParameters `json:"faceDetectorPreset,omitempty" tf:"face_detector_preset,omitempty"`

	// A Transform can define more than one outputs. This property defines what the service should do when one output fails - either continue to produce other outputs, or, stop the other outputs. The overall Job state will not reflect failures of outputs that are specified with ContinueJob. Possibles value are StopProcessingJob or ContinueJob.
	// +kubebuilder:validation:Optional
	OnErrorAction *string `json:"onErrorAction,omitempty" tf:"on_error_action,omitempty"`

	// Sets the relative priority of the TransformOutputs within a Transform. This sets the priority that the service uses for processing Transform Outputs. Possibles value are High, Normal or Low.
	// +kubebuilder:validation:Optional
	RelativePriority *string `json:"relativePriority,omitempty" tf:"relative_priority,omitempty"`

	// A video_analyzer_preset block as defined below.
	// +kubebuilder:validation:Optional
	VideoAnalyzerPreset []VideoAnalyzerPresetParameters `json:"videoAnalyzerPreset,omitempty" tf:"video_analyzer_preset,omitempty"`
}

type TransformObservation struct {

	// The ID of the Transform.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TransformParameters struct {

	// An optional verbose description of the Transform.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Media Services account name. Changing this forces a new Transform to be created.
	// +crossplane:generate:reference:type=ServicesAccount
	// +kubebuilder:validation:Optional
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	MediaServicesAccountNameRef *v1.Reference `json:"mediaServicesAccountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	MediaServicesAccountNameSelector *v1.Selector `json:"mediaServicesAccountNameSelector,omitempty" tf:"-"`

	// One or more output blocks as defined below. At least one output must be defined.
	// +kubebuilder:validation:Optional
	Output []OutputParameters `json:"output,omitempty" tf:"output,omitempty"`

	// The name of the Resource Group where the Transform should exist. Changing this forces a new Transform to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

type VideoAnalyzerPresetObservation struct {
}

type VideoAnalyzerPresetParameters struct {

	// +kubebuilder:validation:Optional
	AudioAnalysisMode *string `json:"audioAnalysisMode,omitempty" tf:"audio_analysis_mode,omitempty"`

	// +kubebuilder:validation:Optional
	AudioLanguage *string `json:"audioLanguage,omitempty" tf:"audio_language,omitempty"`

	// Defines the type of insights that you want the service to generate. The allowed values are AudioInsightsOnly, VideoInsightsOnly, and AllInsights. If you set this to AllInsights and the input is audio only, then only audio insights are generated. Similarly if the input is video only, then only video insights are generated. It is recommended that you not use AudioInsightsOnly if you expect some of your inputs to be video only; or use VideoInsightsOnly if you expect some of your inputs to be audio only. Your Jobs in such conditions would error out.
	// +kubebuilder:validation:Optional
	InsightsType *string `json:"insightsType,omitempty" tf:"insights_type,omitempty"`
}

// TransformSpec defines the desired state of Transform
type TransformSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransformParameters `json:"forProvider"`
}

// TransformStatus defines the observed state of Transform.
type TransformStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransformObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Transform is the Schema for the Transforms API. Manages a Transform.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Transform struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransformSpec   `json:"spec"`
	Status            TransformStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransformList contains a list of Transforms
type TransformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Transform `json:"items"`
}

// Repository type metadata.
var (
	Transform_Kind             = "Transform"
	Transform_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Transform_Kind}.String()
	Transform_KindAPIVersion   = Transform_Kind + "." + CRDGroupVersion.String()
	Transform_GroupVersionKind = CRDGroupVersion.WithKind(Transform_Kind)
)

func init() {
	SchemeBuilder.Register(&Transform{}, &TransformList{})
}
