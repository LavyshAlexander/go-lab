package main

import (
	"fmt"
	"math"
)

const epsilon float64 = 0.1

func squareRoot(x float64) (z float64) {
	z = x / 2
	iter := 0

	for math.Abs(z*z-x) > epsilon {
		fmt.Println(iter, z)
		z -= (z*z - x) / (2 * z)
		iter += 1
	}

	return
}

func main() {
	fmt.Println(squareRoot(200))
}
