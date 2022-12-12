package _039_combination_sum

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var dfs func(cur []int, target int, idx int)
	dfs = func(cur []int, target int, idx int) {
		if target == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		if target < 0 {
			return
		}

		// 当前处理函数
		for i := idx; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			// 只需要一个横坐标即可，不需要传入纵坐标
			dfs(cur, target-candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}

	dfs([]int{}, target, 0)

	return res
}
