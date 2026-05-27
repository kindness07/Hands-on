package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "i'm 008"}

	xp := [][]string{jb, jm}

	for i, v := range xp {
		fmt.Println(i, v)
		for j, val := range v {
			fmt.Printf("\tindex position: %d \t value: %s\n", j, val)
		}
	}
}
