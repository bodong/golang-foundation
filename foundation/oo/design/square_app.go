package main

import (
	"fmt"
	"log"
)

func main() {
	square, err := NewSquare(-1, 3, 5)
	if err != nil {
		log.Fatalf("ERR : unable to create square")
	}
	fmt.Printf("Origin : %+v\n", square)
	square.Move(2, 2)
	fmt.Printf("After move : %+v\n", square)
	fmt.Println(square.Area())
}

//Point class
type Point struct {
	x int
	y int
}

//Move point
func (point *Point) Move(dx int, dy int) {
	point.x += dx
	point.y += dy
}

//Square class
type Square struct {
	center Point
	length int
}

//NewSquare will create new Square object
func NewSquare(x int, y int, length int) (*Square, error) {
	if x < 0 || y < 0 || length < 0 {
		return nil, fmt.Errorf("Given was x:%d, y:%d, length:%d. X or y or length must be > 0", x, y, length)
	}

	square := &Square{
		center: Point{x, y},
		length: length,
	}

	return square, nil
}

//Move square point to given coordinate param
func (square *Square) Move(dx int, dy int) {
	square.center.Move(dx, dy)
}

//Area of the square
func (square *Square) Area() int {
	return square.length * square.length
}
