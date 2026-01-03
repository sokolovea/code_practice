package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataPalindromes = []int{
	121,
	100001,
	10001,
	4,
	55,
}

var testDataNotPalindromes = []int{
	1000021,
	-121,
	10,
	49,
	39479328,
	-3,
}

func TestIsPalindrome_Valid(t *testing.T) {
	for i, data := range testDataPalindromes {
		result := isPalindrome(data)
		assert.True(t, result, fmt.Sprintf("Test case %d failed: result %v but %d is palindrome", i, result, data))
	}
}

func TestIsPalindrome_NotValid(t *testing.T) {
	for i, data := range testDataNotPalindromes {
		result := isPalindrome(data)
		assert.False(t, result, fmt.Sprintf("Test case %d failed: result %v but %d is not palindrome", i, result, data))
	}
}
