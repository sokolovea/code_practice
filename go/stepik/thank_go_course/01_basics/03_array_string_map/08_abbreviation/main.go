/*
 * Напишите программу, которая принимает на вход фразу и составляет аббревиатуру по первым буквам слов:
 * Today I learned → TIL
 * Высшее учебное заведение → ВУЗ
 * Кот обладает талантом → КОТ
 * Если слово начинается не с буквы, игнорируйте его:
 * Ар 2 Ди #2 → АД
 * Разделителями слов считаются только пробельные символы. Дефис, дробь и прочие можно не учитывать:
 * Анна-Мария Волхонская → АВ
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	phrase := readString()
	words := strings.Fields(phrase)
	abbrSlice := make([]rune, 0, len(words))
	for _, word := range words {
		firstLetter := unicode.ToUpper([]rune(word)[0])
		if unicode.IsLetter(firstLetter) {
			abbrSlice = append(abbrSlice, firstLetter)
		}
	}

	abbr := string(abbrSlice)
	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
