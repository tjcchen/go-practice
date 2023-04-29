package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	// defer - FILO
	// defer equals to finally in Java, C#, the relevant logic will be executed finally
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	loopFunc()              // loopFunc 2, 1, 0 - order does not fixed
	time.Sleep(time.Second) // print sequence: 3, 2, 1
}

func loopFunc() {
	// Mutual exclusion, mutex
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		// use a closure to resolve the dead lock issue
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("loopFunc", i)
		}(i)
	}
}