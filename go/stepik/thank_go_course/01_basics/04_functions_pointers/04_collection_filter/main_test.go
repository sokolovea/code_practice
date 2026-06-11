package main

import (
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	evenPredicate := func(number int) bool {
		return number%2 == 0
	}
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 3, 5}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 4}},
		{[]int{2, -2, -6, 4}, []int{2, -2, -6, 4}},
	}
	for _, test := range tests {
		got := filter(evenPredicate, test.input)
		if !slices.Equal(got, test.want) {
			t.Errorf("filter(%v) = %v; but got = %v", test.input, test.want, got)
		}
	}
}
