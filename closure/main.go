package main

import (
	"fmt"
)

func main() {
	// 1. function can be assigned to a variable
	x := func(){}
	fmt.Println(x)

	// 2. function can be invoked directly
	func(x, y int){
		fmt.Println(x + y) // 3
	}(1, 2)

	// 3. function can be a returned value - to be checked
	// func Add() (func(b int) int)

	// 4. defer function
	defer func() {
		if r := recover(); r != nil {
			println("recovered in funcX")
		}
	}()
}

