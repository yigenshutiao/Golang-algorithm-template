package _123_best_time_to_buy_and_sell_stock_iii

func maxProfit(prices []int) int {
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j <= 2; j++ {
			dp[i][j] = make([]int, 3)
		}
	}

	for i := 0; i < len(prices); i++ {
		for j := 2; j >= 1; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}

			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[len(prices)-1][2][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
