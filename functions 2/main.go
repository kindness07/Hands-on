package main

import (
	"fmt"
)

func square(x int) int {
	return x * x
}

func printSquare(f func(int) int, x int) {
	fmt.Printf("The square of %d is %d\n", x, f(x))
}

func main() {
	printSquare(square, 5)
}
