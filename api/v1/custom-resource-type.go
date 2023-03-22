package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MyResourceSpec defines the desired state of MyResource
type MyResourceSpec struct {
	Image string `json:"image"`
	Port  int    `json:"port"`
}

// MyResource is the Schema for the myresources API
type MyResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MyResourceSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// MyResourceList contains a list of MyResource
type MyResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyResource `json:"items"`
}
