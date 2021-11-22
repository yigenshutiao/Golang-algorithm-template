package _438_find_all_anagrams_in_a_string

//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]

func findAnagrams(s string, p string) []int {

	if len(s) < len(p) {
		return nil
	}

	source := map[byte]int{}
	for i := 0; i < len(p); i++ {
		source[p[i]]++
	}

	var result []int

	target := map[byte]int{}

	left, right := 0, len(p)-1

	for i := left; i <= right; i++ {
		target[s[i]]++
	}

	for right < len(s) {
		if check(source, target) {
			result = append(result, left)
		}

		target[s[left]]--
		left++

		right++
		if right == len(s) {
			break
		}
		target[s[right]]++
	}

	return result
}

func check(source, target map[byte]int) bool {
	for k, v := range target {
		if source[k] != v {
			return false
		}
	}
	return true
}
