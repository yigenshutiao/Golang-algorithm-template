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
