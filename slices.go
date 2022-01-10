package main

import (
	"fmt"
)

func main() {
	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust",
	}

	/* Define slices */
	classics := languages[0:3] // alt languages[:3]
	modern := make([]string, 4)
	modern = languages[3:7] // include 3 exclude 7
	new := languages[7:]

	fmt.Printf("classic languages: %v\n", classics)
	fmt.Printf("modern languages: %v\n", modern)
	fmt.Printf("new languages: %v\n", new)
}
