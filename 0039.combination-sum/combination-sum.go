package _039_combination_sum

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var dfs func(cur int, curidx int, tmp []int, res *[][]int)
	dfs = func(cur int, curidx int, tmp []int, res *[][]int) {
		if cur == 0 {
			// 这里为啥要copy：slice是引用类型，如果不复制，那后面的循环会更新值
			t := make([]int, len(tmp))
			copy(t, tmp)
			*res = append(*res, t)
		}

		if cur < 0 {
			return
		}

		for i := curidx; i < len(candidates); i++ {
			target -= candidates[i]
			tmp = append(tmp, candidates[i])
			dfs(target, i, tmp, res)
			tmp = tmp[:len(tmp)-1]
			target += candidates[i]
		}
	}

	dfs(target, 0, []int{}, &res)

	return res
}
