package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataTree = []struct {
	inputSlice  []int
	outputSlice []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{4, 3, 10}, []int{3, 4, 10}},
	{[]int{8, 4, 2, 1, 5, 14, 12, 13, 16}, []int{1, 2, 4, 5, 8, 12, 13, 14, 16}},
}

func TestTreePreorder(t *testing.T) {
	for i, data := range testDataTree {
		var root *TreeNode = nil
		for _, number := range data.inputSlice {
			root = root.add(number)
		}
		result := preorderTraversal(root)
		assert.Equal(t, data.outputSlice, result, fmt.Sprintf("Test case %d failed: result slice %v but must be %v", i, result, data.outputSlice))
	}
}
