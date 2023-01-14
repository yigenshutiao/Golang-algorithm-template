package offer12_ju_zhen_zhong_de_lu_jing_lcof

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	isUsed := make([][]bool, len(board))
	for i := 0; i < m; i++ {
		isUsed[i] = make([]bool, len(board[0]))
	}

	var dfs func(i, j int, idx int) bool

	dfs = func(i, j int, idx int) bool {
		if idx == len(word) {
			return true
		}

		if i < 0 || j < 0 || i >= m || j >= n {
			return false
		}

		if isUsed[i][j] == true || board[i][j] != word[idx] {
			return false
		}

		isUsed[i][j] = true
		next := dfs(i-1, j, idx+1) || dfs(i+1, j, idx+1) || dfs(i, j-1, idx+1) || dfs(i, j+1, idx+1)
		if !next {
			isUsed[i][j] = false
			return false
		} else {
			return true
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
