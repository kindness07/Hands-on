package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	x := sum(xi...)
	fmt.Println(x) // fmt.Println(sum(xi...)) --- IGNORE ---
	// i := [5]int{1, 2, 3, 4, 5}
	// y := bar(i)
	// fmt.Println(y)
	fmt.Println(bar(xi))

}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(i []int) int {
	fmt.Printf("%T\n", i)
	total := 0
	for _, v := range i {
		total += v
	}
	return total

}
