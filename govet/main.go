package main

import (
	"fmt"
)

func main() {
	name := "testing"
	fmt.Printf("%d\n", name)         // fmt.Printf format %d has arg name of wrong type string
	fmt.Printf("%s\n", name, name)   // fmt.Printf call needs 1 arg but has 2 args

	fmt.Printf("%s\n", name)         // correct
	fmt.Printf("%s%s\n", name, name) // correct
}