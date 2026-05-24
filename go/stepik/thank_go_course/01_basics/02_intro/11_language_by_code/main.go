/*
 * Напишите программу, которая определяет название языка по его коду.
 * Правила:
 * en → English
 * fr → French
 * ru или rus → Russian
 * иначе → Unknown
 */

package main

import (
	"fmt"
	"log"
)

func main() {
	var code string
	_, err := fmt.Scan(&code)
	if err != nil {
		log.Fatal("Scan parse error!")
		return
	}
	var language string
	switch code {
	case "en":
		language = "English"
	case "fr":
		language = "French"
	case "ru", "rus":
		language = "Russian"
	default:
		language = "Unknown"
	}
	fmt.Println(language)
}
