package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x is %v and y is %v\n", x, y)

	/*if x < 4 && y < 4 {
		fmt.Println("x and y are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are greater than 6")
	} else if x >= 4 && x <= 6 || x == 6 {
		fmt.Println("x is between 4 and 6 inclusive")
	} else if y != 6 {
		fmt.Println("y is not equal to 6")
	} else {
		fmt.Println("none of the conditiosns were met")
	}*/

	switch {
	case x < 4 && y < 4:

		fmt.Println("x and y are less than 4")

	case x > 6 && y > 6:

		fmt.Println("x and y are greater than 6")

	case x >= 4 && x <= 6 || x == 6:

		fmt.Println("x is between 4 and 6 inclusive")

	case y != 6:

		fmt.Println("y is not equal to 6")

	default:

		fmt.Println("none of the conditiosns were met")

	}
}
