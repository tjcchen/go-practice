package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	// concurrency & parallism
	// 1. concurrency: two or more events occured with intervals at the same time - serial( one CPU )
	// 2. parallism: two or more events occured at the same time - parallel( multiple CPUs )

	// process & thread & coroutine
	// 1. process: an application running in your computer, your computer will allocate memory, and CPU to it
	// 1.1. ps -ef to check system processes
	// 2. thread: a thread is execution flow. a process could have multiple threads ( by using multiple CPUs )
	// 2.1 multiple threads shared the same process memory
	// 3. coroutine: coroutine is a lightweight thread implementation in Go - goroutine

	// CSP( Communicating Sequential Process ) - multi-threads model in Go
	// 1. CSP - communicating model between entities
	// 2. channel - communicating between different threads

	// thread & goroutine
	// 1. goroutine is more lightweight than threads
	// 2. the switching process between different threads is expensive, while goroutine cost less
	// 3. GOMAXPROCS - control the max serial thread number( concurrency )

	// Examples
	loopFunc()
	time.Sleep(time.Second) // loopFunc:  0, 1, 2

	loopFuncWithGoroutine()
	time.Sleep(time.Second) // loopFunc:  0, 2, 1 - order does not fixed

	fmt.Println("-------------------------------------------------------")

	// channel - communication between different threads
	// 1. one side sends data, the other side receives data
	// 2. there is only one goroutine(thread) can access data at a time, there is no resources competition in GO
	// 3. goroutine execute in code order
	// execution sequence:
	// 1. hello from goroutine
	// 2. 0
	ch := make(chan int)
	go func() {
		fmt.Println("hello from goroutine")
		ch <- 0      // write data to channel
	}()
	i := <-ch      // retrieve data from channel and assign to another variable
	fmt.Println(i) // 0

	// channel buffer
	// 1. the communication of channel is in a sync way
	// 2. when the buffer is full, the sending of the data is blocked
	// 3. we can use `make` to define a channel buffer, the default buffer is 0
	// ch1 := make(chan int)        // channel buffer is 0
	// ch2 := make(chan int, 1)     // channel buffer is 1

	// iterate over channel buffer zone
	ch3 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)  // n will be between 0 and 10
			fmt.Println("putting: ", n)
			ch3 <- n
		}
		close(ch3)
	}()

	fmt.Println("hello from main")
	for v := range ch3 {
		fmt.Println("receiving: ", v)
	}
}

func loopFunc() {
	for i := 0; i < 3; i++ {
		fmt.Println("loopFunc: ", i)
	}
}

// with goroutine - multiple CPUs
func loopFuncWithGoroutine() {
	for i := 0; i < 3; i++ {
		go fmt.Println("loopFunc: ", i)
	}
}