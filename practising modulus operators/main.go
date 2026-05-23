package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 40; i++ {
		if i%2 != 1 {
			continue
		}
		fmt.Println(i)
	}
}
