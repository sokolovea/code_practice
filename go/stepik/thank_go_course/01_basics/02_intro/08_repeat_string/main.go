/**
* Программа принимает на вход строку source и число times.
* Требуется склеить source саму с собой times раз и вернуть результат:
* source = x, times = 3 → xxx
* source = omm, times = 2 → ommomm
* Гарантируется, что введенное значение корректное
 */

package main

import (
	"fmt"
)

func main() {
	var source string
	var times int64
	fmt.Scanf("%v%v", &source, &times)
	for range times {
		fmt.Printf("%v", source)
	}
	fmt.Println()
}
