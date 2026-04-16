package lab7

import (
	"math"

	c "github.com/sokolovea/code_practice/go/rsreu/calculation_methods/common"
)

func Bisection(left float64, right float64,
	f c.F, g c.F, eps float64) c.MinimizeResult {
	var i uint64 = 0
	xCenterCurr := (right + left) / 2
	for right-left >= 2*eps {
		if g(xCenterCurr) < 0 {
			left = xCenterCurr
		} else {
			right = xCenterCurr
		}
		i++
		xCenterCurr = (right + left) / 2
	}
	return c.MinimizeResult{XMin: xCenterCurr, YMin: f(xCenterCurr), FuncCallCounter: i + 1}
}

func Chord(left float64, right float64,
	f c.F, g c.F, eps float64) c.MinimizeResult {
	var i uint64 = 0
	xInnerChord := (right-left)/(g(right)-g(left))*(-g(left)) + left
	i += 2
	for xInnerChord-left > eps {
		left = xInnerChord
		xInnerChord = (right-left)/(g(right)-g(left))*(-g(left)) + left
		i += 2
	}
	return c.MinimizeResult{XMin: xInnerChord, YMin: f(xInnerChord), FuncCallCounter: i + 1}
}

func Newton(left float64, right float64,
	f c.F, g c.F, g2 c.F, eps float64) c.MinimizeResult {
	var i uint64 = 0
	xInnerNewton := (right - left) / 2
	for {
		f1 := g(xInnerNewton)
		i += 1
		if math.Abs(f1) <= eps {
			break
		}
		f2 := g2(xInnerNewton)
		xInnerNewtonCopy := xInnerNewton
		xInnerNewton = xInnerNewtonCopy - f1/f2
		i += 1
	}
	return c.MinimizeResult{XMin: xInnerNewton, YMin: f(xInnerNewton), FuncCallCounter: i + 1}
}
