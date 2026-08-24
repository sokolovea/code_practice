package main

import (
	"bufio"
	"fmt"
	"os"
)

// начало решения

// readLines возвращает все строки из указанного файла
func readLines(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	resultLines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newLine := scanner.Text()
		if newLine != "" {
			resultLines = append(resultLines, newLine)
		}
	}
	if scanner.Err() != nil {
		return []string{}, err
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
