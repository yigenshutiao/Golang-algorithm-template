package _287_find_the_duplicate_number

func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1

	for l < r {
		mid := (l + r) >> 1
		cnt := 0

		// 这里一定是小于 等于mid的个数
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}

		// 如果小于mid的cnt > mid,说明左边"]"有重复，否则说明右边(有重复
		if cnt > mid {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
