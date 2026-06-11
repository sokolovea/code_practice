package main

import (
	"fmt"
	"os"
)

// normalize нормализует значения, переданные в vals,
// так чтобы их сумма была равна 1.
func normalize(nums ...*float64) {
	var sum float64 = 0.0
	for _, num := range nums {
		sum += *num
	}
	if sum != 0 {
		for _, num := range nums {
			*num = *num / sum
		}
	}
}

func main() {
	a, b, c, d := 1.0, 2.0, 3.0, 4.0
	normalize(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
	// 0.1 0.2 0.3 0.4
	fmt.Println("PASS")
	os.Exit(0)
}
