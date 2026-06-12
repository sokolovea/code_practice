package main

import "testing"

func TestValidator(t *testing.T) {
	testData := []struct {
		validator   validator
		inputString string
		result      bool
	}{
		{digits, "kjsdhfkj", false},
		{digits, "", false},
		{digits, "27389423", true},
		{digits, "273asdf3s", true},
		{digits, "273ырвороf3", true},
		{letters, "kjsdhfkj", true},
		{letters, "", false},
		{letters, "27389423", false},
		{letters, "273asdf3s", true},
		{letters, "дфывлаодл", true},
		{minlen(5), "", false},
		{minlen(5), "2", false},
		{minlen(5), "1234", false},
		{minlen(5), "1234f", true},
		{minlen(5), "asdfu89i", true},
		{and(digits, letters), "asdfu89i", true},
		{and(digits, letters), "asdfasokjdfjo", false},
		{and(digits, letters), "ыфвра78ры8вшрвы", true},
		{or(digits, letters), "asdfasokjdfjo", true},
		{or(digits, letters), "нфы9врагш", true},
		{or(digits, letters), "офщшы", true},
		{or(and(digits, letters), minlen(10)), "", false},
		{or(and(digits, letters), minlen(10)), "kjshadjhf283974", true},
		{or(and(digits, letters), minlen(10)), "k9", true},
		{or(and(digits, letters), minlen(10)), "пловец", false},
		{or(and(digits, letters), minlen(10)), "оченьхорошийпловец", true},
		{or(and(digits, letters), minlen(10)), "hellowor1d", true},
	}
	for i, test := range testData {
		if got := test.validator(test.inputString); got != test.result {
			t.Errorf("Case %d: for %#v(%v) expected %v but got %v", i, test.validator, test.inputString, test.result, got)
		}
	}
}
