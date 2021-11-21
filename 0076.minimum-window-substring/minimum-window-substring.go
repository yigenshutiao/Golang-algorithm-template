package _076_minimum_window_substring

// t中的值都在滑动窗口中
func check(ori, cnt map[byte]int) bool {
	for k, v := range ori {
		if cnt[k] < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	// ori 用来存t中各个byte的数量
	// cnt 用来存储滑动窗口中各个byte的数量
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	// lens是滑动窗口的长度，默认取个较大值
	lens := 9999999
	ansL, ansR := -1, -1

	for l, r := 0, 0; r < sLen; r++ {
		// 右窗口需小于源字符串，且目标中含有这个字符
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		// 滑动窗口长度需大于等于0
		for check(ori, cnt) && l <= r {

			if r-l+1 < lens {
				// 更新滑动窗口长度
				lens = r - l + 1
				// 把坐标也更新
				ansL, ansR = l, l+lens
			}
			// 在滑动窗口中，若左下标存在于目标值中，把滑动窗口对应的值减1
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			// 滑动窗口向右移动
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
