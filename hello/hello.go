package main

import (
	"example/user/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, world from breezing Go.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Printf(cmp.Diff("Hello World", "Hello Go"))
}
