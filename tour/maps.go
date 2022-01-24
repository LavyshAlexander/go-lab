package main

import "fmt"

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
}
