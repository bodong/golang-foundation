package main

import (
	"fmt"
	"math"
)

func main() {
	square := &Square{20}
	fmt.Printf("square area : %+v \n", square.area())

	circle := &Circle{10}
	fmt.Printf("circle area : %+v \n", circle.area())

	shapes := []Shape{square, circle}
	sa := sumAreas(shapes)

	fmt.Printf("area : %+v \n", sa)
}

//Square type
type Square struct {
	length float64
}

//area of square
func (square *Square) area() float64 {
	return square.length * square.length
}

//Circle type
type Circle struct {
	radius float64
}

//area of circle
func (circle *Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

//Shape interface
type Shape interface {
	area() float64
}

//sumAreas calculate areas
func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.area()
	}

	return total
}
