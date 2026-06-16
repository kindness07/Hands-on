package main

import (
	"fmt"

	"github.com/kindness07/Hands-on/Hands-on-exercise-xi/quote"
	"github.com/kindness07/Hands-on/Hands-on-exercise-xi/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
