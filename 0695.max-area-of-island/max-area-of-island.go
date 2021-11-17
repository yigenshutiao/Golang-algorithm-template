package _695_max_area_of_island

func maxAreaOfIsland(grid [][]int) int {
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				var tmp int
				dfs(grid, i, j, &tmp)
				if tmp > res {
					res = tmp
				}
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int, tmp *int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}

	if grid[i][j] == 0 {
		return
	}
	*tmp += 1

	grid[i][j] = 0

	dfs(grid, i, j-1, tmp)
	dfs(grid, i, j+1, tmp)
	dfs(grid, i-1, j, tmp)
	dfs(grid, i+1, j, tmp)
}
