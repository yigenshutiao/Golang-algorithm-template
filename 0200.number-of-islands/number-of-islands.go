package _200_number_of_islands

func numIslands(grid [][]byte) int {
	var res int

	var dfs func(i, j int)

	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			return
		}

		if grid[i][j] == '1' {
			grid[i][j] = '2'
		} else {
			return
		}

		// 如果这个点是岛屿，往周边蔓延，染色所有岛屿
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}
