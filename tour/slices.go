package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[:2]
	b := names[1:]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// length and capacity
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
}

func printSlice(s []int) {
	panic("unimplemented")
}
