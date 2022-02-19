package _020_valid_parentheses

func isValid(s string) bool {
	var stack []rune
	pairRight := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	pairLeft := map[rune]bool{
		'(': true,
		'[': true,
		'{': true,
	}
	for _, ss := range s {
		// 遇到左括号的就入栈
		if _, exist := pairLeft[ss]; exist {
			stack = append(stack, ss)
		} else {
			// 遇到右括号，把栈顶元素pop出来，看能否匹配
			if len(stack) < 1 {
				return false
			}
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			val := pairRight[ss]
			if val != left {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
