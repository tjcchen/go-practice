package main

import (
	"fmt"
)

func main() {
	// array - we need to specify length and type
	arr := [3]int{1, 2, 3}
	strs := [3]string{"January", "Feburary", "March"}

	fmt.Println(arr[0], arr[1], arr[2])    // 1, 2, 3
	fmt.Println(strs[1], strs[2], strs[0]) // Feburary March January

	// array and slice
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:4]                // namely 2, 3, 4 from myArray
	mySlice1 := []int{6, 5, 7}
	fmt.Printf("mySlice %+v\n", mySlice)   // mySlice [2 3 4]
	fmt.Printf("mySlice1 %+v\n", mySlice1) // mySlice1 [6 5 7]
	fullSlice := myArray[:]
	fmt.Println(fullSlice)                           // [1 2 3 4 5]
	remove3rdItem := deleteItem(fullSlice, 2)
	fmt.Printf("remove3rdItem %+v\n", remove3rdItem) // remove3rdItem [1 2 4 5]
}

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
