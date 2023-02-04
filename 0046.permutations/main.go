package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int

	var dfs func(cur []int, used []bool)
	dfs = func(cur []int, used []bool) {
		if len(cur) == len(nums) {
			t := make([]int, len(nums))
			copy(t, cur)
			res = append(res, t)
			return
		}

		for i := 0; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				cur = append(cur, nums[i])
				dfs(cur, used)
				cur = cur[:len(cur)-1]
				used[i] = false
			}

		}
	}

	used := make([]bool, len(nums))
	dfs([]int{}, used)

	return res
}

//
//func permutes(nums []int) [][]int {
//	var res [][]int
//	n := len(nums)
//	isUsed := make([]bool, n)
//	tmp := make([]int, n)
//
//	var dfs func(nums []int, isUsed []bool, cur []int, i int, res *[][]int)
//
//	dfs = func(nums []int, isUsed []bool, cur []int, idx int, res *[][]int) {
//		if idx == n {
//			t := make([]int, n)
//			copy(t, cur)
//			*res = append(*res, t)
//			return
//		}
//
//		for i := 0; i < len(nums); i++ {
//			if isUsed[i] == false {
//				cur[idx] = nums[i]
//				isUsed[i] = true
//				dfs(nums, isUsed, cur, idx+1, res)
//				isUsed[i] = false
//			}
//		}
//	}
//
//	dfs(nums, isUsed, tmp, 0, &res)
//
//	return res
//}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
