package _301_remove_invalid_parentheses

func removeInvalidParentheses(s string) []string {
	var leftRem, rigtRem int
	// 计算出要删除的左括号和右括号
	for _, v := range s {
		if v == '(' {
			leftRem++
		} else if v == ')' && leftRem > 0 {
			leftRem--
		} else if v == ')' && leftRem == 0 {
			rigtRem++
		}
	}
	var res []string
	var dfs func(res *[]string, index, leftRem, rigtRem int, cur string)
	dfs = func(res *[]string, index, leftRem, rigtRem int, cur string) {
		if leftRem == 0 && rigtRem == 0 {
			// 判断是否有效
			if valid(cur) {
				*res = append(*res, cur)
			}
			return
		}

		// 这里移动idx
		for i := index; i < len(cur); i++ {
			// 这里和三数之和的方法差不多，因为是递归，所以是大于idx
			if i > index && cur[i] == cur[i-1] {
				continue
			}

			// 这里不移动idx，在外层统一移动
			if cur[i] == '(' && leftRem > 0 {
				dfs(res, i, leftRem-1, rigtRem, cur[:i]+cur[i+1:])
			}

			if cur[i] == ')' && rigtRem > 0 {
				dfs(res, i, leftRem, rigtRem-1, cur[:i]+cur[i+1:])
			}
		}
	}

	dfs(&res, 0, leftRem, rigtRem, s)

	return res
}

func valid(ss string) bool {
	var cnt int
	for _, s := range ss {
		if s == '(' {
			cnt++
		} else if s == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}

	return cnt == 0
}
