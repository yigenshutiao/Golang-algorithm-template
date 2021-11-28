package _198_house_robber

//input [2,7,9,3,1]
func rob(nums []int) int {
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[0] = nums[0]
			continue
		}
		if i == 1 {
			dp[1] = nums[1]
			continue
		}
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
