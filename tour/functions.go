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

	fmt.Println("-- variadic functions --")
	sum(1, 2)
	sum(1, 2, 3)
	sum([]int{1, 2, 3, 4, 5, 6}...)

	fmt.Println("-- recursive closure --")
	// Will through UndeclaredName compile error if it will be undeclared.
	var fib func(int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
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
