package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(10, 32))

	resFromClosure := func(a, b int) int { return a * b }(10, 23)
	fmt.Println(resFromClosure)
}
