package common

import (
	"fmt"
)

// Common types

type F1d func(float64) float64

type F2d func(float64, float64) float64

type F2dGrad func(x1 float64, x2 float64) (float64, float64)

type MinimizeResult1d struct {
	XMin            float64
	FuncMin         float64
	FuncCallCounter uint64
}

type MinimizeResult2d struct {
	X1Min           float64
	X2Min           float64
	FuncMin         float64
	FuncCallCounter uint64
}

func (minimizeResult MinimizeResult2d) String() string {
	return fmt.Sprintf("X1min = %0.7f; X2min = %0.7f; FuncMin = %0.7f; f(x) and derivative calls count = %v",
		minimizeResult.X1Min, minimizeResult.X2Min, minimizeResult.FuncMin, minimizeResult.FuncCallCounter)
}

// Common functions and vars for tests (benchmarks)

func FTest(x1 float64, x2 float64) float64 {
	return x1*x1 + 3*x2*x2 - 2*x2 + 3*x1 - 6
}

func FDerX1Test(x1 float64, x2 float64) float64 {
	return 2*x1 + 3
}

func FDerX2Test(x1 float64, x2 float64) float64 {
	return 6*x2 - 2
}

var EpsilonsTest []float64 = []float64{1e-3, 1e-5, 1e-7}
