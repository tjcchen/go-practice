package main

func main() {
	var a *int
	*a += 1

	// doOperation(1, increase)
	doOperation(1, decrease)
}

func doOperation(y int, fn func(int, int)) {
	fn(y, 1)
}

func increase(a, b int) int {
	return a + b
}

func decrease(a, b int) {
	println("decrease result is: ", a-b)
}
