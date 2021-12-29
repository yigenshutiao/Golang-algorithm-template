package _027_remove_element

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != val {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			res += 1
		}
	}
	return res
}

func removeElements(nums []int, val int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
