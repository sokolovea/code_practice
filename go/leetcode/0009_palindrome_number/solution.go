package main

import "math"

// Maybe not 100% optimal because float
// 3 ms Beats 57.13%
func isPalindrome(x int) bool {
	var result bool = true
	if x < 0 {
		result = false
	} else {
		var powLogTenByX int = int(math.Pow10(int(math.Log10(float64(x)))))
		var divider int = 1
		for powLogTenByX > 0 {
			left := x / divider % 10
			right := (x / powLogTenByX) % 10
			if left != right {
				result = false
				break
			}
			powLogTenByX /= 10
			divider *= 10
		}
	}
	return result
}
