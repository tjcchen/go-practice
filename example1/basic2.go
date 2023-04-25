package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// variables
	var num1 int = 10
	var i, j int = 3, 7

	// constants
	const flag bool = false

	// declare and assignment
	retVal := "result"
	c, python, java := true, false, "no!"

	fmt.Println("variables and constants")
	fmt.Println(num1, i, j)      // 10, 3, 7
	fmt.Println(flag)            // false
	fmt.Println(retVal)          // result
	fmt.Println(c, python, java) // true false no!

	// type conversion
	var num2 int = 42;
	var num3 float64 = float64(num2);
	var num4 uint = uint(num3);

	// or
	num5 := 43;
	num6 := float64(num5)
	num7 := uint(num6)

	fmt.Println(num2, num3, num4, num5, num6, num7); // 42 42 42 43 43 43

	// print vars block
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)     // Type: bool Value: false
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt) // Type: uint64 Value: 18446744073709551615
	fmt.Printf("Type: %T Value: %v\n", z, z)           // Type: complex128 Value: (2+3i)

	// type inference
	var x int
	y := x                              // y is an integer too
	fmt.Printf("Type: %T, %T\n", x, y); // Type: int, int
}
