package _056_merge_intervals

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	var res [][]int
	// 先把 intervals进行排序，让他们以第一个元素的大小排序
	quickSort(intervals)

	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		// 如果前者的结尾 >= 后者的开头, 需要合并, 进行合并
		if cur[1] >= intervals[i][0] {
			cur[0] = cur[0]
			if cur[1] > intervals[i][1] {
				cur[1] = cur[1]
			} else {
				cur[1] = intervals[i][1]
			}
		} else if cur[1] < intervals[i][0] {
			// 如果前者的结尾 < 后者的开头，不需要合并, 前者追加到res, 后者变前者
			res = append(res, cur)
			cur = intervals[i]
		}
	}

	res = append(res, cur)

	return res
}

func quickSort(intervals [][]int) {
	if len(intervals) <= 1 {
		return
	}

	index := 0
	pos := len(intervals) - 1

	for i := len(intervals) - 1; i >= 1; i-- {
		if intervals[i][0] >= intervals[index][0] {
			intervals[i], intervals[pos] = intervals[pos], intervals[i]
			pos--
		}
	}

	intervals[pos], intervals[index] = intervals[index], intervals[pos]

	quickSort(intervals[:pos])
	quickSort(intervals[pos+1:])
}
