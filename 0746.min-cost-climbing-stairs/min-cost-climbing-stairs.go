package _746_min_cost_climbing_stairs

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	return min(dp[len(cost)-2], dp[len(cost)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
