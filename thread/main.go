package main

import (
	"time"
	"fmt"
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