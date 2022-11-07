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

func removeInvalidParenthese(s string) []string {
	var res []string
	// 先算出要删除的括号
	var leftRemove, rightRemove int

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftRemove++
		}
		if s[i] == ')' && leftRemove == 0 {
			rightRemove++
		} else if s[i] == ')' && leftRemove > 0 {
			leftRemove--
		}
	}

	// 删除括号，看是否有效
	var dfs func(idx int, leftRemove int, rightRemove int, cur string, res *[]string)

	dfs = func(idx int, leftRemove int, rightRemove int, cur string, res *[]string) {
		if leftRemove == 0 && rightRemove == 0 {
			if valid(cur) {
				*res = append(*res, cur)
			}
			return
		}
		for i := idx; i < len(cur); i++ {
			if i > idx && cur[i] == cur[i-1] {
				continue
			}
			if cur[i] == '(' && leftRemove > 0 {
				dfs(i, leftRemove-1, rightRemove, cur[:i]+cur[i+1:], res)
			}

			if cur[i] == ')' && rightRemove > 0 {
				dfs(i, leftRemove, rightRemove-1, cur[:i]+cur[i+1:], res)
			}
		}
	}

	dfs(0, leftRemove, rightRemove, s, &res)

	// 返回结果
	return res
}
