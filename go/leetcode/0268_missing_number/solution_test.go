package missingnumber

import (
	"fmt"
	"testing"
)

var testData = []struct {
	inputArray []int
	result     int
}{
	{[]int{1}, 0},
	{[]int{0, 2}, 1},
	{[]int{1, 2}, 0},
	{[]int{0, 1, 2}, 3},
	{[]int{0, 2, 3}, 1},
	{[]int{0, 1, 2, 3, 4}, 5},
}

var testFunc = []struct {
	f func([]int) int
}{
	{missingNumberSorting},
	{missingNumberInvariant},
}

func TestMissingNumber(t *testing.T) {
	for i, f := range testFunc {
		for j, testCase := range testData {
			t.Run(fmt.Sprintf("Test func %d on data %d: ", i, j),
				func(t *testing.T) {
					got := f.f(testCase.inputArray)
					if got != testCase.result {
						t.Errorf("Test case %d: array = %v, got = %d, expected = %d!",
							j, testCase.inputArray, got, testCase.result)
					}
				})
		}
	}
}
