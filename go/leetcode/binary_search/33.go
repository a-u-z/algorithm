package main

import "log"

// 0,1,2,3,4,5
// 1,2,3,4,5,0
// 2,3,4,5,0,1
// 3,4,5,0,1,2
// 4,5,0,1,2,3
// 5,0,1,2,3,4
func search33(nums []int, target int) int {
	numLen := len(nums)
	if numLen == 0 {
		return -1
	}

	left, right := 0, numLen-1
	for left <= right {
		mid := (left + right) / 2 // 無條件捨去
		log.Printf("here is left:%+v, mid:%+v, right:%+v", left, mid, right)

		if target == nums[mid] {
			return mid
		}
		// 先找出哪段是有序的
		// 這句話是關鍵，在沒有旋轉過的數列，右邊一定會大於左邊，這個無庸置疑
		// 那這題是旋轉過後的，代表有個斷點會是右邊沒有大於左邊，不過也一定會有某一段
		// 是保持著「右邊大於左邊」這個特性，所以還是可以用這個來判斷
		if nums[right] >= nums[mid] { // 代表右半段是有序排列的
			if target >= nums[mid] && target <= nums[right] {
				// 判斷目標元素是否在右半段，如果是的話要移動的是左界
				left = mid + 1
			} else { // 反之
				right = mid - 1
			}
		} else { // 代表左半段是有序的
			if target <= nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
