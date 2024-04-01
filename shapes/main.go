package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLenght float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		base:   10,
		height: 5,
	}

	s := square{
		sideLenght: 20,
	}

	printArea(t)
	printArea(s)

}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
