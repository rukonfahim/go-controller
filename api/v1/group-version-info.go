package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{Group: "mygroup.com", Version: "v1"}

//func init() {
//	SchemeBuilder.Register(&MyResource{}, &MyResourceList{})
//}

func init() {
	SchemeBuilder.Register(addKnownTypes)
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&MyResource{},
		&MyResourceList{},
	)
	return nil
}