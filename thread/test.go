package main

import (
	"fmt"
)

func main() {
	// we have to use a channel to guarantee execution order, otherwise
	// the "second line of code" won't be executed
	fmt.Println("first line of code")
	ch := make(chan int)
	go func() {
		fmt.Println("second line of code")
		ch <- 1 // data assignment - send data
	}()
	i := <-ch // retrieve data - receive data
	fmt.Println("third line of code", i)
}
