package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []string
	cond *sync.Cond
}

/**
 * Cond implements a condition variable, a rendezvous point for goroutines waiting
 * for or announcing the occurrence of an event.
 */
func main() {
	q := Queue{
		queue: []string{},
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// producer
	go func() {
		for {
			q.Enqueue("a")
			time.Sleep(2 * time.Second)
		}
	}()

	// receiver
	for {
		result := q.Dequeue()
		fmt.Println("retrieved data is: ", result)
		time.Sleep(time.Second)
	}
}

func (q *Queue) Enqueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	fmt.Printf("putting %s to queue, notify all\n", item)
	q.cond.Broadcast() // equals to java notifyAll()
}

func (q *Queue) Dequeue() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.queue) == 0 {
		fmt.Println("no data available, wait")
		q.cond.Wait()
	}
	// when code executes to this part, we assume we have data in queue
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}