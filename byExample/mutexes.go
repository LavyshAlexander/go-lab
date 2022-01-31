package main

import (
	"fmt"
	"sync"
)

type CounterContainer struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *CounterContainer) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := CounterContainer{
		counters: make(map[string]int),
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}

		wg.Done()
	}

	wg.Add(3)
	doIncrement("a", 1000)
	doIncrement("a", 1000)
	doIncrement("b", 1000)

	wg.Wait()
	fmt.Println(c.counters)
}
