package offer03_shu_zu_zhong_zhong_fu_de_shu_zi_lcof

func findRepeatNumberMap(nums []int) int {
	numInfo := make(map[int]bool)

	for _, num := range nums {
		if !numInfo[num] {
			numInfo[num] = true
		} else {
			return num
		}
	}

	return -1
}

// findRepeatNumber nums[i] 需要为 i，一直交换，直到 nums[i]和i相等为止
func findRepeatNumber(nums []int) int {

	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}
