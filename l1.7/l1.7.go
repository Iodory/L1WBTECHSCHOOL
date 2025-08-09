package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	mapa := make(map[string]int)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go Greeting(mapa, &mu, &wg)
	}
	wg.Wait()
	fmt.Println(mapa)
}

func Greeting(m map[string]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	m["hello"]++
	mu.Unlock()
}
