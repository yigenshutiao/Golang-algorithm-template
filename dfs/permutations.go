package dfs

func perDfs(cur int, n int, tmp []int, nums []int, used []bool, res *[][]int) {
	if cur == n {
		t := make([]int, n)
		copy(t, tmp)
		*res = append(*res, t)
		return
	}

	for i := 0; i < n; i++ {
		if !used[i] {
			used[i] = true
			tmp[i] = nums[cur]
			perDfs(cur+1, n, tmp, nums, used, res)
			used[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	tmp := make([]int, n)
	used := make([]bool, n)
	perDfs(0, n, tmp, nums, used, &res)
	return res
}
