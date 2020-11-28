package main

import (
	"fmt"
)

func main() {
	p := Point{2, 3}
	fmt.Printf("Original : %+v \n", p)

	p.Move(2, 2)
	fmt.Printf("after move : %+v \n", p)

}

//Point data
type Point struct {
	x int
	y int
}

//Move the point
func (p *Point) Move(dx int, dy int) {
	p.x += dx
	p.y += dy
}
