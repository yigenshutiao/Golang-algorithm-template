package main

import "fmt"

// 此问题用到两种二分模板来寻找边界
// 题目，在一个单调递增的序列里找到一个数字的起始、终止坐标，找不到返回[-1,-1]
func searchRange(num []int, target int) []int {
	if len(num) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(num)-1
	for l < r {
		mid := (l + r) >> 1
		if num[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if num[l] != target {
		return []int{-1, -1}
	} else {
		start := l
		l, r := 0, len(num)-1
		for l < r {
			mid := (l + r + 1) >> 1
			if num[mid] <= target {
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
	fmt.Println(searchRange([]int{1, 2, 4, 4, 5, 8, 9}, 7))
}
