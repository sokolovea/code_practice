package lab8

import (
	"math"

	c "github.com/sokolovea/code_practice/go/rsreu/calculation_methods/common"
)

func BruteForceDummy(leftX1 float64, rightX1 float64,
	leftX2 float64, rightX2 float64,
	f c.F, eps float64) c.MinimizeResult {
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
	return c.MinimizeResult{X1Min: x1Min, X2Min: x2Min,
		FuncMin: fMin, FuncCallCounter: i}
}
