package _001_two_sum

func twoSum(nums []int, target int) []int {
	s := make(map[int]bool, len(nums))
	for _, v := range nums {
		s[v] = true
	}

	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		if ok := s[t]; ok {
			for j := 0; i < len(nums); j++ {
				if nums[j] == t {
					return []int{i, j}
				}
			}

		}
	}
	return nil
}
