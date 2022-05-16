package _003_longest_substring_without_repeating_characters

// 输入: s = "abcabcbb"
// 输出: 3

// 滑动窗口的题，先加元素，再动坐标
func lengthOfLongestSubstring(s string) int {
	res := 0
	source := map[byte]int{}
	right := 0

	for left := 0; left < len(s); left++ {
		// 找到不重复的子序列
		for right < len(s) && source[s[right]] == 0 {
			source[s[right]]++
			right++
		}
		res = max(res, right-left)
		// 删除掉第一个元素，继续尝试找不重复子序列
		source[s[left]]--
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstrings(s string) int {
	source := make(map[byte]int, len(s))
	var res int

	right := 0
	for left := 0; left < len(s); left++ {
		for right < len(s) && source[s[right]] == 0 {
			source[s[right]]++
			right++
		}

		if res < right-left {
			res = right - left
		}

		source[s[left]]--

	}
	return res
}
