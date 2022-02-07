package _283_move_zeroes

// 双指针
func moveZeroes(nums []int) {
	l, r := 0, 0
	n := len(nums)

	for r < n {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

// 暴力解法
func moveZeroes2(nums []int) {
	if len(nums) < 2 {
		return
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
}
