package main

import "fmt"

func main() {
	a := 10
	b := 5
	a, b = b, a
	fmt.Printf("a:%d \nb:%d ", a, b)
}
