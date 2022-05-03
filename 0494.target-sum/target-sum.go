package _494_target_sum

import "sort"

// left + right = sum
// left - right = target
// left = (sum + target) >> 1
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if target > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	// 计算背包大小
	bag := (sum + target) / 2

	if bag < 0 {
		return 0
	}
	// 定义dp数组
	dp := make([]int, bag+1)

	// 初始化
	dp[0] = 1
	// 遍历顺序
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			//推导公式
			dp[j] += dp[j-nums[i]]
			//fmt.Println(dp)
		}
	}

	return dp[bag]
}

// findTargetSumWay 回溯法
func findTargetSumWay(nums []int, target int) int {
	var sum int

	for _, num := range nums {
		sum += num
	}

	if target > sum {
		return 0
	}

	if (target+sum)%2 == 1 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	tar := (target + sum) / 2
	var res int
	var traverse func(idx int, cur, tar int)
	traverse = func(idx int, cur, tar int) {
		if cur == tar {
			res += 1
		}

		for i := idx; i < len(nums) && cur+nums[i] <= tar; i++ {
			traverse(i+1, cur+nums[i], tar)
		}

	}

	traverse(0, 0, tar)

	return res
}
