package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (l byLength) Len() int {
	return len(l)
}

func (l byLength) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byLength) Less(i, j int) bool {
	return len(l[i]) < len(l[j])
}

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{5, 2, 7, 9, 43, 3, 6, 78}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)

	fmt.Println("-- custom sort function --")
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
