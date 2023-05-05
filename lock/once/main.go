package main

import (
	"fmt"
	"sync"
)

type SliceNum []int

func NewSlice() SliceNum {
	return make(SliceNum, 0)
}

func (s *SliceNum) Add(elem int) *SliceNum {
	*s = append(*s, elem)
	fmt.Println("add", elem)
	fmt.Println("add SliceNum end", *s)
	return s
}

func main() {
	var once sync.Once
	s := NewSlice()
	// use once, we execute Add() once only
	// check source code of once.Do()
	// print result:
	// - add 16
	// - add SliceNum end &[16]
	once.Do(func() {
		s.Add(16)
	})
	once.Do(func() {
		s.Add(16)
	})
	once.Do(func() {
		s.Add(16)
	})
}