package main

import (
	"time"
	"fmt"
)

// mock context
type Context interface {
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
}