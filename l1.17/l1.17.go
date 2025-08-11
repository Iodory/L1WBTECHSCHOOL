package main

import "fmt"

func binarySearch(arr []int, num int) int {
	min := 0
	max := len(arr) - 1

	for min <= max {
		mid := (min + max) / 2
		arrMid := arr[mid]

		if arrMid < num {
			min = mid + 1
		} else if arrMid > num {
			max = mid - 1
		} else {
			return mid
		}

	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num := 7

	find := binarySearch(arr, num)

	fmt.Printf("Число под индексом: %d", find)
}
