package main

import (
	"fmt"
)

func main() {
	// var a *int
	// *a += 1

	DoOperation(1, increase) // increase result is:  2
	// DoOperation(1, decrease) // decrease result is: 0
}

func DoOperation(y int, fn func(int, int)) {
	fn(y, 1)
}

func increase1(a, b int) int {
	return a + b
}

func increase(a, b int) {
	fmt.Println("increase result is: ", a+b)
}

func decrease(a, b int) {
	fmt.Println("decrease result is: ", a-b)
}
