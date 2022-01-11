package main

import "fmt"

func average(x []float64) (avg float64) {
	total := 0.0

	if len(x) == 0 {
		avg = 0
	} else {
		for _, v := range x {
			total += v
		}

		avg = total / float64(len(x))
	}

	return
}

func print_reversed(input int) {
	if input <= 0 {
		return
	}

	for input > 0 {
		fmt.Printf("%d\n", input)
		input -= 1
	}
}

func main() {
	x := []float64{5, 3.6, 7.4, .3, 34.23}
	avg := average(x)

	fmt.Println(avg)

	print_reversed(10)
}
