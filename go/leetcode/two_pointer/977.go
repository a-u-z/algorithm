package twopointer

// 因為題目有說這是排序過的 slice ，所以兩端的數字（不看正負號）會是比較大的
// 所以可以用雙指針
func sortedSquares(nums []int) []int {
	// 2. 判斷 edge case
	if len(nums) < 2 {
		return nums
	}
	// 1. 因為是由小到大的排序，也就是兩端的數字本身（不看正負號）會是比較大的
	// 所以要先比較這兩端的數字，那就需要兩個
	n := len(nums) // 3
	result := make([]int, n)
	left, right, index := 0, n-1, n-1
	for left <= right { // 8
		if abs(nums[left]) > abs(nums[right]) { // 4
			result[index] = nums[left] * nums[left] // 5
			left++                                  // 6
		} else {
			result[index] = nums[right] * nums[right]
			right--
		}
		index-- // 7
	}
	return result // 0
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
