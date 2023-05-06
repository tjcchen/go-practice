package main

import (
	"fmt"
)

func main() {
	// 1. Thread lock
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

	// 2. Thread scheduling - threads are scheduled for execution based on their priority.
	// - process: A process in Linux is nothing but a program in execution. It's a running instance of a program.
	//            It is the smallest unit of resource allocation. A process has it's own memory, fs, files, signal etc
	// - thread: Thread is the set of instructions executed within a process that can range from a single thread to multiple.
	//           It is the smalllest unit for scheduling
  // - No matter it is a process or a thread, the linux would use `task_struct` to describe it. From a linux kernel perspective,
	//   Thread is no difference to process
	// - Glibc's pthread library offers NPTL( Native POSIX Threading Library ) support

	fmt.Println("---thread lock---")

}