package main

import (
	"fmt"

	// "math.Pi"

	// "math.Pow"

	"math"
)

func main() {
	s1 := square{
		Length: 5,
		width:  5,
	}
	c1 := circle{
		radius: 3,
	}
	fmt.Println(info(s1))

	fmt.Println(info(c1))
}

type square struct {
	Length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.Length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2) // πr²
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()

}
