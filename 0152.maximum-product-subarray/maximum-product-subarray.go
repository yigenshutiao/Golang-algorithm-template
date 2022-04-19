package _152_maximum_product_subarray

func maxProduct(nums []int) int {
	dp := make([][]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}

	// [0]存小于0的值，[1]存大于0的值
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			// 大于0, 正数*正数 求大，负数*正数 求小
			dp[i][1] = max(dp[i-1][1]*nums[i], nums[i])
			dp[i][0] = min(dp[i-1][0]*nums[i], nums[i])
		} else {
			// 小于0, 负数*负数 求大，正数*负数 求小
			dp[i][1] = max(dp[i-1][0]*nums[i], nums[i])
			dp[i][0] = min(dp[i-1][1]*nums[i], nums[i])
		}
	}

	res := -99999
	for i := 0; i < len(dp); i++ {
		if dp[i][1] > res {
			res = dp[i][1]
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
