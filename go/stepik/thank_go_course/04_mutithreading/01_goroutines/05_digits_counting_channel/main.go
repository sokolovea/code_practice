// Канал с результатами.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

// counter хранит количество цифр в каждом слове.
// Ключ карты - слово, а значение - количество цифр в слове.
type counter map[string]int

// начало решения

// countDigitsInWords считает количество цифр в словах фразы.
func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)
	counted := make(chan int)
	stats := make(counter)
	go func() {
		// sum +=
		for _, word := range words {
			counted <- countDigits(word)
		}
	}()

	for _, word := range words {
		stats[word] = <-counted
	}

	// Считайте значения из канала counted
	// и заполните stats.

	// В результате stats должна содержать слова
	// и количество цифр в каждом.

	return stats
}

// конец решения

// countDigits возвращает количество цифр в строке.
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats печатает количество цифр в словах.
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	stats := countDigitsInWords(phrase)
	printStats(stats)
}
