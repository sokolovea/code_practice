package main

import (
	"fmt"
	"math/rand"
)

// начало решения

type pair struct {
	initial string
	handled string
}

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case <-cancel:
				return
			case out <- randomWord(5):
			}
		}
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for word := range in {
			if !areLettersRepeated(word) {
				select {
				case <-cancel:
					return
				case out <- word:
				}
			}
		}
	}()
	return out
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan pair {
	out := make(chan pair)
	go func() {
		defer close(out)
		for word := range in {
			select {
			case <-cancel:
				return
			case out <- pair{word, reverseWord(word)}:
			}
		}
	}()
	return out
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan pair) <-chan pair {
	out := make(chan pair)
	go func() {
		defer close(out)
		for {
			select {
			case <-cancel:
				return
			case word := <-c1:
				select {
				case <-cancel:
					return
				case out <- word:
				}
			case word := <-c2:
				select {
				case <-cancel:
					return
				case out <- word:
				}
			}
		}
	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan pair, n int) {
	for range n {
		select {
		case <-cancel:
			return
		case wordPair := <-in:
			fmt.Printf("%s -> %s\n", wordPair.initial, wordPair.handled)
		}
	}
}

func areLettersRepeated(word string) bool {
	runes := []rune(word)
	runesMap := make(map[rune]struct{}, len(runes)/2)
	for _, rune := range word {
		_, ok := runesMap[rune]
		if ok {
			return true
		}
		runesMap[rune] = struct{}{}
	}
	return false
}

func reverseWord(word string) string {
	runes := []rune(word)
	for i := range len(word) / 2 {
		j := len(word) - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
