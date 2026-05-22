package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	x := rand.Intn(250)
	fmt.Printf("%v\t is of type %T\n", x, x)

	/*if x <= 100 {
		fmt.Println("x is less than or equal to 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("x is between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("x is between 201 and 250")
	} */

	switch {
	case x <= 100:
		fmt.Println("x is less than or equal to 100")
	case x >= 101 && x <= 200:
		fmt.Println("x is between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("x is between 201 and 250")
	default:
		fmt.Println("x is greater than 250")
	}
}
