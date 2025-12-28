package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var sumIndexes []int = []int{0, 0}
out:
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				sumIndexes[0] = i
				sumIndexes[1] = j
				break out
			}
		}
	}
	return sumIndexes
}

// Input: count numbers, then target, then numbers
func main() {
	var count int
	fmt.Scanf("%d", &count)
	var target int
	fmt.Scanf("%d", &target)
	var numbers []int = make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &numbers[i])
	}
	resultArray := twoSum(numbers, target)
	fmt.Println(resultArray)
}
