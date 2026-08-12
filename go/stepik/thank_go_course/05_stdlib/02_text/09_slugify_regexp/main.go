package main

import (
	"regexp"
	"strings"
	"testing"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовка:
// только латиница, цифры и дефис
func slugify(src string) string {
	src = strings.ToLower(src)
	regExpr, _ := regexp.Compile("[a-zA-Z0-9-]+")
	matches := regExpr.FindAllString(src, -1)
	return strings.Join(matches, "-")
}

// конец решения

func Test(t *testing.T) {
	const phrase = "Go Is Awesome!"
	const want = "go-is-awesome"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}
