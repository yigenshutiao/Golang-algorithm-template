package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	tmp := make([]int, len(nums))

	var dfs func(cur int, nums []int, used []bool, tmp []int, res *[][]int)

	dfs = func(cur int, nums []int, used []bool, tmp []int, res *[][]int) {
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

	dfs(0, nums, used, tmp, &res)

	return res
}

func permutes(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	isUsed := make([]bool, n)
	tmp := make([]int, n)

	var dfs func(nums []int, isUsed []bool, cur []int, i int, res *[][]int)

	dfs = func(nums []int, isUsed []bool, cur []int, idx int, res *[][]int) {
		if idx == n {
			t := make([]int, n)
			copy(t, cur)
			*res = append(*res, t)
			return
		}

		for i := 0; i < len(nums); i++ {
			if isUsed[i] == false {
				cur[idx] = nums[i]
				isUsed[i] = true
				dfs(nums, isUsed, cur, idx+1, res)
				isUsed[i] = false
			}
		}
	}

	dfs(nums, isUsed, tmp, 0, &res)

	return res
}

func main() {
	fmt.Println(permutes([]int{1, 2, 3}))
}
