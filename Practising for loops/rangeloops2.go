package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		// "James":      48,
		"Moneypenny": 32,
		"Q":          28,
	}
	for k, v := range m {
		fmt.Printf("%v is %v years old\n", k, v)
	}
	age := m["James"]
	fmt.Println(age)

	age = m["Q"]
	fmt.Println(age)

	if James, ok := m["James"]; !ok {
		fmt.Printf("James is not in the map\n", James)
	}
}
