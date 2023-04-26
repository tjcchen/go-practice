package main

import (
	"fmt"
	_ "github.com/tjcchen/go-practice/init/a"
	_ "github.com/tjcchen/go-practice/init/b"
)

/**
 * 1. init function will be executed when package is initiated
 * 2. init is executed before main function
 * 3. be cautious to use init() when we have nested imports
 */
func init() {
	fmt.Println("main init")
}

func main() {
	// 1. execution order
	// init from b
	// init from a
	// main init
	// main
	fmt.Println("main")
}