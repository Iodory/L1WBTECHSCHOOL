package main

import "fmt"

func main() {
	A := []string{"cat", "cat", "dog", "cat", "tree"}
	ynical := []string{}
	checker := make(map[string]bool)

	for _, v := range A {
		if !checker[v] {
			ynical = append(ynical, v)
			checker[v] = true
		}
	}

	fmt.Println(ynical)
}
