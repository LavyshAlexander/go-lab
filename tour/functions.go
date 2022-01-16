package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func mod(a, b int) (mod, div int) {
	mod = a % b
	div = a / b
	return
}

func main() {
	fmt.Println(add(10, 32))

	resFromClosure := func(a, b int) int { return a * b }(10, 23)
	fmt.Println(resFromClosure)

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	mod, div := mod(10, 3)
	fmt.Println(mod, div)
}
