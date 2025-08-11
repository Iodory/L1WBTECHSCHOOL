package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
	mu      sync.Mutex
}

func (c *Counter) plusPlus() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}
func main() {
	var cntr Counter
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cntr.plusPlus()
		}()
	}

	wg.Wait()
	fmt.Println(cntr.counter)
}
