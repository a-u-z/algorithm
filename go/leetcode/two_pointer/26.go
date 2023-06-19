package main

func removeDuplicates(nums []int) int {
	// edge case
	// slice 長度为 0 或 1 时，沒有重複的元素
	if len(nums) < 2 {
		return len(nums)
	}

	// 非重複元素的指针
	slowIndex := 1

	// 從第 1 個開始比
	for i := 1; i < len(nums); i++ {
		// 當前元素與前一个元素不相等时，複製到非重複元素的位置
		if nums[i] != nums[i-1] {
			nums[slowIndex] = nums[i]
			slowIndex++
		}
	}

	// 返回非重複元素的个数
	return slowIndex
}

func AA() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
}

func B(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slowIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[slowIndex] {

			slowIndex++
		}
	}
	return slowIndex
}
