package _136_single_number

// singleNumber 最优解：位运算
func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res ^= num
	}

	return res
}

// 自己的解
func singleNumbers(nums []int) int {
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i != j {
				flag = true
				break
			}
		}

		if !flag {
			return nums[i]
		}
	}

	return -1
}
