package main

import (
	"time"
	"fmt"
)

/**
 * Execution order:
 * 1. go routine code
 * 2. infinite loop
 * 3. done channel is triggered, exit child go routine
 * 4. main thread code
 */
func main() {
	// how to stop a child goroutine? with close(channel)
	done := make(chan bool)
	go func() { // go routine listens to channel write
		fmt.Println("go routine code")
		for {
			fmt.Println("infinite loop")
			select {
			case <-done:
				fmt.Println("done channel is triggered, exit child go routine")
				return
			default:
				fmt.Println("select default branch")
			}
		}
	}()

	done <- true // trigger go routine
	fmt.Println("main thread code")
	time.Sleep(3 * time.Second)
	close(done) // important

	// close child goroutine with Context - please refer to donechannel.go
}
