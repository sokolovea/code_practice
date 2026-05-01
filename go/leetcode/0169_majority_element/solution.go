package majorityelement

import "sort"

func majorityElementMap(nums []int) int {
	var elementsMap map[int]int = make(map[int]int)
	for _, num := range nums {
		elementsMap[num] += 1
	}
	var majorityNum, majorityNumIndex int = 0, 0
	for num, count := range elementsMap {
		if count > majorityNumIndex {
			majorityNum = num
			majorityNumIndex = count
		}
	}
	return majorityNum
}

// Modifies nums, that is problem for correct measurements while benching
// func majorityElementSort(nums []int) int {
// 	if len(nums) == 0 { // STUB
// 		return 0
// 	}
// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] < nums[j]
// 	})
// 	var tempCounter int = 1
// 	var majorityNumCounter int = 1
// 	var majorityNum int = nums[0]
// 	var predNum int = nums[0]
// 	for _, num := range nums[1:] {
// 		if predNum != num {
// 			predNum = num
// 			tempCounter = 1
// 		} else {
// 			tempCounter += 1
// 			if tempCounter > majorityNumCounter {
// 				majorityNum = num
// 				majorityNumCounter = tempCounter
// 			}
// 		}
// 	}
// 	return majorityNum
// }

// No in-place sorting, works with copy
func majorityElementSort(nums []int) int {
	if len(nums) == 0 { // STUB
		return 0
	}
	numsCopy := append([]int(nil), nums...)
	sort.Slice(numsCopy, func(i, j int) bool {
		return numsCopy[i] < numsCopy[j]
	})
	var tempCounter int = 1
	var majorityNumCounter int = 1
	var majorityNum int = numsCopy[0]
	var predNum int = numsCopy[0]
	for _, num := range numsCopy[1:] {
		if predNum != num {
			predNum = num
			tempCounter = 1
		} else {
			tempCounter += 1
			if tempCounter > majorityNumCounter {
				majorityNum = num
				majorityNumCounter = tempCounter
			}
		}
	}
	return majorityNum
}

func majorityElementBoyerMoore(nums []int) int {
	var candidate int = 0
	var counter int = 0
	for _, num := range nums {
		if counter == 0 {
			candidate = num
		}
		if num == candidate {
			counter++
		} else {
			counter--
		}
	}
	return candidate
}

var majorityElement func(nums []int) int = majorityElementBoyerMoore
