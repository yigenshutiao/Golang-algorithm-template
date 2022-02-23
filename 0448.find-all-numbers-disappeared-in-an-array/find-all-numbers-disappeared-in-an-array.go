package _448_find_all_numbers_disappeared_in_an_array

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i] == i+1 { // 当前元素出现在它该出现的位置，无需交换
			i++
			continue
		}
		idealIdx := nums[i] - 1        // idealIdx：当前元素应该出现的位置
		if nums[i] == nums[idealIdx] { // 当前元素=它理应出现的位置上的现有元素，说明重复了
			i++
			continue
		}
		nums[idealIdx], nums[i] = nums[i], nums[idealIdx] // 不重复，进行交换
		// 这里不要i++，因为交换过来的数字本身也需要考察，需要交换到合适的位置上
		// 如果 i++ 就会跳过它，少考察了它
	}

	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 { // 值与索引 不对应
			res = append(res, i+1)
		}
	}
	return res
}
