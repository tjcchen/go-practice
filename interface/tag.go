package main

import (
	"reflect"
)

/**
 * Kubernetes APIServer define all of its resources with Json tag and ProtoBuff tag
 * Eg: NodeName string `json:"nodeName,omitempty" protobuf:"bytes,10,opt,name=nodeName"`
 */
type MyType1 struct {
	Name string `json:"name"`
}

func main() {
	mt := MyType1{Name: "test"}

	myType := reflect.TypeOf(mt) // reflection
	name := myType.Field(0)
	tag := name.Tag.Get("json")

	println(tag) // name
}