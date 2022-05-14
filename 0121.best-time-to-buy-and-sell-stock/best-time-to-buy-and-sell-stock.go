package _121_best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - price[i]) dp[i-1][0] 为0，忽略
	// dp[0][1] = -prices[0]
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		// 今天不持有股票，昨天不持有股票; 昨天持有股票，今天卖
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 今天持有股票，昨天持有股票; 昨天不持有股票，今天买
		// 买股票不用管之前的值，买的利润就是-prices[i]
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
