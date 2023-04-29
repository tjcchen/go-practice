package main

import (
	"fmt"
	"errors"
)

type error interface {
	Error() string
}

type StatusError struct {
	ErrStatus string
}

// Error implements the error interface
func (e *StatusError) Error() string {
	return e.ErrStatus;
}

func main() {
	// We not have exception or try-catch grammar support in Go,
	// we deal with error with error messages. To create an error, we use:
	// 1. errors.New()
	// 2. fmt.errorF()
	var errNotFound error = errors.New("NotFound")
	fmt.Println(errNotFound) // NotFound

	var err error = &StatusError{ErrStatus: "Internal Error"}
	fmt.Println(err.Error()) // Internal Error
}