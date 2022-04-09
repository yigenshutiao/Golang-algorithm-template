package _076_minimum_window_substring

func minWindow(s string, t string) string {
	source, target := map[byte]int{}, map[byte]int{}
	for k := range t {
		// 这里注意通过坐标取value是byte类型，直接取是rune类型
		source[t[k]]++
	}

	lens := 99999
	start, end := -1, -1

	for l, r := 0, 0; r < len(s); r++ {
		target[s[r]] += 1
		for l <= r && check(source, target) {
			if r-l+1 < lens {
				lens = r - l + 1
				// 因为go的数组取不到r,所以end取r+1
				start, end = l, r+1
			}
			target[s[l]]--
			l++
		}
	}

	if start == -1 {
		return ""
	}
	return s[start:end]
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
