package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 0 {
		return arr
	}

	var left, right []int
	pivot := arr[len(arr)/2]

	for _, v := range arr {
		if v > pivot {
			left = append(left, v)
		} else if v < pivot {
			right = append(right, v)
		}
	}

	sortedR := quickSort(left)
	sortedL := quickSort(right)

	return append(append(sortedL, pivot), sortedR...)
}

func main() {
	arr := []int{4, 5, 6, 2, 3, 9, 1, 8, 7}
	sorted := quickSort(arr)

	fmt.Println(sorted)
}
