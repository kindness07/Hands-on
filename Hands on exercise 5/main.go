package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("x is %v\t", x)
		switch x {
		case 0:
			fmt.Printf("iteration %v\t x is %v\n", i, x)
		case 1:
			fmt.Printf("iteration %v\t x is %v\n", i, x)
		case 2:
			fmt.Printf("iteration %v\t x is %v\n", i, x)
		case 3:
			fmt.Printf("iteration %v\t x is %v\n", i, x)
		case 4:
			fmt.Printf("iteration %v\t x is %v\n", i, x)
		default:
			fmt.Printf("iteration %v\t x is greater than %v\n", i, x)
		}
	}
}
