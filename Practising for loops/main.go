package main

import (
	"fmt"
)

func main() {
	x := 20
	for {
		fmt.Printf("x is %v\n", x)
		if x < 0 {
			break
		}
		x--
	}

}
