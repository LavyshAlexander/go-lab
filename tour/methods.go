package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func (v *Vertex) Scale(factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

func ScaleFunc(v *Vertex, factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	fmt.Println("-- Pointer receiver --")
	v.Scale(10)
	fmt.Println(v.Abs())

	p := &Vertex{3, 4}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(p, p.Abs())
}
