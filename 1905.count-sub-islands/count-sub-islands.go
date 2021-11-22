package _905_count_sub_islands

func countSubIslands(grid1 [][]int, grid2 [][]int) int {

	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[0]); j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				dfs(grid2, i, j)
			}
		}
	}

	var res int

	for i := 0; i < len(grid2); i++ {
		for j := 0; j < len(grid2[0]); j++ {
			if grid2[i][j] == 1 {
				res += 1
				dfs(grid2, i, j)
			}
		}
	}

	return res
}

func dfs(grid2 [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid2) || j >= len(grid2[0]) {
		return
	}

	// 如果已经被淹掉了，可以返回
	if grid2[i][j] == 0 {
		return
	}

	// 否则把岛屿淹掉
	grid2[i][j] = 0

	dfs(grid2, i, j-1)
	dfs(grid2, i, j+1)
	dfs(grid2, i-1, j)
	dfs(grid2, i+1, j)
}
