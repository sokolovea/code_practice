package main

import (
	"fmt"
	"os"
)

// начало решения

// readLines возвращает все строки из указанного файла
func readLines(name string) ([]string, error) {
	rawData, err := os.ReadFile(name)
	if err != nil {
		return []string{}, err
	}
	resultLines := make([]string, 0)
	currLineStartIndex := 0
	for i, c := range rawData {
		if c == '\n' {
			resultLines = append(resultLines, string(rawData[currLineStartIndex:i]))
			currLineStartIndex = i + 1
		}
	}
	return resultLines, nil
}

// конец решения

func main() {
	lines, err := readLines("/etc/passwd")
	if err != nil {
		panic(err)
	}
	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx+1, line)
	}
}
