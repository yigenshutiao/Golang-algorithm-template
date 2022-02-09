package _287_find_the_duplicate_number

func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1

	for l < r {
		mid := (l + r) >> 1
		cnt := 0
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}

		if cnt > mid {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
