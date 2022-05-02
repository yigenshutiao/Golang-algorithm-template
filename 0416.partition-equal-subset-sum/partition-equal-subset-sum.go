package _416_partition_equal_subset_sum

//canPartition 0-1背包问题
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 和为单数，肯定不能被均分
	if sum%2 == 1 {
		return false
	}
	// 找出目标值
	target := sum / 2
	dp := make([]int, target+1)

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[target] == target
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
