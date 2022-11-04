package _003_longest_substring_without_repeating_characters

// 输入: s = "abcabcbb"
// 输出: 3

// 滑动窗口的题，先加元素，再动坐标
func lengthOfLongestSubstring(s string) int {
	var res int

	keys := make(map[byte]int)

	// 要在外面声明，为了
	j := 0
	for i := 0; i < len(s); i++ {

		// 无重复时，调整右边界
		for j < len(s) && keys[s[j]] == 0 {
			keys[s[j]]++
			j++
		}

		if j-i > res {
			res = j - i
		}

		// 出现重复字母时，调整左边界
		keys[s[i]]--
	}

	return res
}

func lengthOfLongestSubstrings(s string) int {
	left, right := 0, 0
	var res int

	info := make(map[byte]int)

	for left < len(s) {
		// 节点right 往右走，一直走到有重复字符串
		for right < len(s) {
			if info[s[right]] < 1 {
				info[s[right]]++
				right++
			}
		}

		if right-left > res {
			res = right - left
		}

		info[s[left]]--
		left++
	}

	return res
}
