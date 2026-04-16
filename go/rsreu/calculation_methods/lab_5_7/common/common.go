package common

import (
	"fmt"
	"math"
)

// Common types

type F func(float64) float64

type MinimizeResult struct {
	XMin            float64
	YMin            float64
	FuncCallCounter uint64
}

func (minimizeResult MinimizeResult) String() string {
	return fmt.Sprintf("Xmin = %0.7f; Ymin = %0.7f; f(x) and derivative calls count = %v",
		minimizeResult.XMin, minimizeResult.YMin, minimizeResult.FuncCallCounter)
}

// Common functions and vars for tests (benchmarks)

func FTest(x float64) float64 {
	return math.Log(x+2) + math.Cos(2+x) - 0.35
}

func GTest(x float64) float64 {
	return 1/(x+2) - math.Sin(2+x)
}

func G2Test(x float64) float64 {
	return -1/((x+2)*(x+2)) - math.Cos(2+x)
}

var EpsilonsTest []float64 = []float64{1e-3, 1e-5, 1e-7}
