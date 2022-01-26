package main

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func Walk(t *Tree, ch chan<- int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for i := 0; i < 10; i++ {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if v1 != v2 {
			return false
		}

		if ok1 != ok2 {
			return false
		}
	}

	return true
}
