package _128_longest_consecutive_sequence

func longestConsecutive(nums []int) int {

	isExists := make(map[int]bool)
	res := 0

	for _, num := range nums {
		isExists[num] = true
	}

	for i := 0; i < len(nums); i++ {
		// 左值存在没必要计算，要计算最左值
		if _, exist := isExists[nums[i]-1]; !exist {
			cur := 0
			t := nums[i]
			for {
				if _, exist := isExists[t]; exist {
					cur++
					t++
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
