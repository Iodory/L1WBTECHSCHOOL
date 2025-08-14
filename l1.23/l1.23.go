package main

import "fmt"

func remove(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Неверное значение индекса")
		return slice
	}

	copy(slice[i:], slice[i+1:])

	slice[len(slice)-1] = 0

	return slice[:len(slice)-1]
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newSlcie := remove(slice, 6)
	fmt.Printf("Result: %v", newSlcie)
}
