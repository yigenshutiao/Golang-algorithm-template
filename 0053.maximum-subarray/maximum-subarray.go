package _053_maximum_subarray

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = -999999999
	}
	dp[0] = nums[0]

	// dp[i] 以nums[i] 为下标的最大子字符串的和
	// 对于每个数字来说，只有两种选择：只要当前的数字，当前数字+上次最大的值
	// 初始化dp[0]为nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}

	res := -9999999
	// 没办法找到最大值，需要再遍历一次
	for _, v := range dp {
		if v > res {
			res = v
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
