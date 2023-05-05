package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * Execution order:
 * lock: 0
 * rLock: 0
 * rLock: 1
 * rLock: 2
 * wLock: 0
 *
 * sync.RWMutex{} does not block read operation, it blocks write operation
 */
func main() {
	go rLock()
	go wLock()
	go lock()
	time.Sleep(5 * time.Second)
}

func lock() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("lock:", i)
	}
}

func rLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("rLock:", i)
	}
}

func wLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.Lock();
		defer lock.Unlock()
		fmt.Println("wLock:", i)
	}
}