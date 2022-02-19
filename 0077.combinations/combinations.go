package _077_combinations

func combine(n int, k int) [][]int {
	var res [][]int

	var dfs func(res *[][]int, cur []int, startIndex int)
	dfs = func(res *[][]int, cur []int, startIndex int) {
		if len(cur) == k {
			tmp := make([]int, k)
			copy(tmp, cur)
			*res = append(*res, tmp)
			return
		}

		// 单次逻辑
		for i := startIndex; i <= n; i++ {
			cur = append(cur, i)
			dfs(res, cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(&res, []int{}, 1)
	return res
}
