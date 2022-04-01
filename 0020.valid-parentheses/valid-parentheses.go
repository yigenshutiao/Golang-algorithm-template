package _020_valid_parentheses

func isValid(s string) bool {
	var stack []rune
	parenth := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, v := range s {
		if _, ok := parenth[v]; ok {
			// 如果是左序列，入栈
			stack = append(stack, v)
		} else {
			// 如果开始stack里面没有左序列
			if len(stack) < 1 {
				return false
			}
			// 如果是右序列，和栈顶元素匹配，匹配则继续，不匹配则假
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if v != parenth[top] {
				return false
			}
		}
	}
	// 栈长度为0，true
	if len(stack) == 0 {
		return true
	}

	return false
}
