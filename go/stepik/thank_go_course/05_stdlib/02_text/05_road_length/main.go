package main

import (
	"strconv"
	"strings"
	"unicode"
)

// начало решения

// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
	resultDistance := 0
	for _, direction := range directions {
		for word := range strings.FieldsSeq(direction) {
			if len(word) >= 1 && unicode.IsDigit(rune(word[0])) {
				currentDistance, _ := strconv.ParseFloat(strings.TrimRight(word, "km"), 64)
				if strings.HasSuffix(word, "km") {
					currentDistance *= 1000
				}
				resultDistance += int(currentDistance)
			}
		}
	}
	return resultDistance
}

// конец решения
