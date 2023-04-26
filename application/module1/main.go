package main

import "fmt"

func main() {
	// declare a slice
	arr := []string{"I", "am", "stupid", "and", "weak"}
	
	printArr(arr) // I am stupid and weak

	fmt.Println(len(arr)) // 5

	// replace logic
	arr[2] = "smart"
	arr[4] = "strong"

	printArr(arr) // I am smart and strong
}

func printArr(arr []string) {
	if len(arr) == 0 {
		return;
	}

	for idx, value := range arr {
		fmt.Print(value)
		if idx + 1 == len(arr) {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
}