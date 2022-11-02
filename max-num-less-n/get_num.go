package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 23121

	nums := []int{2, 4, 9}
	num := strconv.Itoa(n)
	var res int

	var flag bool
	var cur int

	var dfs func(idx int, flag bool, cur int) bool

	dfs = func(idx int, flag bool, cur int) bool {
		if idx == len(num) {
			res = cur
			return true
		}

		if flag {
			return dfs(idx+1, true, cur*10+nums[len(nums)-1])
		} else {
			val := num[idx]
			vv, _ := strconv.Atoi(string(val))
			for i := len(nums) - 1; i >= 0; i-- {
				if vv == nums[i] {
					if dfs(idx+1, false, cur*10+nums[i]) {
						return true
					}
				} else if vv > nums[i] {
					if dfs(idx+1, true, cur*10+nums[i]) {
						return true
					}
				}
			}

			if idx != 0 {
				return false
			}

			return dfs(idx+1, true, cur)
		}
	}

	dfs(0, flag, cur)

	fmt.Println(res)
}
