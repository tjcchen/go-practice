package main

import (
	"fmt"
	"time"
)

func main() {
	// timer
	// 1. time.Ticker can specify a time interval to send time value to channel C
	// 2. use case: set timeout for goroutine
	ch := make(chan int)
	timer := time.NewTimer(time.Second)
	select {
	// check normal channel
	case <-ch:
		fmt.Println("received from ch")
	case <-timer.C:
		fmt.Println("timeout waiting from channel ch")
	}
}
