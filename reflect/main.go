package main

import (
	"reflect"
	"fmt"
)

type T struct {
	A string
}

// please note receiver is struct
func (t T) String() string {
	return t.A + "1"
}

func main() {
	// reflect.TypeOf() returns an object type
	// reflect.ValueOf() returns an object value

	// 1. basic type
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	myMap["c"] = "d"

	t := reflect.TypeOf(myMap)
	fmt.Println("type: ", t) // type:  map[string]string

	v := reflect.ValueOf(myMap)
	fmt.Println("value", v)  // value map[a:b c:d]

	// 2. struct
	myStruct := T{A: "a"}
	t1 := reflect.TypeOf(myStruct)
	v1 := reflect.ValueOf(myStruct)
	fmt.Println(t1)          // main.T - type
	fmt.Println(v1)          // a1 - value
	for i := 0; i < v1.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, v1.Field(i))   // Field 0: a
	}
	for i := 0; i < v1.NumMethod(); i++ {
		fmt.Printf("Method %d: %v\n", i, v1.Method(i)) // Method 0: 0x1090de0
	}
	// invoke the struct method String()
	result := v1.Method(0).Call(nil)
	fmt.Println("result:", result) // result: [a1]
}