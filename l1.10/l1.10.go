package main

import (
	"fmt"
	"math"
)

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempGroup := make(map[int][]float64)

	for _, v := range temp {
		key := int(math.Floor(v/10) * 10)
		tempGroup[key] = append(tempGroup[key], v)
	}

	for k, v := range tempGroup {
		fmt.Printf("%d, %v\n", k, v)
	}
}
