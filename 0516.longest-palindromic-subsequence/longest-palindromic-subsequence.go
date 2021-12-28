package _516_longest_palindromic_subsequence

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		if dp[i] == nil {
			dp[i] = make([]int, len(s))
		}
		for j := 0; j < len(s); j++ {
			if i == j {
				dp[i][j] = 1
			}
		}
	}
	// 这里需注意坐标
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
