// Выборка из генератора.
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
func countDigitsInWords(next func() string) counter {
	counted := make(chan pair)

	// считаем цифры в словах
	go func() {
		for {
			// как выйти из горутины,
			// когда слова закончились?
			word := next()
			if word == "" {
				close(counted)
				break
			}
			count := countDigits(word)
			counted <- pair{word, count}
		}
	}()

	// заполняем статистику
	stats := counter{}
	for {
		result := <-counted
		// как выйти из цикла,
		// когда слова закончились?
		if result.word == "" {
			break
		}
		// откуда взять слово?
		stats[result.word] = result.count
	}

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
