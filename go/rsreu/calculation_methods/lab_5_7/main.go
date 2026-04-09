package main

import (
	"fmt"
	"math"

	"github.com/sokolovea/code_practice/go/rsreu/calculation_methods/lab5"
)

func f(x float64) float64 {
	return math.Log(x+2) + math.Cos(2+x) - 0.35
}

func main() {
	fmt.Printf("%v\n", lab5.BruteForceDummy(0, 2, f, 0.005))
	fmt.Printf("%v\n", lab5.BitwiseSearch(0, 2, f, 0.005))
}
