package main

import (
	"fmt"
)

func main() {
	a := Add()

	fmt.Println(a(5, 5))
}

func Add() func(x, y int) int {
	return func(x, y int) int {
		return x + y
	}
}
