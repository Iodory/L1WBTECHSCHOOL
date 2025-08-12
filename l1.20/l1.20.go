package main

import (
	"fmt"
	"strings"
)

func main() {
	words := "snow dog sun"
	oneWord := strings.Fields(words)

	for i, j := 0, len(oneWord)-1; i < j; i, j = i+1, j-1 {
		oneWord[i], oneWord[j] = oneWord[j], oneWord[i]
	}
	words = strings.Join(oneWord, " ")

	fmt.Println(words)
}
