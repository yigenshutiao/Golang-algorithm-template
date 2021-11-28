package _188_best_time_to_buy_and_sell_stock_iv

// k = 2, prices = [3,2,6,5,0,3]
func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, k+1)
		for j := k; j >= 0; j-- {
			dp[i][j] = make([]int, 3)
		}
	}

	for i := 0; i < len(prices); i++ {
		dp[i][0][0] = 0
		dp[i][0][1] = -999999999
	}

	for i := 0; i < len(prices); i++ {
		for j := k; j >= 1; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}

			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
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
