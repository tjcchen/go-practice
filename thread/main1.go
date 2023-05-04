package main

import (
	"fmt"
)

func main() {
	// how to stop a child goroutine? with close(channel)
	done := make(chan bool)
	go func() {
		fmt.Println("go routine code")
		done <- true
		for {
			fmt.Println("infinite loop")
			select {
			case <-done:
				fmt.Println("done channel is triggered, exit child go routine")
				return
			default:
				fmt.Println("select default branch")
				return
			}
		}
	}()
	val := <-done
	fmt.Println("main thread code", val)
	close(done)
}
