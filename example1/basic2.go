package main

import (
	"fmt"
)

func main() {
	// variables
	var num1 int = 10
	var i, j int = 3, 7

	// constants
	const flag bool = false

	// declare and assignment
	retVal := "result"
	c, python, java := true, false, "no!"

	fmt.Println("variables and constants")
	fmt.Println(num1, i, j)      // 10, 3, 7
	fmt.Println(flag)            // false
	fmt.Println(retVal)          // result
	fmt.Println(c, python, java) // true false no!
}
