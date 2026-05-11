package keyboardrow

import "strings"

func findWords(words []string) []string {
	rowRunesSlice := [][]rune{
		{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'},
		{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'},
		{'z', 'x', 'c', 'v', 'b', 'n', 'm'},
	}
	rowRunesMap := make(map[rune]int)
	for i, row := range rowRunesSlice {
		for _, c := range row {
			rowRunesMap[c] = i
		}
	}
	var resultWords []string = make([]string, 0, len(words))
	for _, word := range words {
		var isWordSuitable bool = true
		var firstLetterIndex int
		for j, c := range strings.ToLower(word) {
			if j == 0 {
				firstLetterIndex = rowRunesMap[c]
			} else {
				if firstLetterIndex != rowRunesMap[c] {
					isWordSuitable = false
					break
				}
			}
		}
		if isWordSuitable {
			resultWords = append(resultWords, word)
		}
	}
	return resultWords
}
