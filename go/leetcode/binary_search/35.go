package main

// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.

// 需要 O(log n) 的原因是，如果單存的遍歷這個 array 那時間複雜度會是 O(n)，要比這個好的只有 O(log n) 跟 O(1)，那 O(1) 沒有辦法達成，所以要 O(log n)

// 法一：遍歷 array
// 法二：二元搜尋 binary search
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 左界、右界
	// 設定左右界的目的是找到 middle
	for left <= right {
		middle := left + (right-left)/2
		// Q: 除不盡怎麼辦？
		// A: 誰來當 middle 都沒有關係，只要還是分成左區間與右區間即可，除不盡會捨去

		if target < nums[middle] { // target 不會是中間那個，而且比 nums[middle] 小，所以 target 在左區間，那重設的就是右界
			right = middle - 1 // 而且右界就可以是 middle 少一（因為要往左邊逼近，所以要用減的）
		} else if target > nums[middle] {
			left = middle + 1
		} else { // nums[middle] == target
			return middle
		}
	}
	return left
}
