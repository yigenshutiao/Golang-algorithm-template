package _200_number_of_islands

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}

	var res int

	// len(grid) 纵坐标  len(grid[0]) 横坐标
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res += 1
				dfs(grid, i, j)
			}
		}
	}

	return res
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}

	if grid[i][j] == '1' {
		grid[i][j] = '2'
	} else {
		return
	}

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
