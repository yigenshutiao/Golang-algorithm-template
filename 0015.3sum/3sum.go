package _015_3sum

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		target := nums[i]
		// 开头大于0，后面没得玩了，直接返回
		if target > 0 {
			break
		}

		// 这里是防止重复的优化，且只保留第一次的结果, 注意第一个条件
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 提前定义坐标很重要
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right]+target == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 注意标定变量
				n1, n2 := nums[left], nums[right]
				for left < right && nums[left] == n1 {
					left++
				}
				for left < right && nums[right] == n2 {
					right--
				}
			} else if nums[left]+nums[right]+target > 0 {
				right--
			} else if nums[left]+nums[right]+target < 0 {
				left++
			}
		}
	}
	return res
}
