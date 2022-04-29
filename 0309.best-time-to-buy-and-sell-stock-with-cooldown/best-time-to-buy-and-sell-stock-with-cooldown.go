package _309_best_time_to_buy_and_sell_stock_with_cooldown

//input [1,2,3,0,2]
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	// 第一个元素是天数，第二个元素，0没有股票，1代表有股票
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {

		if i == 1 {
			// 第一天的话，不持有股票： 前一天不持有股票，前一天持有股票，今天卖了
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			// 持有股票：前一天持有股票，前一天不持有股票，今天买
			dp[i][1] = max(dp[i-1][1], -prices[i])
			continue
		}

		// 今天不持有股票：昨天不持有股票，昨天持有股票，今天卖
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 今天持有股票：昨天持有股票，前天不持有股票，昨天冷冻期，今天买
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
