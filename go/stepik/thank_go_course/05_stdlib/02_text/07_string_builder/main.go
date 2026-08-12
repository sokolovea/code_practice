package main

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

// начало решения

// prettify возвращает отформатированное
// строковое представление карты
func prettify(m map[string]int) string {
	strSlice := make([]string, len(m))
	i := 0
	for key, value := range m {
		strSlice[i] = key + ": " + strconv.Itoa(value)
		i++
	}
	slices.Sort(strSlice)

	var sb strings.Builder
	sb.WriteRune('{')
	if len(m) == 1 {
		sb.WriteRune(' ')
	}
	for _, str := range strSlice {
		if len(m) > 1 {
			sb.WriteString("\n    ")
		}
		sb.WriteString(str)
		if len(m) > 1 {
			sb.WriteRune(',')
		}
	}
	if len(m) == 1 {
		sb.WriteRune(' ')
	}
	if len(m) > 1 {
		sb.WriteRune('\n')
	}
	sb.WriteRune('}')
	return sb.String()
}

// конец решения

func Test(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	const want = "{\n    one: 1,\n    three: 3,\n    two: 2,\n}"
	got := prettify(m)
	if got != want {
		t.Errorf("%v\ngot:\n%v\n\nwant:\n%v", m, got, want)
	}
}
