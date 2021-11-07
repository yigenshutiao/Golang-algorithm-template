package main

import "fmt"

// 硬币面额为：c1 ~ cn
// 状态转移方程：f(n) = min(f(n - c1), f(n - c2), ... f(n - cn)) + 1
func coinChange(coins []int, amount int) int {
	// dp[i] 金额为 i 时，所需零钱最小数量
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // 搞个无解的默认值
		for _, c := range coins {
			if i >= c {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
