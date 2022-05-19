package _033_search_in_rotated_sorted_array

// 局部二分法：将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
// 此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环.

// 所以这里再朴素二分的基础上，增加了寻找 有序子序列模块，要先找子序列，然后再查找
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
