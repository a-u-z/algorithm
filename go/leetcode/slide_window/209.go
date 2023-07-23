package sw

import (
	"log"
	"math"
)

func SliminSubArrayLenV2(target int, nums []int) int {
	// 要用滑動窗口
	// 就需要左右指針
	// 窗戶由左右指針構成，不過總和是從 left 到 right-1
	// 順便宣告總和
	left, right, sum := 0, 0, 0

	// 需要一個登記總答案的地方，先設定最大
	resultLen := math.MaxInt32

	// 要遍歷這個 slice
	// 當右指針超過 nums 的長度代表遍歷完成
	for right < len(nums) {
		sum += nums[right]
		right++

		// 當總和已經大於或是等於目標的時候
		for sum >= target {
			// 要去比較這次的長度是否有更短
			// 有的話會去更新
			resultLen = min(right-left, resultLen)

			// 再移動左指針，看看這樣長度改變，內容物改變，值的變化
			sum -= nums[left]
			left++
		}
	}

	// 如果 minLength 沒有改變，表示沒有符合條件的子陣列，返回 0
	if resultLen == math.MaxInt32 {
		return 0
	}
	return resultLen
	// 因為這個 slice 只有被遍歷兩次，左右指針各一次，所以複雜度是 On
}

// [2,3,1,2,4,3]
//
//	-
//	- -
//	- - -
//	- - - -
//	  - - -
//	    - -
func SliminSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	minLength := math.MaxInt32 // 先預設最大
	log.Printf("here is minLength:%+v", minLength)
	currentSum := 0

	for right < len(nums) {
		// 從右側擴展窗口
		currentSum += nums[right]
		right++

		// 當窗口內元素和大於等於 s 時，縮小窗口的左側
		for currentSum >= target {
			// 更新最小長度
			minLength = min(minLength, right-left)
			// 縮小窗口的左側，以繼續尋找最小長度
			currentSum -= nums[left]
			left++
		}
	}

	// 如果 minLength 沒有改變，表示沒有符合條件的子陣列，返回 0
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
