package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << i
		fmt.Println(pow[i])
	}
}
