package main

import "fmt"

type stack struct {
	index int
	data  [5]int
}

func (s *stack) push(k int) {
	s.data[s.index] = k
	s.index++
}

func (s *stack) pop() int {
	s.index--
	return s.data[s.index]
}

func main() {
	s := new(stack)
	s.push(42)
	s.push(23)
	fmt.Printf("stack: %v\n", s)
}
