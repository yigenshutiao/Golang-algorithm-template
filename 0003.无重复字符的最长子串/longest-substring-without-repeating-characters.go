package _003_longest_substring_without_repeating_characters

// 输入: s = "abcabcbb"
// 输出: 3

// 滑动窗口的题，先加元素，再动坐标
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	var res int

	info := make(map[byte]int)

	for left < len(s) {
		// 节点right 往右走，一直走到有重复字符串
		for right < len(s) && info[s[right]] < 1 {
			info[s[right]]++
			right++
		}

		if right-left > res {
			res = right - left
		}

		info[s[left]]--
		left++
	}

	return res
}
