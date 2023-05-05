package main

import (
	// "fmt"
	"sync"
	"time"
)

func main() {
	// myMap := map[int]int{}
	// myMap[2] = 1
	// fmt.Println(myMap)

	// unsafeWrite() // fatal error: concurrent map writes
	safeWrite()
	time.Sleep(time.Second)
}

func unsafeWrite() {
	conflictMap := map[int]int{}
	for i := 0; i < 100; i++ {
		go func() {
			conflictMap[1] = i
		}()
	}
}

type SafeMap struct {
	safeMap map[int]int
	sync.Mutex
}

func (s *SafeMap) Read(k int) (int, bool) {
	s.Lock()
	defer s.Unlock()
	result, ok := s.safeMap[k]
	return result, ok
}

func (s *SafeMap) Write(k, v int) {
	s.Lock()
	defer s.Unlock()
	s.safeMap[k] = v
}

func safeWrite() {
	s := SafeMap{
		safeMap: map[int]int{},
		Mutex:   sync.Mutex{},
	}
	for i := 0; i < 100; i++ {
		go func() {
			s.Write(1, 1)
		}()
	}
}
