package main

// 通過兩個指針，一個快一個慢，在一個 for 迴圈當中，完成需要兩個 for 迴圈的事情
// 又稱為快慢指針
// 快指針：
func removeElement(nums []int, val int) int {
	// 遇到的 val 要刪除掉
	n := len(nums)

	resultPointer := 0 // 遇到目標的時候，指針不動
	for i := 0; i < n; i++ {
		if nums[i] == val {
		} else {
			nums[resultPointer] = nums[i]
			resultPointer++
		}
	}
	return resultPointer
}
