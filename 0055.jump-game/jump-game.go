package _055_jump_game

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))

	dp[0] = true

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] == true && nums[j] >= i-j {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(nums)-1]
}

func canJumps(nums []int) bool {
	k := 0

	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}

		// i + num[i]：已经跳的位置(i), 还能跳多远(num[i]), 和为最远能到达的距离
		k = max(k, i+nums[i])
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
