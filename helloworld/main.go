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

	// variadic parameters
	// Eg: func append(slice []Type, elems ...Type) []Type
	arr := append([]int{1, 2}, 3, 4, 5)
	fmt.Println(arr) // [1 2 3 4 5]

	myArr := []string{}
	myArr = append(myArr, "a", "b", "c")
	fmt.Println(myArr) // [a b c]

	/**
	 * Some useful built-in functions:
	 * 1. close - close pipeline
	 * 2. len, cap - returns the length and capacity of array, slice, or map
	 * 3. new, make - memory allocation
	 * 4. copy, append - manipulate with slice
	 * 5. panic, recover - error catching mechanism
	 * 6. print, println - print messages
	 * 7. complex, real, imag
	 */
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

/**
 * give return value a name, in this case, err & result
 */
func DuplicateString1(input string) (err error, result string) {
	if input == "aaa" {
		err = fmt.Errorf("aaa is not allowed")
		return;
	}
	result = input + input
	return;
}
