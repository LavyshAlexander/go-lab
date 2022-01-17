package main

import (
	"fmt"
	"math/cmplx"
)

var (
	count  int
	ToBe   bool       = true
	MaxInt uint64     = 1<<64 - 1
	root   complex128 = cmplx.Sqrt(-4)
)

func main() {
	var c, python, java = true, true, "no!"
	l := 5

	fmt.Println(count, c, python, java, l)

	fmt.Printf("Type %T, value %v\n", ToBe, ToBe)
	fmt.Printf("Type %T, value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T, value %v", root, root)
}
