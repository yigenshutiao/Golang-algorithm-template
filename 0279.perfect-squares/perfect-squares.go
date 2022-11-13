package _279_perfect_squares

func numSquares(n int) int {
	dp := make([]int, n+1)

	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = i + 1
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
