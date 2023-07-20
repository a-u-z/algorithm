package twopointer

import "sort"

// 以三包胎為單位，三包胎裡面可以有重複的元素
func threeSum(nums []int) [][]int {

	var result [][]int

	// 首先对数组进行排序
	sort.Ints(nums)

	// 遍历数组
	for i := 0; i < len(nums)-2; i++ { // -2 是因為最多可能的是左指針 -2 右指針 -1

		// 為了避免出現重複的三包胎
		// 因為是往右找，所以當左邊有一個跟自己相同的數字，最後算出的結果又是 0 的話，那就會是重複的三包胎
		if i > 0 && nums[i-1] == nums[i] { // >0 是因為要計算 i-1
			continue
		}

		// 使用双指针来找到和为0的三个数
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 跳过重复元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
