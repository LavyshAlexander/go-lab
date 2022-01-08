package main

import "fmt"

var a int

var (
	b bool
	f float32
	s string
)

func main() {
	a = 42
	b, f = true, 24.0
	s = "Viva la Go"
	dAndA := 23

	fmt.Println(a, b, f, s, dAndA)
}
