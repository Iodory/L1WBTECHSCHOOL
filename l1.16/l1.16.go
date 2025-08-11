package main

import "fmt"

func quickSort(arr []int) []int {

	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{4, 5, 6, 2, 3, 9, 1, 8, 7}
	sortArr := quickSort(arr)
	fmt.Println(sortArr)
}
