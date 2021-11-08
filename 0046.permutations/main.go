package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	tmp := make([]int, len(nums))

	dfs(0, nums, used, tmp, &res)

	return res
}

func dfs(cur int, nums []int, used []bool, tmp []int, res *[][]int) {

	if cur == len(nums) {
		t := make([]int, len(nums))
		copy(t, tmp)
		*res = append(*res, t)
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			tmp[cur] = nums[i]
			used[i] = true
			dfs(cur+1, nums, used, tmp, res)
			used[i] = false
		}
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
