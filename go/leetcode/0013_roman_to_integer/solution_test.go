package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	roman   string
	integer int
}{
	{"I", 1},
	{"III", 3},
	{"IV", 4},
	{"VII", 7},
	{"XIX", 19},
	{"XXI", 21},
	{"XLV", 45},
	{"XLIX", 49},
	{"MCMXCIV", 1994},
}

func TestRomanToInt(t *testing.T) {
	for i, data := range testData {
		result := romanToInt(data.roman)
		assert.Equal(t, data.integer, result, fmt.Sprintf("Test case %d failed: for %v result %v but should be %v", i, data.roman, result, data.integer))
	}
}
