package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p2 = &Vertex{100, 23}
)

func main() {
	v := Vertex{1, 2}

	p := &v
	p.X = 1e9

	fmt.Println(v, v2, v3, p2)
}
