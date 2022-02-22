package _079_word_search

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	isUsed := make([][]bool, m)
	for i := 0; i < m; i++ {
		isUsed[i] = make([]bool, n)
	}

	var dfs func(i, j int, l int) bool
	dfs = func(i, j int, l int) bool {
		if l == len(word) {
			return true
		}
		// 单次搜索边界是否有效
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		// 当前单词是否已使用，是否和目标相等
		if isUsed[i][j] == true || board[i][j] != word[l] {
			return false
		}
		// 用了就置为true
		isUsed[i][j] = true

		cur := dfs(i+1, j, l+1) || dfs(i, j+1, l+1) || dfs(i-1, j, l+1) || dfs(i, j-1, l+1)
		// 这里应该是搜一半，然后断了，需要返回结果继续搜
		if cur == false {
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
