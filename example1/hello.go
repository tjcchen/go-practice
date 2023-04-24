package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	str := "hello world!"
	fmt.Println(str)
	// fmt.Printf format %d has arg str of wrong type string
	// fmt.Printf("%d\n", str)
}
