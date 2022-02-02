package _053_maximum_subarray

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums)+1)

	dp[0] = nums[0]

	for i := 1; i < len(dp); i++ {
		dp[i] = -9999999999
	}

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}

	res := -999999
	for _, num := range dp {
		if num > res {
			res = num
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
