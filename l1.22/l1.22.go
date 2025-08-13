package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(123456789010123134)
	b := big.NewInt(4178092340721024)

	sum := new(big.Int).Add(a, b)
	umn := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)
	min := new(big.Int).Sub(a, b)
	fmt.Printf("Сложение: %d\n Вычитание: %d \n Умножение: %d\n Деление: %d\n ", sum, min, umn, div)
}
