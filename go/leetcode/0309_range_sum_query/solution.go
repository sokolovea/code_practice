package main

type NumArray struct {
	nums       []int
	prefixSums []int
}

func Constructor(nums []int) NumArray {
	var result NumArray
	result.nums = nums
	result.prefixSums = make([]int, len(result.nums))
	sum := 0
	for i, r := range nums {
		result.prefixSums[i] = sum
		sum += r
	}
	return result
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.nums[right] + this.prefixSums[right] - this.prefixSums[left]
}
