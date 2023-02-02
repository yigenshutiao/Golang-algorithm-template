package _042_trapping_rain_water

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// trapDP 接雨水dp解法
func trap(height []int) int {
	var res int

	maxLeft := make([]int, len(height))
	maxRigt := make([]int, len(height))

	maxLeft[0] = height[0]
	// 当前值左边界的最大值 为当前值和dp[i-1]的比较
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
	}

	// 当前值右边界的最大值，是当前值和dp[i+1]的比较
	maxRigt[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		maxRigt[i] = max(height[i], maxRigt[i+1])
	}

	for i := 1; i < len(height); i++ {
		high := min(maxRigt[i], maxLeft[i])
		sum := high - height[i]

		if sum > 0 {
			res += sum
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
