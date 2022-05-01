package test

func topKFrequent(nums []int, k int) []int {
	// key:数字, val:数字出现的次数
	numCnt := map[int]int{}
	for i := 0; i < len(nums); i++ {
		numCnt[nums[i]]++
	}

	var a []int
	for k := range numCnt {
		a = append(a, k)
	}

	quickSort(a)

	top := len(a) - k

	var res []int
	isUsed := make(map[int]bool)
	for k := range a {
		isUsed[k] = false
	}

	for i := len(a) - 1; i >= top; i-- {
		n := a[i]
		for k, v := range numCnt {
			if v == n && isUsed[i] == false {
				res = append(res, k)
				isUsed[i] = true
			}
		}
	}

	return res
}

func quickSort(nums []int) {
	if len(nums) < 1 {
		return
	}
	idx, pos := 0, len(nums)-1

	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[idx] {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos--
		}
	}

	nums[idx], nums[pos] = nums[pos], nums[idx]

	quickSort(nums[:pos])
	quickSort(nums[pos+1:])
}
