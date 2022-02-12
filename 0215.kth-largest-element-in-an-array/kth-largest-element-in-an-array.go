package _215_kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	kk := len(nums) - k
	left, right := 0, len(nums)-1

	for {
		pos := partition(nums, left, right)
		if pos == kk {
			return nums[kk]
		} else if pos > kk {
			right = pos - 1
		} else if pos < kk {
			left = pos + 1
		}
	}
}

func partition(nums []int, left, right int) int {
	pivot := left
	pos := right
	for i := right; i >= left; i-- {
		if nums[i] > nums[pivot] {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos--
		}
	}
	nums[pos], nums[pivot] = nums[pivot], nums[pos]
	return pos
}
