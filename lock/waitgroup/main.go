package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	// waitBySleep()
	// waitByChannel()
	waitByWG()
}

func waitBySleep() {
	for i := 0; i < 100; i++ {
		go fmt.Println(i) // the result order cannot be guaranteed with goroutine
	}
	time.Sleep(time.Second)
}

func waitByChannel() {
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true     // putting
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-c             // retrieving
	}
}

func waitByWG() {
	wg := sync.WaitGroup{}
	// We have 100 tasks in WaitGroup, when each task is Done, we invoke wg.Done()
	// After the 100 tasks is completed, we release the blocker
	wg.Add(100) 
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}