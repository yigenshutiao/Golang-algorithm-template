package _032_longest_valid_parentheses

func longestValidParentheses(s string) int {
	var res int
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = 2 + dp[i-2]
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				// 这里只需要看s[i-dp[i-1]-2] 是否存在即可
				if (i - dp[i-1] - 2) >= 0 {
					dp[i] = dp[i-1] + 2 + dp[i-2-dp[i-1]]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
