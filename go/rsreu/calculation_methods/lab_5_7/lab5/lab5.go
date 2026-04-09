package lab5

import (
	"math"

	c "github.com/sokolovea/code_practice/go/rsreu/calculation_methods/common"
)

/*
Найти точку минимума функции одной переменной, определенной вариантом задания,
методами перебора и поразрядного поиска с заданной погрешностью. Для каждого
метода вывести точку минимума, значение функции в точке минимума, число
вычислений значения функции в процессе минимизации. Проверку результата
выполнить с помощью встроенных к Matlab (Scilab) функций минимизации при
погрешности, в 100 раз меньшей заданной.
*/

func sign(number float64) float64 {
	if number > 0 {
		return 1
	} else if number == 0 {
		return 0
	}
	return -1
}

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
	return c.MinimizeResult{XMin: xMin, YMin: yMin, Iters: i}
}

func BitwiseSearch(left float64, right float64,
	f c.F, eps float64) c.MinimizeResult {
	tempEps := (right - left) / 4
	curResult, xCurr := bitwiseSearchInner(left, right, f, tempEps)
	for math.Abs(eps) <= math.Abs(tempEps) {
		if sign(tempEps) > 0 {
			right = xCurr
		} else {
			left = xCurr
		}
		tempEps = -(tempEps / 4)
		tempResult, tempXCurr := bitwiseSearchInner(left, right, f, tempEps)
		xCurr = tempXCurr
		curResult.XMin = tempResult.XMin
		curResult.YMin = tempResult.YMin
		curResult.Iters += tempResult.Iters
	}
	return curResult
}

func bitwiseSearchInner(left float64, right float64, f c.F, eps float64) (c.MinimizeResult, float64) {
	xCurr := left
	xMin := xCurr
	yMin := f(xMin)
	var i uint64 = 0
	for (xCurr-right)*sign(eps) <= 0 {
		yCurr := f(xCurr)
		if (yCurr-yMin)*sign(eps) <= 0 {
			yMin = yCurr
			xMin = xCurr
		} else {
			break
		}
		xCurr += eps
		i++
	}
	return c.MinimizeResult{XMin: xMin, YMin: yMin, Iters: i}, xCurr
}
