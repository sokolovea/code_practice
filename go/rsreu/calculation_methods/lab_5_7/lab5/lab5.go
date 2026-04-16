package lab5

import (
	"math"

	c "github.com/sokolovea/code_practice/go/rsreu/calculation_methods/common"
)

func BruteForceDummy(left float64, right float64,
	f c.F, eps float64) c.MinimizeResult {
	xCurr := left
	xMin := xCurr
	yMin := f(xMin)
	var i uint64 = 0
	for xCurr <= right {
		yCurr := f(xCurr)
		if yCurr < yMin {
			yMin = yCurr
			xMin = xCurr
		}
		xCurr += eps
		i++
	}
	return c.MinimizeResult{XMin: xMin, YMin: yMin, FuncCallCounter: i}
}

func BitwiseSearch(left float64, right float64,
	f c.F, eps float64) c.MinimizeResult {
	leftOrig := left
	rightOrig := right
	tempEps := (right - left) / 4
	curResult := bitwiseSearchInner(left, right, f, tempEps)
	xCurr := curResult.XMin
	for math.Abs(eps) <= math.Abs(tempEps) {
		tempEps = -(tempEps / 4)
		left, right = xCurr, left
		tempResult := bitwiseSearchInner(left, right, f, tempEps)
		xCurr = tempResult.XMin
		if xCurr+tempEps >= leftOrig && xCurr+tempEps <= rightOrig {
			xCurr += tempEps
		}
		curResult.XMin = tempResult.XMin
		curResult.YMin = tempResult.YMin
		curResult.FuncCallCounter += tempResult.FuncCallCounter
	}
	return curResult
}

func bitwiseSearchInner(left float64, right float64, f c.F, eps float64) c.MinimizeResult {
	xCurr := left
	xMin := xCurr
	yMin := f(xMin)
	var i uint64 = 1
	for (xCurr-right)*eps <= 0 {
		yCurr := f(xCurr)
		if yCurr-yMin <= 0 {
			yMin = yCurr
			xMin = xCurr
		} else {
			break
		}
		xCurr += eps
		i++
	}
	return c.MinimizeResult{XMin: xMin, YMin: yMin, FuncCallCounter: i}
}
