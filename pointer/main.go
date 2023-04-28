package main

import (
	"fmt"
)

type ParameterStruct struct {
	Name string
}

func changeParameter(param *ParameterStruct, value string) {
	param.Name = value
}

func cannotChangeParameter(param ParameterStruct, value string) {
	param.Name = value
}

/**
 * To passing a pointer or a value to a function
 */
func main() {
	// 1. pointer and value
	str := "a string value"
	pointer := &str
	anotherString := *&str
	fmt.Println(str)            // a string value
	fmt.Println(pointer)        // 0xc000096230
	fmt.Println(anotherString)  // a string value

	str = "changed string"      
	fmt.Println(str)            // changed string
	fmt.Println(pointer)        // 0xc000096230 - pointer do not change
	fmt.Println(anotherString)  // a string value - pointer value does no change too

}