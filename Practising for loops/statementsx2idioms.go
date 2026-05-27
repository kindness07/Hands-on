package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		x := rand.Intn(5)
		y := 3
		if y == x {
			fmt.Printf("youwin! %v is equal to %v\n", x, y)
		} else {
			fmt.Printf("%v is not equal to %v\n", x, y)
		}
	}

	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteration %v \t x is %v\n", i, x)
		}
	}
}
