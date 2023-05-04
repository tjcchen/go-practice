package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan int, 10)

	// // consumer
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ch:
	// 			fmt.Println("retrieve value from channel", <-ch)
	// 		}
	// 	}
	// }()

	// // producer
	// for i := 0; i < 5; i++ {
	// 	time.Sleep(1 * time.Second)
	// 	ch <- i
	// 	fmt.Println("put value to channel", i)
	// }
	
	// producer1
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("put value to channel", i)
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	// consumer1
	// solution1
	for v := range ch {
		fmt.Println("retrieve value from channel", v);
	}
}
