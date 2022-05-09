package _300_longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	var res int

	dp := make([]int, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}

	for i := 0; i < len(dp); i++ {
		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
