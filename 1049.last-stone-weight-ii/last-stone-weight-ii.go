package _049_last_stone_weight_ii

func lastStoneWeightII(stones []int) int {
	target := 0
	for _, stone := range stones {
		target += stone
	}

	dp := make([]int, target+1)

	for i := 0; i < len(stones); i++ {
		for j := target / 2; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return target - (dp[target/2] * 2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
