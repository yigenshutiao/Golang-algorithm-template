package _213_house_robber_ii

// nums= []int{1,2,3,1}
func rob(nums []int) int {

	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// 偷了第一间，不能偷最后一间
	res1 := robs(nums[0 : len(nums)-1])
	// 偷了最后一间，不能偷第一间
	res2 := robs(nums[1:len(nums)])

	// 所以两种方案都试试，最后选择利润高的方案
	return max(res1, res2)
}
func robs(nums []int) int {
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[1] = nums[0]
			continue
		}
		if i == 1 {
			dp[2] = max(dp[0], nums[i])
		}

		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
	}

	return dp[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
