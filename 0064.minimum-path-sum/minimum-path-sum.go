package _064_minimum_path_sum

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	if m == 1 && n == 1 {
		return grid[m-1][n-1]
	}

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			dp[0][i] += grid[0][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j <= i; j++ {
			dp[i][0] += grid[j][0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
