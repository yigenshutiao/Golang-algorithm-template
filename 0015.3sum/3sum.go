package _015_3sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		// 这里是个优化，最小的数都大于0，相当于无解，直接退出
		if n1 > 0 {
			break
		}
		// 这里是用来去重的，如果num[i] == num[i-1], 跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 这里是要跳过相同的元素
				for l < r && nums[l] == n2 {
					l++
				}
				// 这里是要跳过相同的元素
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else if n1+n2+n3 > 0 {
				r--
			}
		}
	}
	return res
}
