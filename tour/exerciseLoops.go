package main

import (
	"fmt"
	"math"
)

const epsilon float64 = 0.1

func squareRoot(x float64) (z float64) {
	z = 1
	for math.Abs(z*z-x) > epsilon {
		z -= (z*z - x) / (2 * z)
	}

	return
}

func main() {
	fmt.Println(squareRoot(2))
}
