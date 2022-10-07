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

func findAnagrams(s string, p string) []int {
	source := map[byte]int{}
	for i := 0; i < len(p); i++ {
		source[p[i]]++
	}

	left, right := 0, len(p)-1

	target := map[byte]int{}
	for i := left; i <= right; i++ {
		target[s[i]]++
	}

	var res []int
	for right < len(s) {
		if check(target, source) {
			res = append(res, left)
		}

		target[s[left]]--
		left++

		right++
		if right == len(s) {
			break
		}
		target[s[right]]++
	}

	return res
}

func check(source, target map[byte]int) bool {
	for k, v := range target {
		if source[k] != v {
			return false
		}
	}
	return true
}
