package main

import (
	"fmt"
)

func main() {
	// we have to use a channel to guarantee execution order, otherwise
	// the "second line of code" won't be executed
	// execution order:
	// 1. first line of code
	// 2. second line of code
	// 3. third line of code 1
	fmt.Println("first line of code")
	ch := make(chan int)
	go func() {
		fmt.Println("second line of code")
		ch <- 1 // data assignment - send data
	}()
	// [important] this line of code blocks the main thread exit, since channel does not have value
	i := <-ch // retrieve data - receive data
	fmt.Println("third line of code", i)
}
