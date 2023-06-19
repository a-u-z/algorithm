package main

// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
func searchRange(nums []int, target int) []int {

	left, right := 0, len(nums)-1
	// 先設定 edge case 當做初始值
	// 要有初始值的想法，每個東西都要有初始值
	start, end := -1, -1

	// 重點在於逼近的方式

	// 找到第一個等於 target 的元素位置
	// 先問有沒有大於
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		if nums[mid] == target {
			start = mid
		}
	}

	// 找到最後一個等於 target 的元素位置
	left, right = 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		if nums[mid] == target {
			end = mid
		}
	}

	return []int{start, end}
}
