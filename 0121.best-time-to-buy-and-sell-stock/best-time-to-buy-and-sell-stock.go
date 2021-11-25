package _121_best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	dp := [][]int{}

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	// dp[0][1] = -prices[0] 这一行是干什么的...
	dp[0][1] = -999

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
