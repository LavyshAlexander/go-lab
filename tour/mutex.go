package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu     sync.Mutex
	values map[string]int
}

func (sc *SafeCounter) Inc(key string) {
	sc.mu.Lock()
	sc.values[key]++
	sc.mu.Unlock()
}

func (sc *SafeCounter) Value(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.values[key]
}

func main() {
	c := SafeCounter{values: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))

}
