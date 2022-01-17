package main

import (
	"fmt"
	"math/cmplx"
)

const Pi = 3.14

var (
	count  int
	ToBe   bool       = true
	MaxInt uint64     = 1<<64 - 1
	root   complex128 = cmplx.Sqrt(-4)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	var c, python, java = true, true, "no!"
	l := 5

	fmt.Println(count, c, python, java, l)

	fmt.Printf("Type %T, value %v\n", ToBe, ToBe)
	fmt.Printf("Type %T, value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T, value %v\n", root, root)

	var s string
	fmt.Printf("v - %v, q - %q\n", s, s)

	fmt.Println("happy", Pi, "day")

	// Numeric Constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big)) // overflow
	// fmt.Printf("%T", Big) // even here will throw compile error, cause Big is untyped int constant.
}
