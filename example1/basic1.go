package main

import "fmt"

func main() {
	fmt.Println("---basic syntax---")

	// 1. if
	var num = 1
	if num > 1 {
		fmt.Println("greater than 1")
	} else if num < 1 {
		fmt.Println("less than 1")
	} else {
		fmt.Println("equals to 1")
	}

	// if multiple conditions
	// In Go, := is for declaration + assignment, whereas = is for assignment only.
	// For example, var foo int = 10 is the same as foo := 10.
	if v := num - 100; v < 0 {
		fmt.Println("v is less than 0")
	}

	// 2. switch
	counter := 100
	switch counter {
	case 99:
		fmt.Println("less than 100")
	case 100:
		fmt.Println("equals to 100")
		fallthrough // fall through to next switch case
	case 101:
		selfPrint()
	default:
		fmt.Println("fall back to this line")
	}

	// 3. for - Go only has for loop
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum) // 45 - 0 ~ 9

	// Go does not support while keyword, we can use the following case if needed
	for sum < 100 {
		sum += sum
	}
	println(sum) // 180

	// infinite loop
	for {
		sum++
		println("current sum", sum)
		if sum == 190 {
			println("sum equals to 190")
			break
		}
	}

	// 4. for-range
	// iterate over array, string, map, object etc
	for index, char := range []string{"first", "second", "third"} {
		println(index, char)
	}
}

func selfPrint() {
	fmt.Println("---greater than 100---")
}
