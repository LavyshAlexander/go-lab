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

	fmt.Println("-- length and capacity --")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	underlyingArray := *(*[4]int)(s)
	fmt.Printf("type = %T of %v\n", underlyingArray, underlyingArray)

	fmt.Println("-- zero value --")
	var sNil []int
	fmt.Println(sNil, len(sNil), cap(sNil))
	if sNil == nil {
		fmt.Println("It is nil!")
	}

	fmt.Println("-- Creating a slice with make --")
	a5 := make([]int, 5)
	printSlice(a5)

	b05 := make([]int, 0, 5)
	printSlice(b05)

	cb2 := b05[:2]
	printSlice(cb2)

	dc25 := cb2[2:5]
	printSlice(dc25)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
