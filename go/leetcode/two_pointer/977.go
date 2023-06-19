package main

// 因為題目有說這是排序過的 slice ，所以兩端的數字（不看正負號）會是比較大的
// 所以可以用雙指針
func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left := 0
	right := n - 1
	index := n - 1 // 新的 slice 的 index

	for left <= right {
		if abs(nums[left]) > abs(nums[right]) {
			result[index] = nums[left] * nums[left]
			left++
		} else {
			result[index] = nums[right] * nums[right]
			right--
		}
		index--
	}
	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
