package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sidelength float64
}

type shape interface {
	area() float64
}

func main() {
	t := triangle{10, 5}
	s := square{5}

	printArea(t)
	printArea(s)
}

func (s square) area() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.area())
}
