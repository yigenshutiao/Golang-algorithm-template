package characterMatch

func bf1(s, pattern string) int {
	if len(s) == 0 || len(pattern) == 0 || len(s) < len(pattern) {
		return -1
	}
	n, m := len(s), len(pattern)
	for i := 0; i <= n-m; i++ {
		for j := 0; j < m; j++ {
			if s[i+j] == pattern[j] {
				if j == m-1 {
					return i
				} else {
					continue
				}
			} else {
				break
			}
		}
	}
	return -1
}

func bf2(main, pattern string) int {
	if len(main) == 0 || len(pattern) == 0 || len(main) < len(pattern) {
		return -1
	}
	n, m := len(main), len(pattern)
	for i := 0; i <= n-m; i++ {
		sub := main[i : i+m]
		if sub == pattern {
			return i
		}
	}
	return -1
}

func bf3(main, pattern string) bool {
	if len(main) == 0 || len(pattern) == 0 || len(main) < len(pattern) {
		return false
	}
	n, m := len(main), len(pattern)
	for i := 0; i <= n-m; i++ {
		flag := true
		for j := 0; j < m; j++ {
			if main[i+j] != pattern[j] {
				flag = false
				break
			} else {
				if j == m-1 {
					return flag
				}
			}
		}
	}
	return false
}
