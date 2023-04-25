package main

import (
	"fmt"
)

func main() {
	// 1. new and make
	// new returns a pointer
	// make return the first element, we can set memory space, to avoid future memory copying
	mySlice1 := new([]int)            // &[]
	mySlice2 := make([]int, 0)        // []
	mySlice3 := make([]int, 10)       // [0 0 0 0 0 0 0 0 0 0]
	mySlice4 := make([]int, 10, 20)   // [0 0 0 0 0 0 0 0 0 0]

	fmt.Println(mySlice1, mySlice2, mySlice3, mySlice4)

	// 2. error prone code - modify the value within slice
	// will not change slice value
	mySlice5 := []int{10, 20, 30, 40, 50}
	for _, value := range mySlice5 {
		value *= 2
	}
	fmt.Printf("mySlice5 %+v\n", mySlice5) // mySlice5 [10 20 30 40 50]

	// will change slice value
	for index, _ := range mySlice5 {
		mySlice5[index] *= 2
	}
	fmt.Printf("mySlice5 %+v\n", mySlice5) // mySlice5 [20 40 60 80 100]

	// 3. error prone code - try to avoid
	// Go change value only
	a := []int{}
	b := []int{1, 2, 3}
	c := a
	a = append(b, 1)     // after this manipulation, c is empty slice still
	fmt.Println(a, b, c) // [1 2 3 1] [1 2 3] []
}