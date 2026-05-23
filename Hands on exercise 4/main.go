package main

import "fmt"

func main() {
	y := 0
	for y < 100 {
		fmt.Printf("%v\t is of type %T\n", y, y)
		y++
	}
}
