package main

import (
	"fmt"
)

func main() {
	// 1. Thread lock
	// - thread safety( multi-thread code ): thread-safe code only manipulates shared data structures in a manner
	//   that ensures that all threads behave properly and fulfill their design specifications
	//   without unintended interaction.
	// - Golang not only support CSP(Communicating Sequential Process) model, the communcation between different channels,
	//   but also support multi-thread data sharing
	// - Sync syntax:
	//   sync.Mutex( mutual exclusion ) - mutual exclusion: Lock() add lock; UnLock() remove lock
	//   sync.RWMutex - read and write seperately: read does not block thread, but write blocks.
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
	// - check binary size: size basic5
	// - check text & bss: objdump -x basic5
	// - linux page size: getconf PAGE_SIZE( 4096 )
	// - GoLang implements user thread based on GMP( Goroutine, Machine, Process ) model

	// 3. Memory management
	// Heap memory management:
	// - Mutator( Ask for memory from Allocator ),
	// - Allocator( apply for a big memory block once, like 10G. used & unused ),
	// - Heap,
	// - Object Header( object header + metadata ),
	// - Collector( Garabage Collector, scan all the object headers )
	// ThreadCacheMalloc( Heap memory pre-allocation, each thread maintains its own private memory space to avoid multi-threads competition ):
	// - ThreadCache
	// - CentralCache
	// - PageHeap
	// - Virtual Memory
  // Golang memory allocation( two span class matches to one size class, one for pointer, the other for reference ):
	// - mcache
	// - mcentral
	// - mheap
	// - virtual memory
	// Memory collection:
	// - reference counter( Python, PHP, Swift ): each memory block maintains a counter, the an object is released, the counter--
	// - mark & sweep( Golang ): scan code from root, mark reference object, and collect unmarked object
	// - generation collector( Java ): Eden, Survivor, Tenured

	fmt.Println("---Go Routine---")

}