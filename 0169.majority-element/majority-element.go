package _169_majority_element

func majorityElement(nums []int) int {
	candi, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			candi = num
		}

		if num == candi {
			count++
		} else {
			count--
		}
	}

	return candi
}
