package main

func removeDuplicates(nums []int) int {
	// 1. 處理 edge case
	if len(nums) < 2 {
		return len(nums)
	}
	// 遇到相同的，指針不動
	// 遇到不同的，指針往前，並設置上不同的元素
	resultPointer := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[resultPointer] {

		} else {
			resultPointer++
			nums[resultPointer] = nums[i]
		}
	}
	return resultPointer + 1 // 要回傳這個 nums 的長度，所以要 +1
}
