package _438_find_all_anagrams_in_a_string

//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
// findAnagrams 不要每次都把所有值重新输入，每次滑动一个值即可
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	// source是p组成的目标值
	source := map[byte]int{}
	for i := 0; i < len(p); i++ {
		source[p[i]]++
	}

	var result []int

	target := map[byte]int{}

	// 这里要注意一下，right应该初始化为len(p) - 1, 这是因为坐标从0开始计算
	left, right := 0, len(p)-1
	// target是滑动窗口值
	for i := left; i <= right; i++ {
		target[s[i]]++
	}

	for right < len(s) {
		if check(source, target) {
			result = append(result, left)
		}
		// 左边的值滑动走
		target[s[left]]--
		left++

		// 右边的值滑动来
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
