package _647_palindromic_substrings

func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					res++
				} else {
					if dp[i+1][j-1] == true {
						dp[i][j] = true
						res++
					}
				}
			}
		}
	}
	return res
}
