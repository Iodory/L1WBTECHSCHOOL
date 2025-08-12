package main

import "fmt"

func main() {
	str := "Привет"
	runeStr := []rune(str)
	var reversRune []rune

	for i := len(runeStr) - 1; i >= 0; i-- {
		reversRune = append(reversRune, runeStr[i])
	}

	newStr := string(reversRune)

	fmt.Println(newStr)
}
