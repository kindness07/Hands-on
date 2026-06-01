package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world 4")
	defer fmt.Println("world 3")
	defer fmt.Println("world 2")
	defer fmt.Println("world 1")
	fmt.Println("hello") //last in firtst out
}
