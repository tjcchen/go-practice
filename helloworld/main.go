package main

import (
	"flag"
	"fmt"
	"os"
)

/**
 * Eg: ./main --name jack
 */
func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("os args is", os.Args) // os args is [./helloworld/main]
	// The * character is used to define a pointer in both C and Go.
	// Instead of a real value the variable has an address to the location of a value.
	// This name has to be a pointer
	fmt.Println("input parameter is:", *name)              // input parameter is: world
	fullString := fmt.Sprintf("Hello %s from Go\n", *name) // Hello world from Go
	fmt.Println(fullString)

	// duplicate string
	err, result := DuplicateString("abb")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

/**
 * A normal function with multiple return values
 */
func DuplicateString(input string) (error, string) {
	if input == "aaa" {
		return fmt.Errorf("aaa is not allowed"), ""
	}
	return nil, input + input
}
