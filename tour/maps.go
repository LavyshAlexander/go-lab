package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

func main() {

	var m = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)

	fmt.Println("-- mutating maps --")
	m1 := make(map[string]int)

	m1["Answer"] = 42
	fmt.Println("The answer:", m1["Answer"])

	delete(m1, "Answer")
	v, ok := m1["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	fmt.Println("-- Exercise --")
	counts := WordCount("Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.You might find strings.Fields helpful.")
	fmt.Println(counts)
}

func WordCount(s string) (counts map[string]int) {
	counts = make(map[string]int)

	fields := strings.Fields(s)
	for _, f := range fields {
		_, ok := counts[f]
		if !ok {
			counts[f] = 0
		}

		counts[f] += 1
	}

	return
}
