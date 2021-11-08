package main

import "fmt"

func permute(nums []int) [][]int {
	// 回溯问题，主调用里面准备参数
	var res [][]int
	used := make([]bool, len(nums))
	tmp := make([]int, len(nums))

	// dfs里面算调用关系，注意把各种当前状态传进去，本题状态，cur
	permuteDfs(0, nums, used, tmp, &res)

	return res
}

func permuteDfs(cur int, nums []int, used []bool, tmp []int, res *[][]int) {

	if cur == len(nums) {
		t := make([]int, len(nums))
		copy(t, tmp)
		*res = append(*res, t)
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] { // 还可以选的元素
			tmp[cur] = nums[i] // 选择这个元素
			used[i] = true
			permuteDfs(cur+1, nums, used, tmp, res)
			used[i] = false
		}
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
