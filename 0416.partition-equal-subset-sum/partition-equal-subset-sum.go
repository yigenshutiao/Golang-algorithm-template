package _416_partition_equal_subset_sum

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

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
