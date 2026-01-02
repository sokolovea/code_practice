package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataPalindromes = []string{
	"aba",
	"torttrot",
	"aaa",
	"",
	"k",
	"alla",
	"ALLA",
	"rsrsrsrsrsrsrsrsr",
}

func TestIsPalindrome_Valid(t *testing.T) {
	for i, data := range testDataPalindromes {
		result := isPalindrome(data)
		assert.True(t, result, fmt.Sprintf("Test case %d failed: result %v but %s is palindrome", i, result, data))
	}
}

var testDataNotPalindromes = []string{
	"jsdglsf",
	"djfsjdfgsjgfif",
	"ju",
	"torttort",
	"RTOS",
	"RESTAPIrestAPI",
}

func TestIsPalindrome_NotValid(t *testing.T) {
	for i, data := range testDataNotPalindromes {
		result := isPalindrome(data)
		assert.False(t, result, fmt.Sprintf("Test case %d failed: result %v but %s is not palindrome", i, result, data))
	}
}

var testDataLongestPalindrome = []struct {
	inputData string
	expected  []string
}{
	{"babad", []string{"bab", "aba"}},
	{"aba", []string{"aba"}},
	{"Ruble", []string{"R", "u", "b", "l", "e"}},
	{"shalahs", []string{"shalahs"}},
	{"aaac", []string{"aaa"}},
	{"caaa", []string{"aaa"}},
	{"", []string{""}},
}

func TestLongestPalindrome_Simple(t *testing.T) {
	for i, data := range testDataLongestPalindrome {
		result := longestPalindromeSimple(data.inputData)
		assert.True(t, assert.Contains(t, data.expected, result), fmt.Sprintf("Test case %d failed: for %q result %q not found in expected values", i, data.inputData, result))
	}
}

// func TestLongestPalindrome_Optimized(t *testing.T) {
// 	for i, data := range testDataLongestPalindrome {
// 		result := longestPalindromeOptimized(data.inputData)
// 		assert.True(t, assert.Contains(t, data.expected, result), fmt.Sprintf("Test case %d failed: for %q result %q not found in expected values", i, data.inputData, result))
// 	}
// }
