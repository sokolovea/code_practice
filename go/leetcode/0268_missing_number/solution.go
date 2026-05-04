package missingnumber

import "sort"

func missingNumberSorting(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var missingNum int = len(nums)
	for i, val := range nums {
		if val != i {
			return i
		}
	}
	return missingNum
}

func missingNumberInvariant(nums []int) int {
	var sumInvariant int = (len(nums) + 1) * len(nums) / 2
	for _, val := range nums {
		sumInvariant -= val
	}
	return sumInvariant
}

var missingNumber func(nums []int) int = missingNumberInvariant
