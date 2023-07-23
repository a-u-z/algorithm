package sw

// "abcabcbb"
func lengthOfLongestSubstring(s string) int {
	// 要給出原字串中的部分連續字串
	// 要用到滑動窗戶
	// 需要雙指針
	left, right := 0, 0

	// 要判斷有沒有重複，需要一個 map 來記錄
	dict := map[byte]int{}
	result := 0

	// 要遍歷這個 slice
	for right < len(s) { // 右指針要遍歷完的意思
		if index, found := dict[s[right]]; found && index >= left {
			left = index + 1
		}
		result = max(result, right-left+1)
		dict[s[right]] = right
		right++
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
