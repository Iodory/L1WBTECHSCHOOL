package main

import "fmt"

func main() {
	otvet := []int{}
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	checkerA := make(map[int]bool)
	for _, v := range A {
		checkerA[v] = true
	}

	checkerB := make(map[int]bool)
	for _, v := range B {
		if checkerA[v] && !checkerB[v] {
			otvet = append(otvet, v)
			checkerB[v] = true
		}
	}

	fmt.Println(otvet)
}
