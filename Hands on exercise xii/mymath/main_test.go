package mymath

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}
	tests := []test{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 3, 4}, 2.5},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 3.5},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 4},
	}
	for _, v := range tests {
		got := CenteredAvg(v.data)
		if got != v.answer {
			t.Errorf("CenteredAvg(%v) == %v, want %v", v.data, got, v.answer)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 4, 6, 9}))
	// Output:
	// 4
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000})
	}
}
