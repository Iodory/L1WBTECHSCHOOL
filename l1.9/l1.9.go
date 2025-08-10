package main

import (
	"fmt"
	"sync"
)

func main() {
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int, len(myArray))
	chDouble := make(chan int, len(myArray))

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for _, v := range myArray {
			ch <- v
		}
		defer close(ch)
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			chDouble <- v * 2
		}
		defer close(chDouble)
	}()

	go func() {
		defer wg.Done()
		for v := range chDouble {
			fmt.Println(v)
		}
	}()

	wg.Wait()

}
