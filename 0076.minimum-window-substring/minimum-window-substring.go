package _076_minimum_window_substring

func minWindow(s string, t string) string {

	source, target := make(map[byte]int), make(map[byte]int)

	for k := range t {
		source[t[k]]++
	}

	left, right := -1, -1
	lens := 99999999
	for l, r := 0, 0; r < len(s); r++ {
		target[s[r]]++

		for l <= r && check(source, target) {
			// 看长度，做减法
			if r-l+1 < lens {
				lens = r - l + 1
				left, right = l, r+1
			}
			target[s[l]]--
			l++
		}
	}

	if left == -1 {
		return ""
	}

	return s[left:right]
}

// check 看target里面是否包含source中的所有词
func check(source, target map[byte]int) bool {
	for k, cnt := range source {
		if cnt > target[k] {
			return false
		}
	}
	return true
}
