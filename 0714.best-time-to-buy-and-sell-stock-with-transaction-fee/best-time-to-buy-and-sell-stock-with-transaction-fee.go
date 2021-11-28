package _714_best_time_to_buy_and_sell_stock_with_transaction_fee

func maxProfit(prices []int, fee int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			dp[0][0] = 0
			dp[0][1] = -prices[i] - fee
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)

	}
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
