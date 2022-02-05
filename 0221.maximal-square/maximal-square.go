package _221_maximal_square

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))

	square := 0
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))

		for j := 0; j < len(dp[i]); j++ {
			x := matrix[i][j] - '0'
			if x == 1 {
				dp[i][j] = 1
				square = 1
			}
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			// 这个if条件需要注意一下，当是1的时候才需要计算
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > square {
					square = dp[i][j]
				}
			}
		}
	}

	return square * square
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
