package main

import (
	"context"
	"fmt"
	"time"
)

// mock context
type Context1 interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

func main() {
	// Context
	// 1. when we have timeout, cancel manipulation, or errors, we need to cancel subsequent manipulations with context
	// 2. Context is the structure body of setting deadline, sync signal, pass request parameter
	// usages:
	// - context.Background( major context, share values )
	// - context.TODO
	// - context.WithDeadline
	// - context.WithValue
	// - context.WithCancel

	// 1. Context share values
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		fmt.Println(c.Value("a")) // b
	}(ctx)

	time.Sleep(1 * time.Second) // we need to let main thread sleep to trigger child goroutine logic

	// 2. When timeout time reaches, we need to do sth.
	// execution order:
	// enter default
  // child process interrupt...
  // main process exit!
	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second) // with 1 second timeout
	defer cancel()

	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				// ...
				// when timeout time reaches, we need to do sth here.
				// ....
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)

	select {
	case <-timeoutCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}

	// time.Sleep(5 * time.Second) // print "child process interrupt..." only
}
