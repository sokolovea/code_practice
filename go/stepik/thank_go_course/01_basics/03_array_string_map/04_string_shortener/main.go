/*
 * Напишите программу, которая укорачивает строку до указанной длины и добавляет в конце многоточие:
 * text = Eyjafjallajokull, width = 6 → Eyjafj...
 * Если строка не превышает указанной длины, менять ее не следует:
 * text = hello, width = 6 → hello
 * Гарантируется, что в исходной строке text используются только однобайтовые символы без пробелов, а длина width строго больше 0.
 */

package main

import (
	"fmt"
	"log"
)

func main() {
	var text string
	var width int
	_, err := fmt.Scanf("%s %d", &text, &width)
	if err != nil {
		log.Fatal("Error parsing text!")
		return
	}

	var byteArr []byte = []byte(text)
	if width < len(text) {
		byteArr = []byte(text[:width])
		byteArr = append(byteArr, '.', '.', '.')
	}
	var res string = string(byteArr)

	fmt.Println(res)
}
