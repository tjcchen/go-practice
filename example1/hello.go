package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	str := "hello world!"
	fmt.Println(str)

	// go vet xxx.go report includes:
	// 1. type does not match
	// fmt.Printf format %d has arg str of wrong type string
	// fmt.Printf("%d\n", str)

	// 2. redundant expressions
	// i := 0
	// fmt.Println(i != 0 || i != 1);

	// 3. go routine cannot be executed
	// words := []string{"foo", "bar", "baz"}
	// for _, word := range words {
	//  // go routine - a goroutine is a lightweight thread of execution.
	// 	go func() {
	// 		fmt.Println(word)
	// 	}()
	// 	// fmt.Println(word)
	// }

	// 4. unreachable code, like codes after return expression

	// 5. error checking too late
	// res, err := http.Get("https://www.spreadsheetdb.io/")
	// defer res.Body.close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
