package keyboardrow

import (
	"fmt"
	"slices"
	"testing"
)

func TestFindWords(t *testing.T) {
	testData := []struct {
		inputWords  []string
		outputWords []string
	}{
		{[]string{"qwerty", "ewe", "adf", "qaz"}, []string{"qwerty", "ewe", "adf"}},
		{[]string{"er", "fv", "hk"}, []string{"er", "hk"}},
		{[]string{"hello", "not", "joke", "doll"}, []string{}},
		{[]string{"bad", "mad", "lag", "Potter"}, []string{"lag", "Potter"}},
		{[]string{"DAS", "AuTo", "Et"}, []string{"DAS", "Et"}},
	}

	for i, testCase := range testData {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := findWords(testCase.inputWords)
			if !slices.Equal(got, testCase.outputWords) {
				t.Errorf("Error test case: expected = %v; got = %v",
					testCase.outputWords, got)
			}
		})
	}
}
