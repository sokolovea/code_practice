package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	var builder strings.Builder
	var charSlice [][]byte = make([][]byte, numRows)
	for i := range len(charSlice) {
		charSlice[i] = make([]byte, len(s))
	}
	builder.Grow(len(s))
	for _, currentSChar := range numRows {
		for charIndex := currentRowIndex; charIndex < len(s); charIndex += numRows {
			builder.WriteByte(s[charIndex])
		}
	}
	return builder.String()
}

func main() {
	fmt.Printf("(\"%s\"; %d) -> %s; result = %s\n", "PAYPALISHIRING", 3, convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR")
}
