package _122_best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		// dp[1][0] 今天不持有股票的最大利润
		// = max(昨天不持有股票的最大利润，昨天持有股票且今天卖了的最大利润)
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// dp[i][1] 今天持有持股的最大利润
		// = max(昨天持有股票的最大利润，昨天不持有股票且今天买了股票的最大利润)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
