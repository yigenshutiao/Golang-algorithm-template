package _301_remove_invalid_parentheses

func removeInvalidParentheses(s string) []string {
	// 计算出l和r需要删除的次数
	lResult, rResult := 0, 0
	for _, ss := range s {
		if ss == '(' {
			lResult += 1
		} else if ss == ')' && lResult > 0 {
			lResult -= 1
		} else if ss == ')' && lResult == 0 {
			rResult += 1
		}
	}

	// 开始删除
	var res []string
	dfs(&res, 0, lResult, rResult, s)

	return res
}

func dfs(res *[]string, startIndex int, lResult, rResult int, cur string) {
	if lResult == 0 && rResult == 0 {
		if valid(cur) {
			*res = append(*res, cur)
		}
		return
	}

	for i := startIndex; i < len(cur); i++ {
		// 避免相同的情况
		if i > startIndex && cur[i] == cur[i-1] {
			continue
		}

		if cur[i] == '(' && lResult > 0 {
			dfs(res, i, lResult-1, rResult, cur[:i]+cur[i+1:])
		}

		if cur[i] == ')' && rResult > 0 {
			dfs(res, i, lResult, rResult-1, cur[:i]+cur[i+1:])
		}
	}
}

func valid(s string) bool {
	cnt := 0
	for _, v := range s {
		if v == '(' {
			cnt++
		} else if v == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}
