package main

// 前題 1. slice 是有序的 2. 無重複元素
// 二分法的前題就是上述的兩個條件
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if target == nums[middle] {
			return middle
		} else if target < nums[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}
