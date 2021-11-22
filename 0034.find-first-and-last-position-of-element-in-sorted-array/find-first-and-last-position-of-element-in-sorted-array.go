package _034_find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if nums[left] != target {
		return []int{-1, -1}
	}

	start := left

	left, right = 0, len(nums)-1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}

	end := right

	return []int{start, end}
}
