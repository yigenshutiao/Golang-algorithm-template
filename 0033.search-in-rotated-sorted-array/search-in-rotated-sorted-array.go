package _033_search_in_rotated_sorted_array

// search有点没搞懂什么时候用 <= ,什么时候用<
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if target == nums[mid] {
			return mid
		}
		if nums[0] <= nums[mid] {
			// 左半部分有序
			if target >= nums[0] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 右半部分有序
			if target <= nums[len(nums)-1] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
