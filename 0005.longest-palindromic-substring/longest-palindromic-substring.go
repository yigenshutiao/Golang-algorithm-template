package _005_longest_palindromic_substring

func longestPalindrome(s string) string {
	var res string

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	// 从下往上，从左往右
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					if j-i+1 > len(res) {
						res = s[i : j+1]
					}
				} else {
					if dp[i+1][j-1] == true {
						dp[i][j] = true
						if j-i+1 > len(res) {
							res = s[i : j+1]
						}
					}
				}
			}
		}
	}

	return res
}
