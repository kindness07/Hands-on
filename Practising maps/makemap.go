package main

import (
	"fmt"
)

func main() {
	ms := make(map[string][]string)
	ms["bond_james"] = []string{"Shaken, not stirred", "Martinis", "fast cars"}
	ms["moneypenny_miss"] = []string{"intelligence", "Literature", "Computer Science"}
	ms["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	ms[`fleming_ian`] = []string{"steaks", "cigars", "espionage"}

	for k, v := range ms {
		fmt.Printf("k: %s, v: %v\n", k, v)
		for i, val := range v {
			fmt.Printf("\tindex position: %d \t value: %s\n", i, val)
		}
	}

	fmt.Println("********************")

	delete(ms, "no_dr")

	for k, v := range ms {
		fmt.Printf("k: %s, v: %v\n", k, v)
		for i, val := range v {
			fmt.Printf("\tindex position: %d \t value: %s\n", i, val)
		}
	}
}
