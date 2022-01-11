package main

import "fmt"

func main() {
	const a int32 = 5
	const b rune = 7
	const f float32 = 7.777
	var c complex128 = 1 + 4i
	var d uint8 = 14
	var u string = "Ho Hey"

	n := 42
	pi := 3.14
	cx := 2 + 2i
	z := "Go is great"

	fmt.Printf("Types by user:\n %T, %T, %T, %T, %T, %T\n", a, b, f, c, d, u)
	fmt.Printf("Types by default:\n %T, %T, %T, %T\n", n, pi, cx, z)
}
