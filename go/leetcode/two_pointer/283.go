package twopointer

func moveZeroes(nums []int) {
	// 概念是 left 當做完成的最尾部
	// 去蒐集非零的元素，當做完成的，一個一個放在 array 前面

	if len(nums) < 2 {
		return
	}

	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
}
