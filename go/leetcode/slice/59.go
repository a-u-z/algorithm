package slice

func generateMatrix(n int) [][]int {

	// 1. 要輸出 matrix 所以先做一個符合題目的空 matrix
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	// 3. 需要有上下左右
	top, bottom, left, right := 0, n, 0, n

	// 4. 準備要填入的數字，從 1 開始
	num := 1

	// 5. 開始填入數字，要填到 n*n 所以要有 =
	// 6. 用一個九宮格，然後標出上下左右，走一次
	for num <= n*n {
		for i := left; i < right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		for i := top; i < bottom; i++ {
			matrix[i][right-1] = num
			num++
		}
		right--

		for i := right - 1; i >= left; i-- {
			matrix[bottom-1][i] = num
			num++
		}
		bottom--

		for i := bottom - 1; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++

	}

	// 2. return
	return matrix
}
