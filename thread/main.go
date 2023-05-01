package main

import (
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

}