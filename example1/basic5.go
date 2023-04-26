package main

import (
	"fmt"
)

func main() {
	// simple - string, string map
	myMap := make(map[string]string, 10) // exceed length of 10, does not cast error
	myMap["a"] = "b"
	myMap["c"] = "d"
	myMap["e"] = "f"

	fmt.Println(myMap["a"]) // b
	fmt.Println(myMap["e"]) // f

	// complex - string, function map
	myFuncMap := map[string]func() int {
		"funcA": func() int {
			return 1
		},
	}
	fmt.Println(myFuncMap) // map[funcA:0x108dc60]

	// assess myFuncMap value
	f := myFuncMap["funcA"]
	fmt.Println(f())       // 1

	// use the exists parameter to check whether a specific key exists
	value, exists := myMap["a"]
	if exists {
		println(value) // b
	}

	// iterate over myMap
	for k, v := range myMap {
		println(k, v) // order is not fixed
	}
}