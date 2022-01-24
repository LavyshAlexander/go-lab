package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

func ScaleFunc(v *Vertex, factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

type Abser interface {
	Abs() float64
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

	fmt.Println("-- Interfaces --")
	var a Abser

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())

	// a = v // cannot use v (type Vertex) as type Abser in assignment:
	// Vertex does not implement Abser (Abs method has pointer receiver)
	// fmt.Println(a.Abs())
}
