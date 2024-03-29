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

		//搜完了当前的，继续搜下一个字段
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

func exists(board [][]byte, word string) bool {

	if len(board) < 1 {
		return false
	}

	used := make([][]bool, len(board))

	for i := 0; i < len(used); i++ {
		used[i] = make([]bool, len(board[0]))
	}

	var dfs func(i, j, cur int) bool
	dfs = func(i, j, cur int) bool {
		// 返回条件
		if cur == len(word) {
			return true
		}
		// 剪枝
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return false
		}

		if used[i][j] == true || board[i][j] != word[cur] {
			return false
		}

		used[i][j] = true

		cu := dfs(i+1, j, cur+1) || dfs(i-1, j, cur+1) || dfs(i, j-1, cur+1) || dfs(i, j+1, cur+1)
		if cu {
			return true
		} else {
			used[i][j] = false
			return false
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}
