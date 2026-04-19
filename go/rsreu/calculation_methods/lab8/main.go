package main

import (
	"fmt"

	lab8 "github.com/sokolovea/code_practice/go/rsreu/calculation_methods/lab8"
)

func f(x1 float64, x2 float64) float64 {
	return x1*x1 + 3*x2*x2 - 2*x2 + 3*x1 - 6
}

func fDerX1(x1 float64) float64 {
	return 2*x1 + 3
}

func fDerX2(x2 float64) float64 {
	return 6*x2 - 2
}

func fGrad(x1 float64, x2 float64) (float64, float64) {
	return fDerX1(x1), fDerX2(x2)
}

func main() {
	x1Left := -5.0
	x1Right := 5.0
	x2Left := -5.0
	x2Right := 5.0
	startStep := 1.0
	lambda := 0.9
	eps := 1e-3
	fmt.Printf("Searching minimum at X1 = [%v; %v]; X2 = [%v; %v] "+
		"with EPS = %0.7f\n", x1Left, x1Right, x2Left, x2Right, eps)
	fmt.Printf("1) lab8.BruteForceDummy:\n")
	fmt.Printf("%v\n",
		lab8.BruteForceDummy(x1Left, x1Right, x2Left, x2Right, f, eps))
	fmt.Printf("2) lab8.CoordinateDescent:\n")
	fmt.Printf("%v\n",
		lab8.CoordinateDescent(x1Left, x1Right, x2Left, x2Right, f, fDerX1, fDerX2, eps))
	fmt.Printf("3) lab8.GradientDescent:\n")
	fmt.Printf("%v\n",
		lab8.GradientDescent(x1Left, x1Right, f, fGrad, startStep, lambda, eps))
}
