package main

import (
	"fmt"
	"math"

	"github.com/sokolovea/code_practice/go/rsreu/calculation_methods/lab5"
	"github.com/sokolovea/code_practice/go/rsreu/calculation_methods/lab6"
	"github.com/sokolovea/code_practice/go/rsreu/calculation_methods/lab7"
)

func f(x float64) float64 {
	return math.Log(x+2) + math.Cos(2+x) - 0.35
}

func g(x float64) float64 {
	return 1/(x+2) - math.Sin(2+x)
}

func g2(x float64) float64 {
	return -1/((x+2)*(x+2)) - math.Cos(2+x)
}

func main() {
	left := 0.0
	right := 2.0
	eps := 1e-3
	fmt.Printf("Searching minimum at [%v; %v] with EPS = %0.7f\n", left, right, eps)
	fmt.Printf("1) lab5.BruteForceDummy:\n")
	fmt.Printf("%v\n", lab5.BruteForceDummy(left, right, f, eps))
	fmt.Printf("2) lab5.BitwiseSearch:\n")
	fmt.Printf("%v\n", lab5.BitwiseSearch(left, right, f, eps))
	fmt.Printf("3) lab6.Dichotomy:\n")
	fmt.Printf("%v\n", lab6.Dichotomy(left, right, f, eps))
	fmt.Printf("4) lab6.GoldenRatio:\n")
	fmt.Printf("%v\n", lab6.GoldenRatio(left, right, f, eps))
	fmt.Printf("5) lab6.Parabola:\n")
	fmt.Printf("%v\n", lab6.Parabola(left, right, f, eps))
	fmt.Printf("6) lab7.Bisection:\n")
	fmt.Printf("%v\n", lab7.Bisection(left, right, f, g, eps))
	fmt.Printf("7) lab7.Chord:\n")
	fmt.Printf("%v\n", lab7.Chord(left, right, f, g, eps))
	fmt.Printf("8) lab7.Newton:\n")
	fmt.Printf("%v\n", lab7.Newton(left, right, f, g, g2, eps))
}
