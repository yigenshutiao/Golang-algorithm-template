package _039_combination_sum

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int

	var dfs func(res *[][]int, cur []int, sum int, idx int)

	dfs = func(res *[][]int, cur []int, sum int, idx int) {
		if sum > target {
			return
		}

		if sum == target {
			t := make([]int, len(cur))
			copy(t, cur)
			*res = append(*res, t)
			return
		}

		for i := idx; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			sum += candidates[i]
			// 这里应该传i，而不是idx,idx是起始坐标...
			dfs(res, cur, sum, i)
			sum -= candidates[i]
			cur = cur[:len(cur)-1]
		}
	}

	dfs(&res, cur, 0, 0)

	return res
}
