package main

import (
	"fmt"
	"reflect"
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
	fmt.Println("|:-------------------------:|")

	allLangs := languages[:]
	fmt.Println(reflect.TypeOf(allLangs).Kind())

	frameworks := []string{
		"React", "Vue", "Angular", "Svetle",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrameworks := frameworks[0:4:4]
	frameworks = append(frameworks, "Meteor")

	fmt.Printf("all frameworks: %v\n", frameworks)
	fmt.Printf("js frameworks: %v\n", jsFrameworks)
}
