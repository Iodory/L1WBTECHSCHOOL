package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	exampleArray := []int{2, 4, 6, 8, 10}

	for i := 0; i < len(exampleArray); i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(exampleArray[i])
	}

	wg.Wait()
}
