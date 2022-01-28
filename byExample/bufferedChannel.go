package main

import "fmt"

func main() {
	// Buffered channels accept a limited number of values
	// without a corresponding receiver for those values.
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
