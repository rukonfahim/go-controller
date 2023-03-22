package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime" // add this line
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

// DeepCopyObject creates a deep copy of this object and returns the new object.
func (in *MyResource) DeepCopyObject() runtime.Object {
	if in == nil {
		return nil
	}
	out := new(MyResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto creates a deep copy of this object and stores it in the specified object.
func (in *MyResource) DeepCopyInto(out *MyResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}


// DeepCopyObject creates a deep copy of this object and returns the new object.
func (in *MyResourceList) DeepCopyObject() runtime.Object {
	if in == nil {
		return nil
	}
	out := new(MyResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto creates a deep copy of this object and stores it in the specified object.
func (in *MyResourceList) DeepCopyInto(out *MyResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MyResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}