package lab8

import (
	"math"

	c "github.com/sokolovea/code_practice/go/rsreu/calculation_methods/common"
)

func BruteForceDummy(leftX1 float64, rightX1 float64,
	leftX2 float64, rightX2 float64,
	f c.F2d, eps float64) c.MinimizeResult2d {
	x1Min, x2Min := leftX1, leftX2
	fMin := math.Inf(+1)
	var i uint64 = 0
	for x1Curr := leftX1; x1Curr <= rightX1-eps; x1Curr += eps {
		for x2Curr := leftX2; x2Curr <= rightX2-eps; x2Curr += eps {
			fCurr := f(x1Curr, x2Curr)
			if fCurr < fMin {
				fMin = fCurr
				x1Min = x1Curr
				x2Min = x2Curr
			}
			i++
		}
	}
	return c.MinimizeResult2d{X1Min: x1Min, X2Min: x2Min,
		FuncMin: fMin, FuncCallCounter: i}
}

func CoordinateDescent(leftX1 float64, rightX1 float64,
	leftX2 float64, rightX2 float64, f c.F2d,
	gX1 c.F1d, gX2 c.F1d, eps float64) c.MinimizeResult2d {
	x1Min, x2Min := leftX1, leftX2
	var x1MinPrev, x2MinPrev float64
	fMin := math.Inf(+1)
	var i uint64 = 0
	var isMinimizingX1 bool = true
	for {
		x1MinPrev, x2MinPrev = x1Min, x2Min
		if isMinimizingX1 {
			fX1 := getFX1(f, x2Min)
			result := bisectionInner(leftX1, rightX1, fX1, gX1, eps)
			x1Min = result.XMin
			i += result.FuncCallCounter
		} else {
			fX2 := getFX2(f, x1Min)
			result := bisectionInner(leftX2, rightX2, fX2, gX2, eps)
			x2Min = result.XMin
			i += result.FuncCallCounter
		}
		fCurr := f(x1Min, x2Min)
		i += 1
		if fCurr < fMin {
			fMin = fCurr
		}
		isMinimizingX1 = !isMinimizingX1
		if math.Abs(x1MinPrev-x1Min) <= eps &&
			math.Abs(x2MinPrev-x2Min) <= eps {
			break
		}
	}
	return c.MinimizeResult2d{X1Min: x1Min, X2Min: x2Min,
		FuncMin: fMin, FuncCallCounter: i}
}

func bisectionInner(left float64, right float64,
	f c.F1d, g c.F1d, eps float64) c.MinimizeResult1d {
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
	return c.MinimizeResult1d{XMin: xCenterCurr, FuncMin: f(xCenterCurr), FuncCallCounter: i + 1}
}

func getFX1(f c.F2d, x2 float64) c.F1d {
	return func(x1 float64) float64 {
		return f(x1, x2)
	}
}

func getFX2(f c.F2d, x1 float64) c.F1d {
	return func(x2 float64) float64 {
		return f(x1, x2)
	}
}

func GradientDescent(startX1 float64, startX2 float64,
	f c.F2d, fGrad c.F2dGrad, startStep float64,
	lambda float64, eps float64) c.MinimizeResult2d {
	x1Min, x2Min := startX1, startX2
	fMin := math.Inf(+1)
	var i uint64 = 0
	for {
		gradX1, gradX2 := fGrad(x1Min, x2Min)
		x1Min = x1Min - startStep*gradX1
		x2Min = x2Min - startStep*gradX2

		startStep *= lambda

		fCurr := f(x1Min, x2Min)
		i += 3
		if fCurr < fMin {
			fMin = fCurr
		}

		gradNorm := math.Sqrt(gradX1*gradX1 + gradX2*gradX2)
		if gradNorm < eps {
			break
		}
	}
	return c.MinimizeResult2d{X1Min: x1Min, X2Min: x2Min,
		FuncMin: fMin, FuncCallCounter: i}
}

func sign(x float64) float64 {
	if x > 0 {
		return 1
	} else {
		return -1
	}
}
