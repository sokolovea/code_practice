package main

import (
	"errors"
)

func getRomanLetterValue(letter rune) (int, error) {
	var result int
	var convertError error = nil
	switch letter {
	case 'I':
		result = 1
	case 'V':
		result = 5
	case 'X':
		result = 10
	case 'L':
		result = 50
	case 'C':
		result = 100
	case 'D':
		result = 500
	case 'M':
		result = 1000
	case '_':
		convertError = errors.New("Unknown literal")
	}
	return result, convertError
}

func isTwoRomanLettersSubstractionValid(letterFirst rune, letterSecond rune) bool {
	var isValid bool = false
	isValid = (letterFirst == 'I' && (letterSecond == 'V' || letterSecond == 'X') ||
		letterFirst == 'X' && (letterSecond == 'L' || letterSecond == 'C') ||
		letterFirst == 'C' && (letterSecond == 'D' || letterSecond == 'M'))
	return isValid
}

func processRuneQueue(runeQueue *[]rune, isEnd bool) (int, error) {
	var result int = 0
	var error error = nil
	if runeQueue != nil {
		var runeQueueLength int = len(*runeQueue)
		switch {
		case runeQueueLength == 0:
			{
				error = errors.New("Empty queue!")
			}
		case runeQueueLength == 1 && isEnd:
			{
				result, error = getRomanLetterValue((*runeQueue)[0])
				if error != nil {
					(*runeQueue) = (*runeQueue)[:0]
				}
			}
		case runeQueueLength == 2:
			{
				leftRune := (*runeQueue)[0]
				rightRune := (*runeQueue)[1]

				leftRuneValue, leftRuneError := getRomanLetterValue(leftRune)
				rightRuneValue, rightRuneError := getRomanLetterValue(rightRune)
				if leftRuneError != nil || rightRuneError != nil {
					error = leftRuneError
				}
				if isTwoRomanLettersSubstractionValid(leftRune, rightRune) {
					result += (rightRuneValue - leftRuneValue)
					(*runeQueue) = (*runeQueue)[:0]
				} else {
					result += leftRuneValue
					(*runeQueue)[0] = rightRune
					(*runeQueue) = (*runeQueue)[:1]
				}
			}
		}
	}
	return result, error
}

// 0 ms | Beats 100.00%
func romanToInt(s string) int {
	var result int = 0
	var parseError error = nil
	rune_queue := make([]rune, 0, 2)
	var tempResult int = 0
	for _, letter := range s {
		rune_queue = append(rune_queue, letter)
		tempResult, parseError = processRuneQueue(&rune_queue, false)
		if parseError != nil {
			break
		}
		result += tempResult
	}
	tempResult, parseError = processRuneQueue(&rune_queue, true)
	result += tempResult
	return result
}
