package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] != target {
		return []int{-1, -1}
	} else {
		start := l
		l, r := 0, len(nums)-1
		for l < r {
			mid := (l + r + 1) >> 1
			if nums[mid] <= target {
				l = mid
			} else {
				r = mid - 1
			}
		}
		end := r
		return []int{start, end}
	}
}

func main() {
	a := searchRange([]int{2, 2, 2, 2, 2}, 2)
	fmt.Print(a)
}
