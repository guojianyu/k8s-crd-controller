package v1alpha1

import (
    meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodGroup describes a PodGroup resource
type LkController struct {
    // TypeMeta is the metadata for the resource, like kind and apiversion
    meta_v1.TypeMeta `json:",inline"`
    // ObjectMeta contains the metadata for the particular object, including
    // things like...
    //  - name
    //  - namespace
    //  - self link
    //  - labels
    //  - ... etc ...
    meta_v1.ObjectMeta `json:"metadata,omitempty"`

    // Spec is the custom resource spec
    Spec LkControllerSpec `json:"spec"`
}

// PodGroupSpec is the spec for a MyResource resource
type LkControllerSpec struct {
    Min int `json:"min"`
    Replicas       *int32 `json:"replicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodGroupList is a list of PodGroup resources
type LkControllerList struct {
    meta_v1.TypeMeta `json:",inline"`
    meta_v1.ListMeta `json:"metadata"`

    Items []LkController `json:"items"`
}
