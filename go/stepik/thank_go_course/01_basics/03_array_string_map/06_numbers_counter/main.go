package main

import (
	"fmt"
	"sort"
)

func main() {
	var number int
	fmt.Scanf("%d", &number)

	// Посчитайте, сколько раз каждая цифра встречается
	// в числе `number`. Запишите результат в карту `counter`,
	// где ключом является цифра числа, а значением -
	// сколько раз она встречается
	counter := make(map[int]int)
	for {
		curNumber := number % 10
		number /= 10
		counter[curNumber] += 1
		if number == 0 {
			break
		}
	}
	printCounter(counter)
}

// key1:val1 key2:val2 ... keyN:valN
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
