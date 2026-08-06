package main

// не удаляйте импорты, они используются при проверке
import (
	"strings"
)

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// Библиотечная
func MatchContains(pattern string, src string) bool {
	return strings.Contains(src, pattern)
}

// Самописная
func MatchContainsCustom(pattern string, src string) bool {
	if pattern == "" {
		return true
	}
	if len(pattern) > len(src) {
		return false
	}
	pat_len := len(pattern)
	for idx := 0; idx < len(src)-pat_len+1; idx++ {
		if src[idx:idx+pat_len] == pattern {
			return true
		}
	}
	return false
}
