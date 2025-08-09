package main

import "fmt"

func setBit(n int64, i uint, bitValue uint) int64 {
	mask := int64(1) << i

	if bitValue == 1 {
		return n | mask
	} else {
		return n &^ mask
	}
}

func main() {
	var n int64 = 5
	var i uint = 0
	var bitValue uint = 0

	fmt.Printf("Исходное число: %d\n", n)

	result := setBit(n, i, bitValue)

	fmt.Printf("Результат после установки %d-го бита в %d: %d\n", i, bitValue, result)
}
