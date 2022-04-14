package _139_word_break

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)

	wd := make(map[string]bool)

	for _, word := range wordDict {
		wd[word] = true
	}

	// 长度为i的单词，如果在wordDict中出现过的话，结果为true
	dp[0] = true
	// dp[i] = dp[j] + s[j:i]
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
