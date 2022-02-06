package _139_word_break

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)

	wd := make(map[string]bool)

	for _, word := range wordDict {
		wd[word] = true
	}

	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wd[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}