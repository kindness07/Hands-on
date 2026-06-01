package main

import (
	"fmt"
)

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "James Bond"
}

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}
