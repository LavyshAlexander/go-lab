package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

const epsilon float64 = 0.0001

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(x / 2)
	iter := 0

	for math.Abs(z*z-x) > epsilon {
		fmt.Println(iter, z)
		z -= (z*z - x) / (2 * z)
		iter += 1
	}

	return z, nil
}

func main() {
	fmt.Println(squareRoot(2))
	fmt.Println(squareRoot(-2))
}
