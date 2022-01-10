package _018_4sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	res := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		// 这里要注意一下...不要写成 i+1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			left, right := j+1, len(nums)-1
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for left < right {
				// 重新定义坐标, n1, n2要写在compare 循环里面，不然不会更新...
				n1, n2 := nums[left], nums[right]
				if nums[i]+nums[j]+n1+n2 == target {
					res = append(res, []int{nums[i], nums[j], n1, n2})
					for left < right && n1 == nums[left] {
						left++
					}
					for left < right && n2 == nums[right] {
						right--
					}
				} else if nums[i]+nums[j]+n1+n2 < target {
					left++
				} else if nums[i]+nums[j]+n1+n2 > target {
					right--
				}
			}
		}
	}

	return res
}
