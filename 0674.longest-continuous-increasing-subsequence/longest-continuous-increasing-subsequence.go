package _674_longest_continuous_increasing_subsequence

func findLengthOfLCIS(nums []int) int {
	var res int
	var tmp int
	for i := 0; i < len(nums); i++ {
		tmp = 1
		cur := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > cur {
				cur = nums[j]
				tmp++
			} else {
				break
			}
		}
		if tmp > res {
			res = tmp
		}
	}

	return res
}
