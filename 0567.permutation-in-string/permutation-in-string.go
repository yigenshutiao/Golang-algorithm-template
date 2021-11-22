package _567_permutation_in_string

func checkInclusion(s1 string, s2 string) bool {
	source := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		source[s1[i]]++
	}

	for left, right := 0, len(s1)-1; right < len(s2); {
		target := map[byte]int{}
		for i := left; i <= right; i++ {
			target[s2[i]]++
		}

		if check(source, target) {
			return true
		}
		left++
		right++
	}

	return false
}

func check(source, target map[byte]int) bool {
	for k, v := range target {
		if source[k] != v {
			return false
		}
	}
	return true
}
