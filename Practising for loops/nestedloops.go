package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("\nthe outer loop runs 1 time\n\n")
		for j := 0; j < 5; j++ {
			fmt.Printf("the inner loop runs %v times\n", j)
		}
	}
}
