package main

import (
	"fmt"

	"math"
)

func power(x float64) func() float64 {
	var y float64
	return func() float64 {
		y++
		return math.Pow(x, y)
	}
}

func main() {
	a := power(2)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}
