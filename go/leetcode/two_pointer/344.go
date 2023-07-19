package twopointer

func reverseString(s []byte) {
	// 左跟右無限對調，直到左不是左，右不是右
	// 然後記得左要往右，右要往左
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func Aa() {
	reverseString([]byte{})
}
