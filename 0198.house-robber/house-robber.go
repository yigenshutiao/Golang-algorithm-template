package _198_house_robber

//input [2,7,9,3,1]
func rob(nums []int) int {
	dp := make([][]int, len(nums))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	// [0]代表不偷， [1]代表偷
	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < len(nums); i++ {
		// 今天偷：昨天不偷，今天偷; 昨天偷
		dp[i][1] = max(dp[i-1][0]+nums[i], dp[i-1][1])
		// 今天不偷：昨天偷了, 昨天不偷
		dp[i][0] = max(dp[i-1][1], dp[i-1][0])
	}

	if dp[len(nums)-1][0] > dp[len(nums)-1][1] {
		return dp[len(nums)-1][0]
	}
	return dp[len(nums)-1][1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
