// Читатель и счетовод.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

// nextFunc возвращает следующее слово из генератора.
type nextFunc func() string

// counter хранит количество цифр в каждом слове.
// Ключ карты - слово, а значение - количество цифр в слове.
type counter map[string]int

// pair хранит слово и количество цифр в нем.
type pair struct {
	word  string
	count int
}

// начало решения

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next().
func countDigitsInWords(next nextFunc) counter {
	pending := make(chan string)
	counted := make(chan pair)

	// отправляет слова на подсчет
	go func() {
		for {
			word := next()
			pending <- word
			if word == "" {
				break
			}
		}
		close(pending)
	}()

	// считает цифры в словах
	go func() {
		for {
			word := <-pending
			counted <- pair{word, countDigits(word)}
			if word == "" {
				break
			}
		}
		close(counted)
	}()

	stats := counter{}

	for {
		result := <-counted
		if result.word == "" {
			break
		}
		stats[result.word] = result.count
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

// wordGenerator возвращает генератор,
// который выдает слова из фразы.
func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
