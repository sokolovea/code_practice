package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	inputString string
	result      bool
}{
	{"()", true},
	{")()", false},
	{"(hello)", true},
	{"({[]})", true},
	{"([(])", false},
	{"([[[[[[[[]]]]]]]])", true},
	{"{}[]()", true},
	{"asas{asdf}d[asdf](asdf)asdfa", true},
	{"(([])[]()()", false},
	{"", true},
	{"sjkdfhjkasd", true},
	{"([{]}])", false},
}

func TestParenthesesCode(t *testing.T) {
	for i, data := range testData {
		result := isValid(data.inputString)
		assert.Equal(t, data.result, result, fmt.Sprintf("Test case %d failed: result for \"%s\" = %v but really is %v!", i, data.inputString, result, data.result))
	}
}
