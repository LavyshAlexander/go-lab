package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	pr := fmt.Printf

	pr("struct1: %v\n", p)
	pr("struct2: %+v\n", p)
	pr("struct3: %#v\n", p)
}
