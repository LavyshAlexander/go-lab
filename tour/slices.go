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
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
