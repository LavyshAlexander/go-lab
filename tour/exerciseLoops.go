package main

import (
	"fmt"
	"math"
)

const epsilon float64 = 0.0001

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
	values := []float64{
		200,
		10087,
		1000046543,
	}

	for _, val := range values {
		sqrt := squareRoot(val)
		fmt.Printf("squareRoot(%v) = %v - math.Sqrt() = %v\n", val, sqrt, sqrt-math.Sqrt(val))
	}
}
