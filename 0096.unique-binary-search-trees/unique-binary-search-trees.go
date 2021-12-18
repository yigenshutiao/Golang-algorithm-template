package _096_unique_binary_search_trees

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			left := j
			right := i - j - 1
			dp[i] += dp[left] * dp[right]
		}
	}
	return dp[n]
}
