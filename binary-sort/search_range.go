package binary_sort

// 此问题用到两种二分模板来寻找边界
func searchRange(nums []int, target int) []int {
	// 此函数接收一个array和一个target, 返回target在array中的起始和终止index值，
	// 若找不到，返回[-1,-1]
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] != target {
		return []int{-1, -1}
	} else {
		start := l
		l, r := 0, len(nums)-1
		for l < r {
			mid := (l + r + 1) >> 1
			if nums[mid] <= target {
				l = mid
			} else {
				r = mid - 1
			}
		}
		end := r
		return []int{start, end}
	}
}
