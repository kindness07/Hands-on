package main

import (
	"fmt"
)

func main() {
	c := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(c)

	c = append(c[:2], c[6:]...)
	fmt.Println(c)

	c = append(c, 52)

	fmt.Println(c)
	c = append(c, 53, 54, 55)

	fmt.Println(c)
	c = append(c, 56, 57, 58, 59, 60)

	fmt.Println(c)

}
