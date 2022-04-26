package _283_move_zeroes

// 双指针
func moveZeroes(nums []int) []int {
	// l想要指的是0值所处的位置
	// r是非0值的位置，要不断移动
	l, r := 0, 0
	n := len(nums)

	for r < n {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
	return nums
}

// 暴力解法
func moveZeroes2(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			t := i
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					i = j
				}
			}
			i = t
		}
	}
	return nums
}
