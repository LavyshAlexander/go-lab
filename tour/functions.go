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

	fmt.Println("-- Closures --")
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fmt.Println("-- Exercise --")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	current, next := 0, 1

	return func() int {
		defer func() { current, next = next, (current + next) }()
		return current
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
