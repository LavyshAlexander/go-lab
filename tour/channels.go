package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	fmt.Println("-- Buffered channel --")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)

	fmt.Println("-- Range and Close --")

	cFib := make(chan int, 10)
	go fibonacci(cap(cFib), cFib)
	for i := range cFib {
		fmt.Println(i)
	}

	fmt.Println("-- select --")
	cSel := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-cSel)
			// if i == 7 {
			// 	quit <- 0
			// }
		}
		quit <- 0 // if comment -> obviously deadlock
	}()

	fibonacciSelect(cSel, quit)
}

func fibonacci(n int, cFib chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		cFib <- x
		x, y = y, x+y
	}
	close(cFib) // if not close will be deadlock
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
