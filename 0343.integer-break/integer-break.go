package _343_integer_break

func integerBreak(n int) int {
	dp := make([]int, n+1)

	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
