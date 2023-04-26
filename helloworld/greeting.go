package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args   // Passing multiple parameters. Eg: ./greeting param1 param2
	fmt.Println(args)
	if len(args) != 0 {
		fmt.Println("Does not accept any argument")
		os.Exit(1)
	}
	fmt.Println("hello world")
}