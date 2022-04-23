package _215_kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	// 找k大的，就是找pos小的
	resPos := len(nums) - k
	left, right := 0, len(nums)-1
	for {
		pos := getPos(nums, left, right)
		if pos == resPos {
			return nums[pos]
		} else if pos > resPos {
			// 如果找到的位置在结果的右边, 说明答案在左边，往左边收缩
			right = pos - 1
		} else if pos < resPos {
			// 如果找到位置在pos左边，说明答案在右边，往右边收缩
			left = pos + 1
		}
	}
}

func getPos(nums []int, left, right int) int {
	idx, pos := left, right
	for i := right; i > left; i-- {
		if nums[i] > nums[idx] {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos--
		}
	}
	nums[pos], nums[idx] = nums[idx], nums[pos]
	return pos
}
