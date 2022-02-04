package _518_coin_change_2

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	dp[0] = 1

	// 这个题是求组合，要先遍历背包元素
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			// 求组合元素的常用递推公式
			dp[j] += dp[j-coin]
		}
		fmt.Println(dp)
	}
	return dp[amount]
}
