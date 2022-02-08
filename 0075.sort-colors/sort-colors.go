package _075_sort_colors

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	index, pos := 0, len(nums)-1

	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[index] {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos--
		}
	}

	nums[pos], nums[index] = nums[index], nums[pos]

	sortColors(nums[:pos])
	sortColors(nums[pos+1:])
}
