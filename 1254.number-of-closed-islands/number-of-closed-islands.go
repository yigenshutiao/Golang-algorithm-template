package _254_number_of_closed_islands

func closedIsland(grid [][]int) int {
	if grid == nil {
		return 0
	}
	var res int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				if dfs(grid, i, j) {
					res += 1
				}
			}
		}
	}

	return res
}

func dfs(grid [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return false
	}

	if grid[i][j] == 1 {
		return true
	}

	grid[i][j] = 1

	up := dfs(grid, i, j-1)
	down := dfs(grid, i, j+1)
	left := dfs(grid, i-1, j)
	right := dfs(grid, i+1, j)

	return up && down && left && right
}
