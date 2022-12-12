package _010_regular_expression_matching

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	s = " " + s
	p = " " + p
	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	// 为了表达式中包含 f[0][0] 的时候计算正确，f[0][0] 置为 true, f[0][0] 无实际意义

	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j+1 < m && p[j+1] == '*' {
				continue // 如果 p 中的下一个字符是 *，不计算，放到和 * 一起计算
			}
			if i > 0 && p[j] != '*' {
				f[i][j] = f[i-1][j-1] && (s[i] == p[j] || p[j] == '.')
			} else if p[j] == '*' {
				f[i][j] = (f[i][j-2]) || ((i > 0 && f[i-1][j]) && (s[i] == p[j-1] || p[j-1] == '.'))
			}
		}
	}
	return f[n][m]
}
