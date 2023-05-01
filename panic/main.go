package main

import (
	"fmt"
)

func main() {
	// panic: occurs when there is a system error, panic will make current thread crash
	// defer: transfer the control to the panic caller
	// recover: function recover from the panic or system errors
	// use `defer + recover` logic to fix panic

	// How to handle NPE in Go:
	// 1. adding nil check
	// 2. use defer & recover conbinations to catch errors

	// execution order:
	// 1. defer func is called
	// 2. a panic is triggered
	defer func() {
		fmt.Println("defer func is called")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("a panic is triggered")
}