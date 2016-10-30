package main

import (
	"encoding/json"

	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/pkg/api/unversioned"
	"k8s.io/client-go/1.4/pkg/runtime"
)

type DemoSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Demo struct {
	unversioned.TypeMeta `json:",inline"`
	api.ObjectMeta       `json:"metadata,omitempty"`

	Spec DemoSpec `json:"spec"`
}

type DemoList struct {
	unversioned.TypeMeta `json:",inline"`
	api.ObjectMeta       `json:"metadata,omitempty"`

	Items []Demo `json:"items"`
}

const GroupName = "third.com"

var (
	SchemeGroupVersion = unversioned.GroupVersion{Group: GroupName, Version: "v1"}
	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme        = SchemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	// TODO this gets cleaned up when the types are fixed
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Demo{},
		&DemoList{},
	)
	return nil
}

func init() {
	if err := AddToScheme(api.Scheme); err != nil {
		panic(err)
	}
}

func (d *Demo) Decode(data []byte) error {
	return json.Unmarshal(data, *d)
}
