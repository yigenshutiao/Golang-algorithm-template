package _128_longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	numExist := make(map[int]bool, len(nums))

	for _, num := range nums {
		numExist[num] = true
	}

	var res int

	for i := 0; i < len(nums); i++ {
		// 看当前的数字左边的数字是否存在
		if _, exist := numExist[nums[i]-1]; !exist {
			// 如果不存在，计算从这个数字为起点的最大长度
			cur := 0
			t := nums[i]
			for {
				if _, exist := numExist[t]; exist {
					t++
					cur++
				} else {
					break
				}
			}

			if cur > res {
				res = cur
			}
		}
	}

	return res
}
