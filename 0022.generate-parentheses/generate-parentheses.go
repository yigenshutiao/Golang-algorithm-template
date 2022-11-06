package _022_generate_parentheses

func generateParenthesis(n int) []string {
	res := []string{}

	if n == 0 {
		return res
	}

	var dfs func(cur string, left, right int)
	dfs = func(cur string, left, right int) {
		if left == 0 && right == 0 {
			res = append(res, cur)
			return
		}
		// 这里相当于剪枝
		if left > right {
			return
		}
		if left > 0 {
			dfs(cur+"(", left-1, right)
		}
		if right > 0 {
			dfs(cur+")", left, right-1)
		}
	}

	dfs("", n, n)
	return res
}

func generateParenthesi(n int) []string {
	res := []string{}

	// 剩余左括号，剩余右括号
	left, right := n, n

	var dfs func(cur string, left, right int)
	dfs = func(cur string, left, right int) {
		// 终止条件
		if left == 0 && right == 0 {
			res = append(res, cur)
			return
		}

		// 减枝
		if left > right {
			return
		}

		// 单层逻辑

		// 左括号
		if left > 0 {
			dfs(cur+"(", left-1, right)
		}

		// 右括号
		if right > 0 {
			dfs(cur+")", left, right-1)
		}

	}

	dfs("", left, right)

	return res
}
