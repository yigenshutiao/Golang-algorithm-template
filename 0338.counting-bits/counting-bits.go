package _338_counting_bits

func countBits(n int) []int {
	dp := make([]int, n+1)

	dp[0] = 0

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i/2]
		}
	}
	return dp
}
