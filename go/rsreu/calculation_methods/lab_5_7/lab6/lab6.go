package lab6

import (
	"math"

	c "github.com/sokolovea/code_practice/go/rsreu/calculation_methods/common"
)

func Dichotomy(left float64, right float64,
	f c.F, eps float64) c.MinimizeResult {
	var i uint64 = 0
	xCenterCurr := (right + left) / 2
	for right-left >= 2*eps {
		xLeftCurr := xCenterCurr - eps
		xRightCurr := xCenterCurr + eps
		yLeftCurr := f(xLeftCurr)
		yRightCurr := f(xRightCurr)
		if yLeftCurr < yRightCurr {
			right = xCenterCurr
		} else {
			left = xCenterCurr
		}
		i += 2
		xCenterCurr = (right + left) / 2
	}
	return c.MinimizeResult{XMin: xCenterCurr, YMin: f(xCenterCurr), FuncCallCounter: i + 1}
}

func GoldenRatio(left float64, right float64,
	f c.F, eps float64) c.MinimizeResult {
	var i uint64 = 0
	τ := (math.Sqrt(5) - 1) / 2
	distance := right - left
	xLeftCurr := right - τ*distance
	xRightCurr := left + τ*distance
	yLeftCurr := f(xLeftCurr)
	yRightCurr := f(xRightCurr)
	i += 2
	for distance >= 2*eps {
		if yLeftCurr <= yRightCurr {
			right = xRightCurr
			distance = right - left
			xRightCurr = xLeftCurr
			xLeftCurr = right - τ*distance
			yRightCurr = yLeftCurr
			yLeftCurr = f(xLeftCurr)
		} else {
			left = xLeftCurr
			distance = right - left
			xLeftCurr = xRightCurr
			xRightCurr = left + τ*distance
			yLeftCurr = yRightCurr
			yRightCurr = f(xRightCurr)
		}
		i += 1
	}
	xCenterCurr := (right + left) / 2
	return c.MinimizeResult{XMin: xCenterCurr, YMin: f(xCenterCurr), FuncCallCounter: i}
}

func Parabola(left float64, right float64,
	f c.F, eps float64) c.MinimizeResult {
	var i uint64 = 0
	x1 := left
	x3 := right
	x2Preresult := findX3ParabolaInner(x1, x3, f)
	i += x2Preresult.FuncCallCounter
	x2 := x2Preresult.XMin

	f1 := f(x1)
	f2 := f(x2)
	f3 := f(x3)
	i += 3
	a1 := (f2 - f1) / (x2 - x1)
	a2 := ((f3-f1)/(x3-x1) - (f2-f1)/(x2-x1)) / (x3 - x2)
	x3Parabola := (x1 + x2 - a1/a2) / 2
	f3Parabola := f(x3Parabola)
	var x3ParabolaPrev float64
	for {
		x3ParabolaPrev = x3Parabola
		x1 = x2
		x3 = x3Parabola
		x2Preresult := findX3ParabolaInner(x1, x3, f)
		i += x2Preresult.FuncCallCounter
		x2 = x2Preresult.XMin

		f1 = f2
		f2 = f(x2)
		f3 = f3Parabola

		a1 = (f2 - f1) / (x2 - x1)
		a2 = ((f3-f1)/(x3-x1) - (f2-f1)/(x2-x1)) / (x3 - x2)
		x3Parabola = (x1 + x2 - a1/a2) / 2
		f3Parabola = f(x3Parabola)

		i += 2

		if math.Abs(x3Parabola-x3ParabolaPrev) <= eps {
			break
		}
	}
	return c.MinimizeResult{XMin: x3Parabola, YMin: f3Parabola, FuncCallCounter: i}
}

func findX3ParabolaInner(left float64, right float64, f c.F) c.MinimizeResult {
	x1 := left
	x3 := right
	var x2 float64
	isX2Found := false
	i := uint64(0)

	x2 = (x3 + x1) / 2
	if areXCorrect(x1, x2, x3, f) {
		i += 3
		isX2Found = true
	}
	if !isX2Found {
		initStep := (right - left) / 5
		for x2temp := left + initStep; x2temp < right-initStep; x2temp++ {
			i += 3
			if areXCorrect(x1, x2temp, x3, f) {
				x2 = x2temp
				isX2Found = true
			}
		}
		if !isX2Found {
			tempResult := GoldenRatio(left, right, f, (x3-x1)/10)
			x2 = tempResult.XMin
			i += tempResult.FuncCallCounter
		}
	}
	return c.MinimizeResult{XMin: x2, FuncCallCounter: i}
}

func areXCorrect(x1, x2, x3 float64, f c.F) bool {
	return (f(x1) <= f(x2) && f(x2) < f(x3)) || (f(x1) < f(x2) && f(x2) <= f(x3))
}
