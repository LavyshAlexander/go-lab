package main

import (
	"fmt"
	"math/rand"
)

func main() {
	favorite_number := rand.Intn(100)
	fmt.Println("My favorite number is:", favorite_number)
}
