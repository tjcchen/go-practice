package main

import (
	"fmt"
	"time"
)

/**
 * execution order:
 *   send message: 0
 *   send message: 1
 *   send message: 2
 *   send message: 3
 *   send message: 4
 *   child process interrupt...
 *   main process exit!
 */
func main() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages) // close messages channel eventually

	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C { // infinite loop
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	time.Sleep(5 * time.Second) // main thread time.Sleep will trigger its goroutine
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
