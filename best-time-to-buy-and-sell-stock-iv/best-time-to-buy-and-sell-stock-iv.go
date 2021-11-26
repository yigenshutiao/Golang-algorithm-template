package best_time_to_buy_and_sell_stock_iv

// k = 2, prices = [3,2,6,5,0,3]
func maxProfit(k int, prices []int) int {
	dp := make([][][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, k+1)
		for j := k; j > 0; j-- {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < len(prices); i++ {
		for j := k; j >= 1; j-- {
			if i == 0 {
				dp[0][k][0] = 0
				dp[0][k][1] = -prices[i]
				continue
			}

			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	return dp[len(prices)-1][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
