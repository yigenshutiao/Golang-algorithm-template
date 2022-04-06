package _053_maximum_subarray

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums)+1)

	dp[0] = nums[0]

	for i := 1; i < len(dp); i++ {
		dp[i] = -9999999999
	}
	// dp[i] 以nums[i] 为下标的最大子字符串的和
	// 对于每个数字来说，只有两种选择：只要当前的数字，当前数字+上次最大的值
	// 初始化dp[0]为nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}

	res := -999999
	// 没办法找到最大值，再遍历一次
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
