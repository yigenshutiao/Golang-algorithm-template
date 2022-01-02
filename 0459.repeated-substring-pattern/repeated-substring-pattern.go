package _459_repeated_substring_pattern

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i <= n/2; i++ {
		// i 代表子串，限制n%i==0，是因为n一定是子串的倍数
		if n%i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
