package _704_binary_search

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 这里需要是<=，如果是<, 在只有一个元素的情况下，run不起来
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
