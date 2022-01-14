package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	favoriteNumber := rand.Intn(100)
	fmt.Println("My favorite number is:", favoriteNumber)
}
