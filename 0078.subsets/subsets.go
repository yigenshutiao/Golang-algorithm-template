package _078_subsets

func subsets(nums []int) [][]int {
	var res [][]int

	var dfs func(res *[][]int, cur []int, startIndex int)

	dfs = func(res *[][]int, cur []int, startIndex int) {
		t := make([]int, len(cur))
		copy(t, cur)
		*res = append(*res, t)

		for i := startIndex; i < len(nums); i++ {
			cur = append(cur, nums[i])
			dfs(res, cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(&res, []int{}, 0)

	return res
}
