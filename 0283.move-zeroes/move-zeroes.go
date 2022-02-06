package _283_move_zeroes

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
