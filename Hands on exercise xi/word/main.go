// Package word provides custom functions for counting words in a string.
package word

import "strings"

// UseCount returns a the number of times each word appears in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
	//return len(strings.Fields(s))
}
