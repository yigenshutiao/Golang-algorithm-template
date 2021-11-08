package main

func dfs(y int, n int, pic []string, res *[][]string, col, bias, rbias []bool) {
	if y == n {
		t := make([]string, n)
		copy(t, pic)
		*res = append(*res, t)
	}

	for x := 0; x < n; x++ {
		if !col[x] && !bias[x+y] && !rbias[y-x+n] {
			col[x], bias[x+y], rbias[y-x+n] = true, true, true
			t := make([]byte, n)
			for k := range t {
				t[k] = '.'
			}
			t[x] = 'Q'

			pic[y] = string(t)

			dfs(y+1, n, pic, res, col, bias, rbias)

			col[x], bias[x+y], rbias[y-x+n] = false, false, false

			t[x] = '.'
		}
	}

}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	pic := make([]string, n)

	col := make([]bool, n)
	bias := make([]bool, n*2)
	rbias := make([]bool, n*2)

	dfs(0, n, pic, &res, col, bias, rbias)

	return res
}
