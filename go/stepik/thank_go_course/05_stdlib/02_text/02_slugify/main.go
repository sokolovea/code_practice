package main

import (
	"strings"
	"unicode"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовка:
// только латиница, цифры и дефис
func slugify(src string) string {
	splitFunc := func(r rune) bool {
		return !(unicode.IsLetter(r) && r <= unicode.MaxLatin1) && !unicode.IsNumber(r) && r != '-'
	}
	return strings.Join(strings.FieldsFunc(strings.ToLower(src), splitFunc), "-")
}
