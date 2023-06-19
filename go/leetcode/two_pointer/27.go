package main

// 通過兩個指針，一個快一個慢，在一個 for 迴圈當中，完成需要兩個 for 迴圈的事情
// 又稱為快慢指針
// 快指針：
func removeElement(nums []int, val int) int {
	slowIndex := 0

	for i := 0; i < len(nums); i++ { // 效率與記憶體位置都比 range 好
		// 上面這個 i 是快指針
		if nums[i] != val {
			nums[slowIndex] = nums[i]
			slowIndex++
		}
	}
	nums = nums[:slowIndex]
	return slowIndex
}
