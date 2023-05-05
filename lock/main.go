package main

import (
	"fmt"
)

func main() {
	// 1. thread lock
	// - thread safety( multi-thread code ): thread-safe code only manipulates shared data structures in a manner
	//   that ensures that all threads behave properly and fulfill their design specifications
	//   without unintended interaction.
	// - Go not only support CSP(Communicating Sequential Process) model, the communcation between different channels,
	//   but also support multi-thread data sharing
	// - Sync syntax:
	//   sync.Mutex( mutual exclusion ) - mutual exclusion: Lock() add lock; UnLock() remove lock
	//   sync.RWMutex - read and write seperately
	//   sync.WaitGroup - wait for a group of goroutine returns
	//   sync.Once - code being executed once only
	//   sync.Cond - make a group of goroutine being revoked under certain conditions
	//   
	fmt.Println("---thread lock---")


}