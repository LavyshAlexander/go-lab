package main

import "fmt"

func cookingGopher(id int, channel chan int) {
	fmt.Println("gopher", id, "started cooking")
	channel <- id
}

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c)
	}

	for i := 0; i < 5; i++ {
		gId := <-c
		fmt.Println("gopher", gId, "finished the dish")
	}
}
