package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	sb := strings.Builder{}
	for reader.Scan() {
		lines := strings.Split(reader.Text(), " ")
		for i, _ := range lines {
			sb.WriteString(strings.Title(strings.ToLower(lines[i])))
			sb.WriteByte(' ')
		}
		fmt.Println(strings.TrimSpace(sb.String()))
	}
	err := reader.Err()
	if err != nil {
		panic(err)
	}
}
