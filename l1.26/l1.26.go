package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "acd"

	if Ynucal(str) == true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

func Ynucal(str string) bool {
	stRune := []rune(str)
	checker := make(map[rune]bool)

	for _, v := range stRune {
		v = unicode.ToLower(v)

		if checker[v] {
			return false
		}

		checker[v] = true
	}

	return true
}
