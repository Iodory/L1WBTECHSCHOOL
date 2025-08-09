package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		Writer(ctx, ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Reader(ch)
	}()

	wg.Wait()
}

func Reader(ch <-chan int) {
	for v := range ch {
		fmt.Printf("Прочитал: %d\n", v)
	}
}

func Writer(ctx context.Context, ch chan<- int) {
	defer close(ch)

	var counter int
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- counter:
			counter++
		}
	}
}
