package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	t := triangle{base: 5, height: 10}
	r := square{sideLength: 6}

	printArea(t)
	printArea(r)
}
