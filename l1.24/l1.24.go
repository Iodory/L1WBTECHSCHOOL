package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other *Point) float64 {
	step1 := (other.x-p.x)*(other.x-p.x) + (other.y-p.y)*(other.y-p.y)
	step2 := math.Sqrt(step1)
	return step2
}

func main() {
	point1 := NewPoint(5, 7)
	point2 := NewPoint(7, 5)

	fmt.Println(point1.Distance(point2))
}
